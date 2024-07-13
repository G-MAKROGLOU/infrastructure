package azstorageaccount

import "time"

type (
	// StorageAccount ~ the az cli response when retrieving storage accounts
	StorageAccount struct {
		AccessTier                            interface{}                    `json:"accessTier"`
		AllowBlobPublicAccess                 bool                           `json:"allowBlobPublicAccess"`
		AllowCrossTenantReplication           interface{}                    `json:"allowCrossTenantReplication"`
		AllowSharedKeyAccess                  interface{}                    `json:"allowSharedKeyAccess"`
		AllowedCopyScope                      interface{}                    `json:"allowedCopyScope"`
		AzureFilesIdentityBasedAuthentication interface{}                    `json:"azureFilesIdentityBasedAuthentication"`
		BlobRestoreStatus                     interface{}                    `json:"blobRestoreStatus"`
		CreationTime                          time.Time                      `json:"creationTime"`
		CustomDomain                          interface{}                    `json:"customDomain"`
		DefaultToOAuthAuthentication          bool                           `json:"defaultToOAuthAuthentication"`
		DNSEndpointType                       interface{}                    `json:"dnsEndpointType"`
		EnableHTTPSTrafficOnly                bool                           `json:"enableHttpsTrafficOnly"`
		EnableNfsV3                           interface{}                    `json:"enableNfsV3"`
		Encryption                            storageAccountEncryption       `json:"encryption"`
		ExtendedLocation                      interface{}                    `json:"extendedLocation"`
		FailoverInProgress                    interface{}                    `json:"failoverInProgress"`
		GeoReplicationStats                   interface{}                    `json:"geoReplicationStats"`
		ID                                    string                         `json:"id"`
		Identity                              interface{}                    `json:"identity"`
		ImmutableStorageWithVersioning        interface{}                    `json:"immutableStorageWithVersioning"`
		IsHnsEnabled                          interface{}                    `json:"isHnsEnabled"`
		IsLocalUserEnabled                    interface{}                    `json:"isLocalUserEnabled"`
		IsSftpEnabled                         interface{}                    `json:"isSftpEnabled"`
		KeyCreationTime                       storageAccountKeyCreationTime  `json:"keyCreationTime"`
		KeyPolicy                             interface{}                    `json:"keyPolicy"`
		Kind                                  string                         `json:"kind"`
		LargeFileSharesState                  interface{}                    `json:"largeFileSharesState"`
		LastGeoFailoverTime                   interface{}                    `json:"lastGeoFailoverTime"`
		Location                              string                         `json:"location"`
		MinimumTLSVersion                     string                         `json:"minimumTlsVersion"`
		Name                                  string                         `json:"name"`
		NetworkRuleSet                        storageAccountNetworkRuleset   `json:"networkRuleSet"`
		PrimaryEndpoints                      storageAccountPrimaryEndpoints `json:"primaryEndpoints"`
		PrimaryLocation                       string                         `json:"primaryLocation"`
		PrivateEndpointConnections            []interface{}                  `json:"privateEndpointConnections"`
		ProvisioningState                     string                         `json:"provisioningState"`
		PublicNetworkAccess                   interface{}                    `json:"publicNetworkAccess"`
		ResourceGroup                         string                         `json:"resourceGroup"`
		RoutingPreference                     interface{}                    `json:"routingPreference"`
		SasPolicy                             interface{}                    `json:"sasPolicy"`
		SecondaryEndpoints                    interface{}                    `json:"secondaryEndpoints"`
		SecondaryLocation                     interface{}                    `json:"secondaryLocation"`
		Sku                                   storageAccountSkus             `json:"sku"`
		StatusOfPrimary                       string                         `json:"statusOfPrimary"`
		StatusOfSecondary                     interface{}                    `json:"statusOfSecondary"`
		StorageAccountSkuConversionStatus     interface{}                    `json:"storageAccountSkuConversionStatus"`
		Tags                                  struct{}                       `json:"tags"`
		Type                                  string                         `json:"type"`
	}

	storageAccountEncryption struct {
		EncryptionIdentity              interface{}                      `json:"encryptionIdentity"`
		KeySource                       string                           `json:"keySource"`
		KeyVaultProperties              interface{}                      `json:"keyVaultProperties"`
		RequireInfrastructureEncryption interface{}                      `json:"requireInfrastructureEncryption"`
		Services                        storageAccountEncryptionServices `json:"services,squash"`
	}

	storageAccountEncryptionServices struct {
		Blob  storageAccountEncryptionService `json:"blob,squash"`
		File  storageAccountEncryptionService `json:"file,squash"`
		Queue interface{}                     `json:"queue"`
		Table interface{}                     `json:"table"`
	}

	storageAccountEncryptionService struct {
		Enabled         bool      `json:"enabled"`
		KeyType         string    `json:"keyType"`
		LastEnabledTime time.Time `json:"lastEnabledTime"`
	}

	storageAccountKeyCreationTime struct {
		Key1 time.Time `json:"key1"`
		Key2 time.Time `json:"key2"`
	}

	storageAccountNetworkRuleset struct {
		Bypass              string        `json:"bypass"`
		DefaultAction       string        `json:"defaultAction"`
		IPRules             []interface{} `json:"ipRules"`
		ResourceAccessRules interface{}   `json:"resourceAccessRules"`
		VirtualNetworkRules []interface{} `json:"virtualNetworkRules"`
	}

	storageAccountPrimaryEndpoints struct {
		Blob               string      `json:"blob"`
		Dfs                interface{} `json:"dfs"`
		File               string      `json:"file"`
		InternetEndpoints  interface{} `json:"internetEndpoints"`
		MicrosoftEndpoints interface{} `json:"microsoftEndpoints"`
		Queue              string      `json:"queue"`
		Table              string      `json:"table"`
		Web                interface{} `json:"web"`
	}

	storageAccountSkus struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	}

)
