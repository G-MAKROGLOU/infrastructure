package azwebapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

// CreateAzureWebApp checks if an azure web app already exists and creates it if it doesn't.
func CreateAzureWebApp(details WebAppCreate) error {
	color.Cyan("AZ WEBAPP | CHECKING IF WEBAPP %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ WEBAPP | RETRIEVING WEBAPPS")
	webApps, err := getWebApps()
	if err != nil {
		return err
	}
	color.Cyan("AZ WEBAPP | WEBAPPS RETRIEVED SUCCESSFULLY")

	if !webAppExists(webApps, details.Name) {
		color.Cyan("AZ WEBAPP | WEBAPP %s DOES NOT EXIST. CREATING IT", details.Name)
		if _, err := createWebApp(details); err != nil {
			return err
		}
		color.Green("AZ WEBAPP | WEBAPP %s CREATED SUCCESSFULLY", details.Name)
	} else {
		color.Yellow("AZ WEBAPP | WEBAPP %s ALREADY EXISTS. SKIPPING WEBAPP CREATION", details.Name)
	}

	return nil
}

// getWebApps fetches all web apps and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getWebApps() ([]WebApp, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "webapp", "list")
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az webapp list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var webApps []WebApp
	if err := json.Unmarshal(out, &webApps); err != nil {
		return nil, fmt.Errorf("az webapp list: failed to parse response: %w", err)
	}
	return webApps, nil
}

func webAppExists(webApps []WebApp, name string) bool {
	for _, wa := range webApps {
		if wa.Name == name {
			return true
		}
	}
	return false
}

func createWebApp(details WebAppCreate) (WebApp, error) {
	var webApp WebApp
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "webapp", "create",
		"--resource-group", details.ResourceGroup,
		"--plan", details.AppServicePlan,
		"--name", details.Name,
		"--runtime", details.Runtime,
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return webApp, fmt.Errorf("az webapp create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &webApp); err != nil {
		return webApp, fmt.Errorf("az webapp create: failed to parse response: %w", err)
	}
	return webApp, nil
}
