package azappservice

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/fatih/color"
)

var appServicePlans []AppServicePlan

// CreateAzureAppServicePlan - checkS if an app service plan exists and creates it if it doesn't
func CreateAzureAppServicePlan(aspDetails AppServicePlanCreate) error {
	color.Cyan("AZ APPSERVICE | CHECKING IF APP SERVICE PLAN %s ALREADY EXISTS", aspDetails.Name)

	color.Cyan("AZ APPSERVICE | RETRIEVING APP SERVICE PLANS")
	rgError := getAppServicePlans()
	if rgError != nil {
		return rgError
	}
	color.Cyan("AZ APPSERVICE | APP SERVICE PLANS RETRIEVED SUCCESSFULLY")

	exists, existsError := appServicePlanExists(aspDetails.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Cyan("AZ APPSERVICE | APP SERVICE PLAN %s DOES NOT EXIST. CREATING IT", aspDetails.Name)
		_, swCreateErr := createAppServicePlan(aspDetails)

		if swCreateErr != nil {
			return swCreateErr
		}
		color.Green("AZ APPSERVICE | APP SERVICE PLAN %s CREATED SUCCESSFULLY", aspDetails.Name)
		return nil
	}

	if exists {
		color.Yellow("AZ APPSERVICE | APP SERVICE PLAN %s ALREADY EXISTS. ABORTING ANY FURTHER OPERATIONS", aspDetails.Name)
	}

	return nil
}

func getAppServicePlans() error {
	aseListOut, aseErr := exec.Command("az", "appservice", "plan", "list").Output()

	if aseErr != nil {
		return aseErr
	}

	unMarshalErr := json.Unmarshal(aseListOut, &appServicePlans)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}

func appServicePlanExists(aseName string) (bool, error) {
	exists := false

	if appServicePlans == nil {
		return exists, errors.New("App Service Plans not initialized yet")
	}

	for _, sw := range appServicePlans {
		if sw.Name == aseName {
			exists = true
			break
		}
	}

	return exists, nil
}

func createAppServicePlan(aspDetails AppServicePlanCreate) (AppServicePlan, error) {
	var appServicePlan AppServicePlan

	aseCreateOut, aseCreateErr := exec.Command("az", "appservice", "plan", "create", "--resource-group", aspDetails.ResourceGroup, "--name", aspDetails.Name, "--sku", "F1", "--location", aspDetails.Location, "--per-site-scaling", "true").Output()

	if aseCreateErr != nil {
		return appServicePlan, aseCreateErr
	}

	unMarshalErr := json.Unmarshal(aseCreateOut, &appServicePlan)

	if unMarshalErr != nil {
		return appServicePlan, unMarshalErr
	}

	return appServicePlan, nil
}
