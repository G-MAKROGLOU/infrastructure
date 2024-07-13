package azdevops

type (
	// ServiceAccountGet represents the details of the service account to be retrieved
	ServiceAccountGet struct {
		Organization string
		Project      string
	}
	// ServiceAccount ~ The az cli response when retrieving service accounts
	ServiceAccount struct {
		AdministratorsGroup              interface{}                        `json:"administratorsGroup"`
		Authorization                    serviceAccountAuthorization        `json:"authorization"`
		CreatedBy                        serviceAccountCreatedBy            `json:"createdBy"`
		Data                             serviceAccountData                 `json:"data"`
		Description                      string                             `json:"description"`
		GroupScopeID                     interface{}                        `json:"groupScopeId"`
		ID                               string                             `json:"id"`
		IsOutdated                       bool                               `json:"isOutdated"`
		IsReady                          bool                               `json:"isReady"`
		IsShared                         bool                               `json:"isShared"`
		Name                             string                             `json:"name"`
		OperationStatus                  serviceAccountOperationStatus      `json:"operationStatus"`
		Owner                            string                             `json:"owner"`
		ReadersGroup                     interface{}                        `json:"readersGroup"`
		ServiceEndpointProjectReferences []serviceEndpointProjectReferences `json:"serviceEndpointProjectReferences"`
		Type                             string                             `json:"type"`
		URL                              string                             `json:"url"`
	}

	serviceAccountAuthorizationParameters struct {
		Serviceprincipalid string `json:"serviceprincipalid"`
		TenantID           string `json:"tenantId"`
	}

	serviceAccountAuthorization struct {
		Parameters serviceAccountAuthorizationParameters `json:"parameters"`
		Scheme     string                                `json:"scheme"`
	}

	serviceAccountCreatedBy struct {
		Descriptor        string      `json:"descriptor"`
		DirectoryAlias    interface{} `json:"directoryAlias"`
		DisplayName       string      `json:"displayName"`
		ID                string      `json:"id"`
		ImageURL          string      `json:"imageUrl"`
		Inactive          interface{} `json:"inactive"`
		IsAadIdentity     interface{} `json:"isAadIdentity"`
		IsContainer       interface{} `json:"isContainer"`
		IsDeletedInOrigin interface{} `json:"isDeletedInOrigin"`
		ProfileURL        interface{} `json:"profileUrl"`
		UniqueName        string      `json:"uniqueName"`
		URL               string      `json:"url"`
	}

	serviceAccountData struct {
		AppObjectID              string `json:"appObjectId"`
		AzureSpnPermissions      string `json:"azureSpnPermissions"`
		AzureSpnRoleAssignmentID string `json:"azureSpnRoleAssignmentId"`
		CreationMode             string `json:"creationMode"`
		Environment              string `json:"environment"`
		ScopeLevel               string `json:"scopeLevel"`
		SpnObjectID              string `json:"spnObjectId"`
		SubscriptionID           string `json:"subscriptionId"`
		SubscriptionName         string `json:"subscriptionName"`
	}

	serviceAccountOperationStatus struct {
		Severity      interface{} `json:"severity"`
		State         string      `json:"state"`
		StatusMessage string      `json:"statusMessage"`
	}

	projectReference struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	serviceEndpointProjectReferences struct {
		Description      string           `json:"description"`
		Name             string           `json:"name"`
		ProjectReference projectReference `json:"projectReference"`
	}
)
