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
func CreateAzureResourceGroup(rgName string, location string) error {
	color.Cyan("AZ RESOURCE GROUP | CHECKING IF RESOURCE GROUP %s ALREADY EXISTS", rgName)

	color.Cyan("AZ RESOURCE GROUP | RETRIEVING RESOURCE GROUPS")
	rgError := getResourceGroups()
	if rgError != nil {
		return rgError
	}
	color.Green("AZ RESOURCE GROUP | RESOURCE GROUPS RETRIEVED SUCCESSFULLY")

	exists, existsError := resourceGroupExists(rgName)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s DID NOT EXIST. CREATING IT", rgName)
		_, rgCreateErr := createResourceGroup(rgName, location)

		if rgCreateErr != nil {
			return rgCreateErr
		}
		color.Green("AZ RESOURCE GROUP | RESOURCE GROUP %s CREATED SUCCESSFULLY", rgName)
	}

	if exists {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s ALREADY EXISTS. SKIPPING RESOURCE GROUP CREATION", rgName)
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

func createResourceGroup(rgName string, location string) (ResourceGroup, error) {
	var resourceGroup ResourceGroup
	rgCreateOut, rgCreateErr := exec.Command("az", "group", "create", "--name", rgName, "--location", location).Output()

	if rgCreateErr != nil {
		return resourceGroup, rgCreateErr
	}

	unmarshalErr := json.Unmarshal(rgCreateOut, &resourceGroup)

	if unmarshalErr != nil {
		return resourceGroup, unmarshalErr
	}

	return resourceGroup, nil
}
