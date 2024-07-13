package azstorageaccount

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/fatih/color"
)

var storageAccounts []StorageAccount

// CreateAzureStorageAccount ~ checks if a storage account exists and creates it if it doesn't
func CreateAzureStorageAccount(details StorageAccountCreate) error {
	color.Cyan("AZ STORAGE ACCOUNT | CHECKING IF STORAGE ACCOUNT %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ STORAGE ACCOUNT | RETRIEVING STORAGE ACCOUNTS")
	rgError := getStorageAccounts()
	if rgError != nil {
		return rgError
	}
	color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNTS RETRIEVED SUCCESSFULLY")

	exists, existsError := storageAccountExists(details.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s DOES NOT EXIST. CREATING IT", details.Name)
		_, rgCreateErr := createStorageAccount(details)

		if rgCreateErr != nil {
			return rgCreateErr
		}
		color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s CREATED SUCCESSFULLY", details.Name)
	}

	if exists {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s ALREADY EXISTS. SKIPPING STORAGE ACCOUNT CREATION", details.Name)
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

func createStorageAccount(details StorageAccountCreate) (StorageAccount, error) {
	var storageAccount StorageAccount
	saCreateOut, saErr := exec.Command("az", "storage", "account", "create", "--name", details.Name, "--location", details.Location, "--resource-group", details.ResourceGroup, "--sku", "Standard_LRS", "--allow-blob-public-access", "false").Output()
	if saErr != nil {
		return storageAccount, saErr
	}

	unmarshalErr := json.Unmarshal(saCreateOut, &storageAccount)

	if unmarshalErr != nil {
		return storageAccount, unmarshalErr
	}

	return storageAccount, nil
}
