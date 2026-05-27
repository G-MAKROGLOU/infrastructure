package azappservice

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CreateAzureAppServicePlan checks if an app service plan exists and creates it if it doesn't.
func CreateAzureAppServicePlan(aspDetails AppServicePlanCreate) error {
	color.Cyan("AZ APPSERVICE | CHECKING IF APP SERVICE PLAN %s ALREADY EXISTS", aspDetails.Name)

	color.Cyan("AZ APPSERVICE | RETRIEVING APP SERVICE PLANS")
	appServicePlans, err := getAppServicePlans()
	if err != nil {
		return err
	}
	color.Cyan("AZ APPSERVICE | APP SERVICE PLANS RETRIEVED SUCCESSFULLY")

	if !appServicePlanExists(appServicePlans, aspDetails.Name) {
		color.Cyan("AZ APPSERVICE | APP SERVICE PLAN %s DOES NOT EXIST. CREATING IT", aspDetails.Name)
		if _, err := createAppServicePlan(aspDetails); err != nil {
			return err
		}
		color.Green("AZ APPSERVICE | APP SERVICE PLAN %s CREATED SUCCESSFULLY", aspDetails.Name)
	} else {
		color.Yellow("AZ APPSERVICE | APP SERVICE PLAN %s ALREADY EXISTS. SKIPPING APP SERVICE PLAN CREATION", aspDetails.Name)
	}

	return nil
}

// getAppServicePlans fetches all app service plans and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getAppServicePlans() ([]AppServicePlan, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "appservice", "plan", "list")
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az appservice plan list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var appServicePlans []AppServicePlan
	if err := json.Unmarshal(out, &appServicePlans); err != nil {
		return nil, fmt.Errorf("az appservice plan list: failed to parse response: %w", err)
	}
	return appServicePlans, nil
}

func appServicePlanExists(appServicePlans []AppServicePlan, name string) bool {
	for _, asp := range appServicePlans {
		if asp.Name == name {
			return true
		}
	}
	return false
}

func createAppServicePlan(aspDetails AppServicePlanCreate) (AppServicePlan, error) {
	var appServicePlan AppServicePlan
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "appservice", "plan", "create",
		"--resource-group", aspDetails.ResourceGroup,
		"--name", aspDetails.Name,
		"--sku", "F1",
		"--location", aspDetails.Location,
		"--per-site-scaling", "true",
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return appServicePlan, fmt.Errorf("az appservice plan create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &appServicePlan); err != nil {
		return appServicePlan, fmt.Errorf("az appservice plan create: failed to parse response: %w", err)
	}
	return appServicePlan, nil
}
