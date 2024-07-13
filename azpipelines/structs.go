package azpipelines

import "time"


type (
	// AgentPool JSON representation of AgentPool
	AgentPool struct {
		AgentCloudID  interface{}        `json:"agentCloudId"`
		AutoProvision bool               `json:"autoProvision"`
		AutoSize      bool               `json:"autoSize"`
		AutoUpdate    bool               `json:"autoUpdate"`
		CreatedBy     agentPoolCreatedBy `json:"createdBy,squash"`
		CreatedOn     time.Time          `json:"createdOn"`
		ID            int                `json:"id"`
		IsHosted      bool               `json:"isHosted"`
		IsLegacy      bool               `json:"isLegacy"`
		Name          string             `json:"name"`
		Options       string             `json:"options"`
		Owner         agentPoolOwner     `json:"owner,squash"`
		PoolType      string             `json:"poolType"`
		Properties    interface{}        `json:"properties"`
		Scope         string             `json:"scope"`
		Size          int                `json:"size"`
		TargetSize    interface{}        `json:"targetSize"`
	}


	agentPoolCreatedBy struct {
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

	agentPoolOwner struct {
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

	// Pipeline ~ the az cli response when retrieving the pipelines
	Pipeline struct {
		AuthoredBy           pipelineAuthoredBy `json:"authoredBy"`
		CreatedDate          time.Time          `json:"createdDate"`
		DraftOf              interface{}        `json:"draftOf"`
		Drafts               []interface{}      `json:"drafts"`
		ID                   int                `json:"id"`
		LatestBuild          interface{}        `json:"latestBuild"`
		LatestCompletedBuild interface{}        `json:"latestCompletedBuild"`
		Metrics              interface{}        `json:"metrics"`
		Name                 string             `json:"name"`
		Path                 string             `json:"path"`
		Project              pipelineProject    `json:"project"`
		Quality              string             `json:"quality"`
		Queue                pipelineQueue      `json:"queue"`
		QueueStatus          string             `json:"queueStatus"`
		Revision             int                `json:"revision"`
		Type                 string             `json:"type"`
		URI                  string             `json:"uri"`
		URL                  string             `json:"url"`
	}

	pipelineAuthoredBy struct {
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

	pipelineProject struct {
		Abbreviation        interface{} `json:"abbreviation"`
		DefaultTeamImageURL interface{} `json:"defaultTeamImageUrl"`
		Description         string      `json:"description"`
		ID                  string      `json:"id"`
		LastUpdateTime      time.Time   `json:"lastUpdateTime"`
		Name                string      `json:"name"`
		Revision            int         `json:"revision"`
		State               string      `json:"state"`
		URL                 string      `json:"url"`
		Visibility          string      `json:"visibility"`
	}

	pipelineQueue struct {
		ID   int               `json:"id"`
		Name string            `json:"name"`
		Pool pipelineQueuePool `json:"pool"`
		URL  string            `json:"url"`
	}

	pipelineQueuePool struct {
		ID       int    `json:"id"`
		IsHosted bool   `json:"isHosted"`
		Name     string `json:"name"`
	}

	// PipelineStatus ~ the az cli response when polling the status of a pipeline
	PipelineStatus struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
		Result string `json:"result"`
	}

	// PipelineQueueRes ~ the az cli response after queueing a pipeline
	PipelineQueueRes struct {
		CreatedDate        time.Time                 `json:"createdDate"`
		FinalYaml          interface{}               `json:"finalYaml"`
		FinishedDate       interface{}               `json:"finishedDate"`
		ID                 int                       `json:"id"`
		Name               string                    `json:"name"`
		Pipeline           pipelineQueueResFolder    `json:"pipeline"`
		Resources          pipelineQueueResResources `json:"resources"`
		Result             string                    `json:"result"`
		State              string                    `json:"state"`
		TemplateParameters interface{}               `json:"templateParameters"`
		URL                string                    `json:"url"`
		Variables          interface{}               `json:"variables"`
	}

	pipelineQueueResFolder struct {
		Folder   string `json:"folder"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Revision int    `json:"revision"`
		URL      string `json:"url"`
	}

	pipelineQueueResResources struct {
		Repositories pipelineQueueResResourcesRepositories `json:"repositories"`
	}

	pipelineQueueResResourcesRepositories struct {
		Self pipelineQueueResResourcesRepositoriesSelf `json:"self"`
	}

	pipelineQueueResResourcesRepositoriesSelf struct {
		RefName    string                                              `json:"refName"`
		Repository pipelineQueueResResourcesRepositoriesSelfRepository `json:"repository"`
		Version    string
	}

	pipelineQueueResResourcesRepositoriesSelfRepository struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	}
)
