package azfunction

type (
	// FunctionApp ~ The az cli response when retrieving azure function apps
	FunctionApp struct {
		AppServicePlanID            string                        `json:"appServicePlanId"`
		AvailabilityState           string                        `json:"availabilityState"`
		ClientAffinityEnabled       bool                          `json:"clientAffinityEnabled"`
		ClientCertEnabled           bool                          `json:"clientCertEnabled"`
		ClientCertExclusionPaths    interface{}                   `json:"clientCertExclusionPaths"`
		ClientCertMode              string                        `json:"clientCertMode"`
		CloningInfo                 interface{}                   `json:"cloningInfo"`
		ContainerSize               int                           `json:"containerSize"`
		CustomDomainVerificationID  string                        `json:"customDomainVerificationId"`
		DailyMemoryTimeQuota        int                           `json:"dailyMemoryTimeQuota"`
		DefaultHostName             string                        `json:"defaultHostName"`
		Enabled                     bool                          `json:"enabled"`
		EnabledHostNames            []string                      `json:"enabledHostNames"`
		ExtendedLocation            interface{}                   `json:"extendedLocation"`
		HostNameSslStates           []functionAppHostNameSslState `json:"hostNameSslStates,squash"`
		HostNames                   []string                      `json:"hostNames"`
		HostNamesDisabled           bool                          `json:"hostNamesDisabled"`
		HostingEnvironmentProfile   interface{}                   `json:"hostingEnvironmentProfile"`
		HTTPSOnly                   bool                          `json:"httpsOnly"`
		HyperV                      bool                          `json:"hyperV"`
		ID                          string                        `json:"id"`
		Identity                    interface{}                   `json:"identity"`
		InProgressOperationID       interface{}                   `json:"inProgressOperationId"`
		IsDefaultContainer          interface{}                   `json:"isDefaultContainer"`
		IsXenon                     bool                          `json:"isXenon"`
		KeyVaultReferenceIdentity   string                        `json:"keyVaultReferenceIdentity"`
		Kind                        string                        `json:"kind"`
		LastModifiedTimeUtc         string                        `json:"lastModifiedTimeUtc"`
		Location                    string                        `json:"location"`
		MaxNumberOfWorkers          interface{}                   `json:"maxNumberOfWorkers"`
		Name                        string                        `json:"name"`
		OutboundIPAddresses         string                        `json:"outboundIpAddresses"`
		PossibleOutboundIPAddresses string                        `json:"possibleOutboundIpAddresses"`
		PublicNetworkAccess         interface{}                   `json:"publicNetworkAccess"`
		RedundancyMode              string                        `json:"redundancyMode"`
		RepositorySiteName          string                        `json:"repositorySiteName"`
		Reserved                    bool                          `json:"reserved"`
		ResourceGroup               string                        `json:"resourceGroup"`
		ScmSiteAlsoStopped          bool                          `json:"scmSiteAlsoStopped"`
		SiteConfig                  functionAppSiteConfig         `json:"siteConfig,squash"`
		SlotSwapStatus              interface{}                   `json:"slotSwapStatus"`
		State                       string                        `json:"state"`
		StorageAccountRequired      bool                          `json:"storageAccountRequired"`
		SuspendedTill               interface{}                   `json:"suspendedTill"`
		Tags                        interface{}                   `json:"tags"`
		TargetSwapSlot              interface{}                   `json:"targetSwapSlot"`
		TrafficManagerHostNames     interface{}                   `json:"trafficManagerHostNames"`
		Type                        string                        `json:"type"`
		UsageState                  string                        `json:"usageState"`
		VirtualNetworkSubnetID      interface{}                   `json:"virtualNetworkSubnetId"`
		VnetContentShareEnabled     bool                          `json:"vnetContentShareEnabled"`
		VnetImagePullEnabled        bool                          `json:"vnetImagePullEnabled"`
		VnetRouteAllEnabled         bool                          `json:"vnetRouteAllEnabled"`
	}

	functionAppHostNameSslState struct {
		CertificateResourceID interface{} `json:"certificateResourceId"`
		HostType              string      `json:"hostType"`
		IPBasedSslResult      interface{} `json:"ipBasedSslResult"`
		IPBasedSslState       string      `json:"ipBasedSslState"`
		Name                  string      `json:"name"`
		SslState              string      `json:"sslState"`
		Thumbprint            interface{} `json:"thumbprint"`
		ToUpdate              interface{} `json:"toUpdate"`
		ToUpdateIPBasedSsl    interface{} `json:"toUpdateIpBasedSsl"`
		VirtualIPv6           interface{} `json:"virtualIPv6"`
		VirtualIP             interface{} `json:"virtualIp"`
	}

	functionAppSiteConfig struct {
		AcrUseManagedIdentityCreds             bool        `json:"acrUseManagedIdentityCreds"`
		AcrUserManagedIdentityID               interface{} `json:"acrUserManagedIdentityId"`
		AlwaysOn                               bool        `json:"alwaysOn"`
		AntivirusScanEnabled                   interface{} `json:"antivirusScanEnabled"`
		APIDefinition                          interface{} `json:"apiDefinition"`
		APIManagementConfig                    interface{} `json:"apiManagementConfig"`
		AppCommandLine                         interface{} `json:"appCommandLine"`
		AppSettings                            interface{} `json:"appSettings"`
		AutoHealEnabled                        interface{} `json:"autoHealEnabled"`
		AutoHealRules                          interface{} `json:"autoHealRules"`
		AutoSwapSlotName                       interface{} `json:"autoSwapSlotName"`
		AzureMonitorLogCategories              interface{} `json:"azureMonitorLogCategories"`
		AzureStorageAccounts                   interface{} `json:"azureStorageAccounts"`
		ClusteringEnabled                      bool        `json:"clusteringEnabled"`
		ConnectionStrings                      interface{} `json:"connectionStrings"`
		Cors                                   interface{} `json:"cors"`
		CustomAppPoolIdentityAdminState        interface{} `json:"customAppPoolIdentityAdminState"`
		CustomAppPoolIdentityTenantState       interface{} `json:"customAppPoolIdentityTenantState"`
		DefaultDocuments                       interface{} `json:"defaultDocuments"`
		DetailedErrorLoggingEnabled            interface{} `json:"detailedErrorLoggingEnabled"`
		DocumentRoot                           interface{} `json:"documentRoot"`
		ElasticWebAppScaleLimit                interface{} `json:"elasticWebAppScaleLimit"`
		Experiments                            interface{} `json:"experiments"`
		FileChangeAuditEnabled                 interface{} `json:"fileChangeAuditEnabled"`
		FtpsState                              interface{} `json:"ftpsState"`
		FunctionAppScaleLimit                  int         `json:"functionAppScaleLimit"`
		FunctionsRuntimeScaleMonitoringEnabled interface{} `json:"functionsRuntimeScaleMonitoringEnabled"`
		HandlerMappings                        interface{} `json:"handlerMappings"`
		HealthCheckPath                        interface{} `json:"healthCheckPath"`
		HTTP20Enabled                          bool        `json:"http20Enabled"`
		HTTP20ProxyFlag                        interface{} `json:"http20ProxyFlag"`
		HTTPLoggingEnabled                     interface{} `json:"httpLoggingEnabled"`
		IPSecurityRestrictions                 interface{} `json:"ipSecurityRestrictions"`
		IPSecurityRestrictionsDefaultAction    interface{} `json:"ipSecurityRestrictionsDefaultAction"`
		JavaContainer                          interface{} `json:"javaContainer"`
		JavaContainerVersion                   interface{} `json:"javaContainerVersion"`
		JavaVersion                            interface{} `json:"javaVersion"`
		KeyVaultReferenceIdentity              interface{} `json:"keyVaultReferenceIdentity"`
		Limits                                 interface{} `json:"limits"`
		LinuxFxVersion                         string      `json:"linuxFxVersion"`
		LoadBalancing                          interface{} `json:"loadBalancing"`
		LocalMySQLEnabled                      interface{} `json:"localMySqlEnabled"`
		LogsDirectorySizeLimit                 interface{} `json:"logsDirectorySizeLimit"`
		MachineKey                             interface{} `json:"machineKey"`
		ManagedPipelineMode                    interface{} `json:"managedPipelineMode"`
		ManagedServiceIdentityID               interface{} `json:"managedServiceIdentityId"`
		Metadata                               interface{} `json:"metadata"`
		MinTLSCipherSuite                      interface{} `json:"minTlsCipherSuite"`
		MinTLSVersion                          interface{} `json:"minTlsVersion"`
		MinimumElasticInstanceCount            int         `json:"minimumElasticInstanceCount"`
		NetFrameworkVersion                    interface{} `json:"netFrameworkVersion"`
		NodeVersion                            interface{} `json:"nodeVersion"`
		NumberOfWorkers                        int         `json:"numberOfWorkers"`
		PhpVersion                             interface{} `json:"phpVersion"`
		PowerShellVersion                      interface{} `json:"powerShellVersion"`
		PreWarmedInstanceCount                 interface{} `json:"preWarmedInstanceCount"`
		PublicNetworkAccess                    interface{} `json:"publicNetworkAccess"`
		PublishingPassword                     interface{} `json:"publishingPassword"`
		PublishingUsername                     interface{} `json:"publishingUsername"`
		Push                                   interface{} `json:"push"`
		PythonVersion                          interface{} `json:"pythonVersion"`
		RemoteDebuggingEnabled                 interface{} `json:"remoteDebuggingEnabled"`
		RemoteDebuggingVersion                 interface{} `json:"remoteDebuggingVersion"`
		RequestTracingEnabled                  interface{} `json:"requestTracingEnabled"`
		RequestTracingExpirationTime           interface{} `json:"requestTracingExpirationTime"`
		RoutingRules                           interface{} `json:"routingRules"`
		RuntimeADUser                          interface{} `json:"runtimeADUser"`
		RuntimeADUserPassword                  interface{} `json:"runtimeADUserPassword"`
		ScmIPSecurityRestrictions              interface{} `json:"scmIpSecurityRestrictions"`
		ScmIPSecurityRestrictionsDefaultAction interface{} `json:"scmIpSecurityRestrictionsDefaultAction"`
		ScmIPSecurityRestrictionsUseMain       interface{} `json:"scmIpSecurityRestrictionsUseMain"`
		ScmMinTLSVersion                       interface{} `json:"scmMinTlsVersion"`
		ScmType                                interface{} `json:"scmType"`
		SitePort                               interface{} `json:"sitePort"`
		SitePrivateLinkHostEnabled             interface{} `json:"sitePrivateLinkHostEnabled"`
		StorageType                            interface{} `json:"storageType"`
		SupportedTLSCipherSuites               interface{} `json:"supportedTlsCipherSuites"`
		TracingOptions                         interface{} `json:"tracingOptions"`
		Use32BitWorkerProcess                  interface{} `json:"use32BitWorkerProcess"`
		VirtualApplications                    interface{} `json:"virtualApplications"`
		VnetName                               interface{} `json:"vnetName"`
		VnetPrivatePortsCount                  interface{} `json:"vnetPrivatePortsCount"`
		VnetRouteAllEnabled                    interface{} `json:"vnetRouteAllEnabled"`
		WebSocketsEnabled                      interface{} `json:"webSocketsEnabled"`
		WebsiteTimeZone                        interface{} `json:"websiteTimeZone"`
		WinAuthAdminState                      interface{} `json:"winAuthAdminState"`
		WinAuthTenantState                     interface{} `json:"winAuthTenantState"`
		WindowsConfiguredStacks                interface{} `json:"windowsConfiguredStacks"`
		WindowsFxVersion                       interface{} `json:"windowsFxVersion"`
		XManagedServiceIdentityID              interface{} `json:"xManagedServiceIdentityId"`
	}

	// CreateFunction represents the data of the Azure Function to be created
	CreateFunction struct {
		Name           string
		StorageAccount string
		Location       string
		ResourceGroup  string
		Os             string
		Runtime        string
		Settings       []Setting
	}

	// Setting represents the environment variables of the Azure Function
	Setting struct {
		Value string
		Name  string
	}
)
