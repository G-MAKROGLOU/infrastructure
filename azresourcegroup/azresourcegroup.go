package azresourcegroup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CreateAzureResourceGroup - checks if a resource group exists and creates it if it doesn't.
func CreateAzureResourceGroup(details ResourceGroupCreate) error {
	color.Cyan("AZ RESOURCE GROUP | CHECKING IF RESOURCE GROUP %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ RESOURCE GROUP | RETRIEVING RESOURCE GROUPS")
	resourceGroups, rgError := getResourceGroups()
	if rgError != nil {
		return rgError
	}
	color.Green("AZ RESOURCE GROUP | RESOURCE GROUPS RETRIEVED SUCCESSFULLY")

	if !resourceGroupExists(resourceGroups, details.Name) {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s DID NOT EXIST. CREATING IT", details.Name)
		if _, err := createResourceGroup(details); err != nil {
			return err
		}
		color.Green("AZ RESOURCE GROUP | RESOURCE GROUP %s CREATED SUCCESSFULLY", details.Name)
	} else {
		color.Yellow("AZ RESOURCE GROUP | RESOURCE GROUP %s ALREADY EXISTS. SKIPPING RESOURCE GROUP CREATION", details.Name)
	}

	return nil
}

// getResourceGroups fetches all resource groups from Azure and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getResourceGroups() ([]ResourceGroup, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "group", "list")
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az group list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var resourceGroups []ResourceGroup
	if err := json.Unmarshal(out, &resourceGroups); err != nil {
		return nil, fmt.Errorf("az group list: failed to parse response: %w", err)
	}
	return resourceGroups, nil
}

func resourceGroupExists(resourceGroups []ResourceGroup, name string) bool {
	for _, rg := range resourceGroups {
		if rg.Name == name {
			return true
		}
	}
	return false
}

// DeleteAzureResourceGroup deletes a resource group and every resource inside it (cascade).
// This is intentionally destructive — only call this when a full teardown is desired.
func DeleteAzureResourceGroup(name string) error {
	color.Cyan("AZ RESOURCE GROUP | DELETING RESOURCE GROUP %s (cascade)", name)

	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "group", "delete", "--name", name, "--yes")
	cmd.Stderr = &stderrBuf

	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("az group delete: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	color.Green("AZ RESOURCE GROUP | RESOURCE GROUP %s DELETED SUCCESSFULLY", name)
	return nil
}

func createResourceGroup(details ResourceGroupCreate) (ResourceGroup, error) {
	var resourceGroup ResourceGroup
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "group", "create", "--name", details.Name, "--location", details.Location)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return resourceGroup, fmt.Errorf("az group create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &resourceGroup); err != nil {
		return resourceGroup, fmt.Errorf("az group create: failed to parse response: %w", err)
	}
	return resourceGroup, nil
}
