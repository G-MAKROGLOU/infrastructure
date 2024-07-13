package azresourcegroup


type (
	// ResourceGroup ~ the az cli response when retrieving resource groups
	ResourceGroup struct {
		ID         string                  `json:"id"`
		Location   string                  `json:"location"`
		ManagedBy  string                  `json:"managedBy"`
		Name       string                  `json:"name"`
		Tags       []string                `json:"tags"`
		Properties resourceGroupProperties `json:"properties,squash"`
		Type       string                  `json:"type"`
	}

	resourceGroupProperties struct {
		ProvisioningState string `json:"provisioningState"`
	}
)