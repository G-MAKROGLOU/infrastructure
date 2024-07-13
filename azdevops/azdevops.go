package azdevops

import (
	"encoding/json"
	"os/exec"

	"github.com/fatih/color"
)

var (
	serviceAccounts []ServiceAccount
)

func getServiceAccounts(organization string, project string) error {
	color.Cyan("AZ DEVOPS | RETRIEVING SERVICE ACCOUNTS TO SELECT FOR THE DEPLOYMENT")

	serviceAccountOut, serviceAccountErr := exec.Command("az", "devops", "service-endpoint", "list", "--organization", organization, "--project", project).Output()
	if serviceAccountErr != nil {
		return serviceAccountErr
	}

	unMarshalErr := json.Unmarshal(serviceAccountOut, &serviceAccounts)

	if unMarshalErr != nil {
		return unMarshalErr
	}

	return nil
}