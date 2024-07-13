package azresourcegroup

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/fatih/color"
)

var (
	resourceGroups []ResourceGroup
)

// CreateAzureResourceGroup - check is a resource group exists and creates it if it doesn't
func CreateAzureResourceGroup(details ResourceGroupCreate) error {
	color.Cyan("AZ RESOURCE GROUP | CHECKING IF RESOURCE GROUP %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ RESOURCE GROUP | RETRIEVING RESOURCE GROUPS")
	rgError := getResourceGroups()
	if rgError != nil {
		return rgError
	}
	color.Green("AZ RESOURCE GROUP | RESOURCE GROUPS RETRIEVED SUCCESSFULLY")

	exists, existsError := resourceGroupExists(details.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s DID NOT EXIST. CREATING IT", details.Name)
		_, rgCreateErr := createResourceGroup(details)

		if rgCreateErr != nil {
			return rgCreateErr
		}
		color.Green("AZ RESOURCE GROUP | RESOURCE GROUP %s CREATED SUCCESSFULLY", details.Name)
	}

	if exists {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s ALREADY EXISTS. SKIPPING RESOURCE GROUP CREATION", details.Name)
	}

	return nil
}

func getResourceGroups() error {

	resGroupListOut, resGroupErr := exec.Command("az", "group", "list").Output()

	if resGroupErr != nil {
		return resGroupErr
	}

	unMarshalErr := json.Unmarshal(resGroupListOut, &resourceGroups)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}

func resourceGroupExists(rgName string) (bool, error) {
	exists := false

	if resourceGroups == nil {
		return exists, errors.New("Resource Groups not initialized yet")
	}

	for _, rg := range resourceGroups {
		if rg.Name == rgName {
			exists = true
			break
		}
	}

	return exists, nil
}

func createResourceGroup(details ResourceGroupCreate) (ResourceGroup, error) {
	var resourceGroup ResourceGroup
	rgCreateOut, rgCreateErr := exec.Command("az", "group", "create", "--name", details.Name, "--location", details.Location).Output()

	if rgCreateErr != nil {
		return resourceGroup, rgCreateErr
	}

	unmarshalErr := json.Unmarshal(rgCreateOut, &resourceGroup)

	if unmarshalErr != nil {
		return resourceGroup, unmarshalErr
	}

	return resourceGroup, nil
}
