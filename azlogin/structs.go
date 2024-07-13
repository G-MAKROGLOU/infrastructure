package azlogin

type (
	// AzureSubscription ~ ~ The az cli response after logging in
	AzureSubscription struct {
		CloudName        string        `json:"cloudName"`
		HomeTenantID     string        `json:"homeTenantId"`
		ID               string        `json:"id"`
		IsDefault        bool          `json:"isDefault"`
		ManagedByTenants []interface{} `json:"managedByTenants"`
		Name             string        `json:"name"`
		State            string        `json:"state"`
		TenantID         string        `json:"tenantId"`
		User             user          `json:"user,squash"`
	}

	user struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
)
