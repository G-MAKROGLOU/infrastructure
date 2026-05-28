package azstorageaccount

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CreateAzureStorageAccount checks if a storage account exists and creates it if it doesn't.
func CreateAzureStorageAccount(details StorageAccountCreate) error {
	color.Cyan("AZ STORAGE ACCOUNT | CHECKING IF STORAGE ACCOUNT %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ STORAGE ACCOUNT | RETRIEVING STORAGE ACCOUNTS")
	storageAccounts, err := getStorageAccounts()
	if err != nil {
		return err
	}
	color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNTS RETRIEVED SUCCESSFULLY")

	if !storageAccountExists(storageAccounts, details.Name) {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s DOES NOT EXIST. CREATING IT", details.Name)
		if _, err := createStorageAccount(details); err != nil {
			return err
		}
		color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s CREATED SUCCESSFULLY", details.Name)
	} else {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s ALREADY EXISTS. SKIPPING STORAGE ACCOUNT CREATION", details.Name)
	}

	return nil
}

// DeleteAzureStorageAccount deletes an Azure Storage Account.
// It is a no-op if the storage account does not exist.
func DeleteAzureStorageAccount(name, resourceGroup string) error {
	color.Cyan("AZ STORAGE ACCOUNT | DELETING STORAGE ACCOUNT %s", name)

	storageAccounts, err := getStorageAccounts()
	if err != nil {
		return err
	}

	if !storageAccountExists(storageAccounts, name) {
		color.Yellow("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s DOES NOT EXIST. SKIPPING DELETION", name)
		return nil
	}

	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "storage", "account", "delete",
		"--name", name,
		"--resource-group", resourceGroup,
		"--yes",
	)
	cmd.Stderr = &stderrBuf

	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("az storage account delete: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	color.Green("AZ STORAGE ACCOUNT | STORAGE ACCOUNT %s DELETED SUCCESSFULLY", name)
	return nil
}

// getStorageAccounts fetches all storage accounts and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getStorageAccounts() ([]StorageAccount, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "storage", "account", "list")
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az storage account list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var storageAccounts []StorageAccount
	if err := json.Unmarshal(out, &storageAccounts); err != nil {
		return nil, fmt.Errorf("az storage account list: failed to parse response: %w", err)
	}
	return storageAccounts, nil
}

func storageAccountExists(storageAccounts []StorageAccount, name string) bool {
	for _, sa := range storageAccounts {
		if sa.Name == name {
			return true
		}
	}
	return false
}

func createStorageAccount(details StorageAccountCreate) (StorageAccount, error) {
	var storageAccount StorageAccount
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "storage", "account", "create",
		"--name", details.Name,
		"--location", details.Location,
		"--resource-group", details.ResourceGroup,
		"--sku", "Standard_LRS",
		"--allow-blob-public-access", "false",
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return storageAccount, fmt.Errorf("az storage account create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &storageAccount); err != nil {
		return storageAccount, fmt.Errorf("az storage account create: failed to parse response: %w", err)
	}
	return storageAccount, nil
}
