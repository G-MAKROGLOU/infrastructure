package azfunction

import (
	"fmt"
	"os/exec"

	"encoding/json"
	"errors"

	"github.com/G-MAKROGLOU/infrastructure"
	"github.com/fatih/color"
)

var (
	functionApps []FunctionApp
)

// CreateAzureFunction - checks if an azure function app exists and creates it if it doesn't
func CreateAzureFunction(funcApp infrastructure.AppDetails) error {
	color.Cyan("AZ FUNCTIONAPP | CHECKING IF FUNCTIONAPP %s ALREADY EXISTS", funcApp.Name)

	color.Cyan("AZ FUNCTIONAPP | RETRIEVING FUNCTIONAPPS")
	faError := getFunctionApps()
	if faError != nil {
		return faError
	}
	color.Green("AZ FUNCTIONAPP | FUNCTIONAPPS RETRIEVED SUCCESSFULLY")

	exists, existsError := functionAppExists(funcApp.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s DOES NOT EXIST. CREATING IT", funcApp.Name)
		_, faCreateErr := createFunctionApp(funcApp)

		if faCreateErr != nil {
			return faCreateErr
		}
		color.Green("AZ FUNCTIONAPP | FUNCTIONAPP %s CREATED SUCCESSFULLY", funcApp.Name)
	}

	if exists {
		color.Yellow("AZ FUNCTIONAPP | FUNCTIONAPP %s ALREADY EXISTS. SKIPPING FUNCTION APP CREATION", funcApp.Name)
	}

	return nil
}

// SetAzureFunctionEnv - sets the environment variables for an azure function app
func SetAzureFunctionEnv(faName string, rgName string, faSettings []infrastructure.AppSettings) error {
	color.Cyan("AZ FUNCTIONAPP SETTINGS | UPDATING SETTINGS FOR FUNCTIONAPP %s", faName)

	cmd := exec.Command("az", "functionapp", "config", "appsettings", "set", "--name", faName, "--resource-group", rgName, "--settings")

	for _, setting := range faSettings {
		arg := fmt.Sprintf("%s=\"%s\"", setting.Name, setting.Value)
		cmd.Args = append(cmd.Args, arg)
	}

	_, settingsErr := cmd.Output()

	if settingsErr != nil {
		return settingsErr
	}
	color.Green("AZ FUNCTIONAPP SETTINGS | SETTINGS FOR FUNCTIONAPP %s UPDATED SUCCESSFULLY", faName)

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

func createFunctionApp(funcApp infrastructure.AppDetails) (FunctionApp, error) {
	var functionApp FunctionApp

	funcOut, funcErr := exec.Command("az", "functionapp", "create", "--resource-group", funcApp.ResourceGroup, "--consumption-plan-location", funcApp.Location, "--runtime", funcApp.Runtime, "--os-type", funcApp.Os, "--functions-version", "4", "--name", funcApp.Name, "--storage-account", funcApp.StorageAccount).Output()

	if funcErr != nil {
		return functionApp, funcErr
	}

	unmarshalErr := json.Unmarshal(funcOut, &functionApp)

	if unmarshalErr != nil {
		return functionApp, unmarshalErr
	}

	return functionApp, nil
}
