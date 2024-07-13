package azfunction

import (
	"fmt"
	"os/exec"

	"encoding/json"
	"errors"

	"github.com/fatih/color"
)

var (
	functionApps []FunctionApp
)

// CreateAzureFunction - checks if an azure function app exists and creates it if it doesn't
func CreateAzureFunction(funcDetails CreateFunction) error {
	color.Cyan("AZ FUNCTIONAPP | CHECKING IF FUNCTIONAPP %s ALREADY EXISTS", funcDetails.Name)

	color.Cyan("AZ FUNCTIONAPP | RETRIEVING FUNCTIONAPPS")
	faError := getFunctionApps()
	if faError != nil {
		return faError
	}
	color.Green("AZ FUNCTIONAPP | FUNCTIONAPPS RETRIEVED SUCCESSFULLY")

	exists, existsError := functionAppExists(funcDetails.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s DOES NOT EXIST. CREATING IT", funcDetails.Name)
		_, faCreateErr := createFunctionApp(funcDetails)

		if faCreateErr != nil {
			return faCreateErr
		}
		color.Green("AZ FUNCTIONAPP | FUNCTIONAPP %s CREATED SUCCESSFULLY", funcDetails.Name)
	}

	if exists {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s ALREADY EXISTS. SKIPPING FUNCTION APP CREATION", funcDetails.Name)
	}

	return nil
}

// SetAzureFunctionEnv - sets the environment variables for an azure function app
func SetAzureFunctionEnv(funcDetails CreateFunction) error {
	color.Cyan("AZ FUNCTIONAPP SETTINGS | UPDATING SETTINGS FOR FUNCTIONAPP %s", funcDetails.Name)

	cmd := exec.Command("az", "functionapp", "config", "appsettings", "set", "--name", funcDetails.Name, "--resource-group", funcDetails.ResourceGroup, "--settings")

	for _, setting := range funcDetails.Settings {
		arg := fmt.Sprintf("%s=\"%s\"", setting.Name, setting.Value)
		cmd.Args = append(cmd.Args, arg)
	}

	_, settingsErr := cmd.Output()

	if settingsErr != nil {
		return settingsErr
	}
	color.Green("AZ FUNCTIONAPP SETTINGS | SETTINGS FOR FUNCTIONAPP %s UPDATED SUCCESSFULLY", funcDetails.Name)

	return nil
}

func getFunctionApps() error {

	funcAppListOut, funcAppErr := exec.Command("az", "functionapp", "list").Output()

	if funcAppErr != nil {
		return funcAppErr
	}

	unMarshalErr := json.Unmarshal(funcAppListOut, &functionApps)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}

func functionAppExists(faName string) (bool, error) {
	exists := false

	if functionApps == nil {
		return exists, errors.New("Resource Groups not initialized yet")
	}

	for _, fa := range functionApps {
		if fa.Name == faName {
			exists = true
			break
		}
	}

	return exists, nil
}

func createFunctionApp(funcDetails CreateFunction) (FunctionApp, error) {
	var functionApp FunctionApp

	funcOut, funcErr := exec.Command("az", "functionapp", "create", "--resource-group", funcDetails.ResourceGroup, "--consumption-plan-location", funcDetails.Location, "--runtime", funcDetails.ResourceGroup, "--os-type", funcDetails.Os, "--functions-version", "4", "--name", funcDetails.Name, "--storage-account", funcDetails.StorageAccount).Output()

	if funcErr != nil {
		return functionApp, funcErr
	}

	unmarshalErr := json.Unmarshal(funcOut, &functionApp)

	if unmarshalErr != nil {
		return functionApp, unmarshalErr
	}

	return functionApp, nil
}
