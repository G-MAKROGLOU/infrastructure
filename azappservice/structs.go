package azappservice

type (
	appServicePlanSku struct {
		Capabilities interface{} `json:"capabilities"`
		Capacity     int         `json:"capacity"`
		Family       string      `json:"family"`
		Locations    interface{} `json:"locations"`
		Name         string      `json:"name"`
		Size         string      `json:"size"`
		SkuCapacity  interface{} `json:"skuCapacity"`
		Tier         string      `json:"tier"`
	}

	// AppServicePlan ~ The az cli response when retrieving app service plans
	AppServicePlan struct {
		ElasticScaleEnabled       bool              `json:"elasticScaleEnabled"`
		ExtendedLocation          interface{}       `json:"extendedLocation"`
		FreeOfferExpirationTime   interface{}       `json:"freeOfferExpirationTime"`
		HostingEnvironmentProfile interface{}       `json:"hostingEnvironmentProfile"`
		HyperV                    bool              `json:"hyperV"`
		ID                        string            `json:"id"`
		IsSpot                    bool              `json:"isSpot"`
		IsXenon                   bool              `json:"isXenon"`
		Kind                      string            `json:"kind"`
		KubeEnvironmentProfile    interface{}       `json:"kubeEnvironmentProfile"`
		Location                  string            `json:"location"`
		MaximumElasticWorkerCount int               `json:"maximumElasticWorkerCount"`
		MaximumNumberOfWorkers    int               `json:"maximumNumberOfWorkers"`
		Name                      string            `json:"name"`
		NumberOfSites             int               `json:"numberOfSites"`
		NumberOfWorkers           int               `json:"numberOfWorkers"`
		PerSiteScaling            bool              `json:"perSiteScaling"`
		ProvisioningState         interface{}       `json:"provisioningState"`
		Reserved                  bool              `json:"reserved"`
		ResourceGroup             string            `json:"resourceGroup"`
		Sku                       appServicePlanSku `json:"sku"`
		SpotExpirationTime        interface{}       `json:"spotExpirationTime"`
		Status                    string            `json:"status"`
		Tags                      interface{}       `json:"tags"`
		TargetWorkerCount         int               `json:"targetWorkerCount"`
		TargetWorkerSizeID        int               `json:"targetWorkerSizeId"`
		Type                      string            `json:"type"`
		WorkerTierName            interface{}       `json:"workerTierName"`
		ZoneRedundant             bool              `json:"zoneRedundant"`
	}
	// AppServicePlanCreate represents the detauls of the  app service plan to be created
	AppServicePlanCreate struct {
		Name          string
		ResourceGroup string
		Location      string
	}
)
