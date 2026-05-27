package azlogin

import (
	"bufio"
	"bytes"
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
	// SelectedSubscription is the subscription chosen by the user for deployments.
	SelectedSubscription AzureSubscription
)

// AzureLogin logs in to Azure and populates the list of available subscriptions.
func AzureLogin() error {
	color.Cyan("AZ LOGIN => WAITING FOR LOGIN APPROVAL")

	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "login")
	cmd.Stderr = &stderrBuf

	output, loginErr := cmd.Output()

	// Bug fix: check loginErr before attempting to unmarshal. When login fails,
	// output may be empty or contain an error string, causing a misleading
	// "unexpected end of JSON input" error to surface instead of the real cause.
	if loginErr != nil {
		if stderrBuf.Len() > 0 {
			return fmt.Errorf("az login: %w: %s", loginErr, strings.TrimSpace(stderrBuf.String()))
		}
		return fmt.Errorf("az login: %w", loginErr)
	}

	if err := json.Unmarshal(output, &subscriptions); err != nil {
		return fmt.Errorf("az login: failed to parse subscription list: %w", err)
	}

	color.Cyan("AZ LOGIN => LOGIN SUCCESSFUL")
	return nil
}

// SelectSubscription prompts the user to choose an Azure subscription.
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

		trimmed := strings.TrimSpace(s)
		if trimmed == "" {
			color.Red("[ERR:] No input provided. Please enter a number.")
			continue
		}

		index, err := strconv.Atoi(trimmed)
		if err != nil {
			color.Red("[ERR:] Invalid input %q — please enter a number.", trimmed)
			continue
		}

		if index >= 1 && index <= len(subscriptions) {
			subscriptionIndex = index
			break
		}

		color.Red("[ERR:] %d is out of range. Please enter a number between 1 and %d.", index, len(subscriptions))
	}
	SelectedSubscription = subscriptions[subscriptionIndex-1]
	color.Cyan("[INFO:] AZURE SUBSCRIPTION: %s (%s)", SelectedSubscription.Name, SelectedSubscription.TenantID)
}
