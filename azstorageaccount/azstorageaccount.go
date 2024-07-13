package azstorageaccount

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/fatih/color"
)

var storageAccounts []StorageAccount

// CreateAzureStorageAccount ~ checks if a storage account exists and creates it if it doesn't
func CreateAzureStorageAccount(saName string, location string, resourceGroup string) error {
	color.Cyan("AZ STORAGE ACCOUNT | CHECKING IF STORAGE ACCOUNT %s ALREADY EXISTS", saName)

	color.Cyan("AZ STORAGE ACCOUNT | RETRIEVING STORAGE ACCOUNTS")
	rgError := getStorageAccounts()
	if rgError != nil {
		return rgError
	}
	color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNTS RETRIEVED SUCCESSFULLY")

	exists, existsError := storageAccountExists(saName)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s DOES NOT EXIST. CREATING IT", saName)
		_, rgCreateErr := createStorageAccount(saName, location, resourceGroup)

		if rgCreateErr != nil {
			return rgCreateErr
		}
		color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s CREATED SUCCESSFULLY", saName)
	}

	if exists {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s ALREADY EXISTS. SKIPPING STORAGE ACCOUNT CREATION", saName)
	}

	return nil
}

func getStorageAccounts() error {
	saListOut, resGroupErr := exec.Command("az", "storage", "account", "list").Output()

	if resGroupErr != nil {
		return resGroupErr
	}

	unMarshalErr := json.Unmarshal(saListOut, &storageAccounts)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}

func storageAccountExists(saName string) (bool, error) {
	exists := false

	if storageAccounts == nil {
		return exists, errors.New("Storage Accounts not initialized yet")
	}

	for _, sa := range storageAccounts {
		if sa.Name == saName {
			exists = true
			break
		}
	}

	return exists, nil
}

func createStorageAccount(saName string, location string, resourceGroup string) (StorageAccount, error) {
	var storageAccount StorageAccount
	saCreateOut, saErr := exec.Command("az", "storage", "account", "create", "--name", saName, "--location", location, "--resource-group", resourceGroup, "--sku", "Standard_LRS", "--allow-blob-public-access", "false").Output()
	if saErr != nil {
		return storageAccount, saErr
	}

	unmarshalErr := json.Unmarshal(saCreateOut, &storageAccount)

	if unmarshalErr != nil {
		return storageAccount, unmarshalErr
	}

	return storageAccount, nil
}
