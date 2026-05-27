package azfunction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CreateAzureFunction checks if an azure function app exists and creates it if it doesn't.
func CreateAzureFunction(funcDetails CreateFunction) error {
	color.Cyan("AZ FUNCTIONAPP | CHECKING IF FUNCTIONAPP %s ALREADY EXISTS", funcDetails.Name)

	color.Cyan("AZ FUNCTIONAPP | RETRIEVING FUNCTIONAPPS")
	functionApps, err := getFunctionApps()
	if err != nil {
		return err
	}
	color.Green("AZ FUNCTIONAPP | FUNCTIONAPPS RETRIEVED SUCCESSFULLY")

	if !functionAppExists(functionApps, funcDetails.Name) {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s DOES NOT EXIST. CREATING IT", funcDetails.Name)
		if _, err := createFunctionApp(funcDetails); err != nil {
			return err
		}
		color.Green("AZ FUNCTIONAPP | FUNCTIONAPP %s CREATED SUCCESSFULLY", funcDetails.Name)
	} else {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s ALREADY EXISTS. SKIPPING FUNCTION APP CREATION", funcDetails.Name)
	}

	return nil
}

// SetAzureFunctionEnv sets the environment variables for an azure function app.
func SetAzureFunctionEnv(funcDetails CreateFunction) error {
	color.Cyan("AZ FUNCTIONAPP SETTINGS | UPDATING SETTINGS FOR FUNCTIONAPP %s", funcDetails.Name)

	cmd := exec.Command("az", "functionapp", "config", "appsettings", "set",
		"--name", funcDetails.Name,
		"--resource-group", funcDetails.ResourceGroup,
		"--settings",
	)

	for _, setting := range funcDetails.Settings {
		// Bug fix: pass name=value directly — exec.Command does not involve a shell,
		// so no quoting is needed. Adding literal '"' chars here would corrupt the value.
		cmd.Args = append(cmd.Args, setting.Name+"="+setting.Value)
	}

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("az functionapp config appsettings set: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	color.Green("AZ FUNCTIONAPP SETTINGS | SETTINGS FOR FUNCTIONAPP %s UPDATED SUCCESSFULLY", funcDetails.Name)
	return nil
}

// getFunctionApps fetches all function apps and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getFunctionApps() ([]FunctionApp, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "functionapp", "list")
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az functionapp list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var functionApps []FunctionApp
	if err := json.Unmarshal(out, &functionApps); err != nil {
		return nil, fmt.Errorf("az functionapp list: failed to parse response: %w", err)
	}
	return functionApps, nil
}

func functionAppExists(functionApps []FunctionApp, name string) bool {
	for _, fa := range functionApps {
		if fa.Name == name {
			return true
		}
	}
	return false
}

func createFunctionApp(funcDetails CreateFunction) (FunctionApp, error) {
	var functionApp FunctionApp
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "functionapp", "create",
		"--resource-group", funcDetails.ResourceGroup,
		"--consumption-plan-location", funcDetails.Location,
		"--runtime", funcDetails.Runtime, // Bug fix: was funcDetails.ResourceGroup
		"--os-type", funcDetails.Os,
		"--functions-version", "4",
		"--name", funcDetails.Name,
		"--storage-account", funcDetails.StorageAccount,
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return functionApp, fmt.Errorf("az functionapp create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &functionApp); err != nil {
		return functionApp, fmt.Errorf("az functionapp create: failed to parse response: %w", err)
	}
	return functionApp, nil
}
