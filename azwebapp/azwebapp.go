package azwebapp

import (
	"encoding/json"
	"errors"
	"os/exec"

	"github.com/G-MAKROGLOU/infrastructure"
	"github.com/fatih/color"
)

var (
	webApps []WebApp
)

// CreateAzureWebApp - checks if an azure web app already exists and creates it if it doesn't
func CreateAzureWebApp(wa infrastructure.AppDetails) error {
	color.Cyan("AZ WEBAPP | CHECKING IF WEBAPP %s ALREADY EXISTS", wa.Name)

	color.Cyan("AZ WEBAPP | RETRIEVING WEBAPPS")
	rgError := getWebApps()
	if rgError != nil {
		return rgError
	}
	color.Cyan("AZ WEBAPP | WEBAPPS RETRIEVED SUCCESSFULLY")

	exists, existsError := webAppExists(wa.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Cyan("AZ WEBAPP | WEBAPP %s DOES NOT EXIST. CREATING IT", wa.Name)
		_, waCreateErr := createWebApp(wa)

		if waCreateErr != nil {
			return waCreateErr
		}
		color.Green("AZ WEBAPP | WEBAPP %s CREATED SUCCESSFULLY", wa.Name)
		return nil
	}

	if exists {
		color.Yellow("AZ WEBAPP | WEBAPP %s ALREADY EXISTS. SKIPPING WEBAPP CREATION", wa.Name)
	}

	return nil
}

func getWebApps() error {
	waListOut, waErr := exec.Command("az", "webapp", "list").Output()

	if waErr != nil {
		return waErr
	}

	unMarshalErr := json.Unmarshal(waListOut, &webApps)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}

func webAppExists(waName string) (bool, error) {
	exists := false

	if webApps == nil {
		return exists, errors.New("Web Apps not initialized yet")
	}

	for _, sw := range webApps {
		if sw.Name == waName {
			exists = true
			break
		}
	}

	return exists, nil
}

func createWebApp(wa infrastructure.AppDetails) (WebApp, error) {
	// --plan ASP-WebApps-af28 --name test-octant --runtime NODE:16LTS
	var webApp WebApp

	waCreateOut, waCreateErr := exec.Command("az", "webapp", "create", "--resource-group", wa.ResourceGroup, "--plan", wa.AppServicePlan, "--name", wa.Name, "--runtime", wa.Runtime).Output()
	if waCreateErr != nil {
		return webApp, waCreateErr
	}

	unMarshalErr := json.Unmarshal(waCreateOut, &webApp)

	if unMarshalErr != nil {
		return webApp, unMarshalErr
	}

	return webApp, nil
}
