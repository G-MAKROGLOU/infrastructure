package azlogin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	subscriptions []AzureSubscription
	// SelectedSubscription the subcription to use instead of hardcoding it
	SelectedSubscription AzureSubscription
)

// AzureLogin logins to azure and get all subscriptions
func AzureLogin() error {
	color.Cyan("AZ LOGIN => WAITING FOR LOGIN APPROVAL")

	output, loginErr := exec.Command("az", "login").Output()

	unmarshalErr := json.Unmarshal(output, &subscriptions)

	if unmarshalErr != nil {
		return unmarshalErr
	}

	if loginErr == nil {
		color.Cyan("AZ LOGIN => LOGIN SUCCESSFUL")
		return nil
	}

	return loginErr
}

// SelectSubscription select the subscription to be used for deployments etc.
func SelectSubscription() {
	var s string
	var subscriptionIndex int
	r := bufio.NewReader(os.Stdin)
	for {
		for index, sub := range subscriptions {
			fmt.Printf("%d) %s (%s)\n", index+1, sub.Name, sub.TenantID)
		}
		fmt.Fprint(os.Stderr, "Select the Azure subscription you would like to use: ")
		s, _ = r.ReadString('\n')

		index, err := strconv.Atoi(strings.TrimSpace(s))

		if s != "" && err != nil {
			color.Red("[ERR:] Invalid Input => %s", err.Error())
		}

		if index-1 >= 0 && index-1 < len(subscriptions) {
			subscriptionIndex = index
			break
		}
	}
	SelectedSubscription = subscriptions[subscriptionIndex-1]
	color.Cyan("[INFO:] AZURE SUBCRIPTION: %s (%s)", SelectedSubscription.Name, SelectedSubscription.TenantID)
}
