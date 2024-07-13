package azpipelines

import (
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

var (
	agentPools []AgentPool
	pipelines  []Pipeline
)

// CreatePipelineFromYaml - creates a pipeline from an existing YAML file
func CreatePipelineFromYaml(details PipelineCreate) error {
	color.Cyan("AZ PIPELINES | CHECKING IF PIPELINE %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ PIPELINES | RETRIEVING PIPELINES")
	rgError := getPipelines(details.DevOPSOrg, details.Project)
	if rgError != nil {
		return rgError
	}
	color.Green("AZ PIPELINES | PIPELINES RETRIEVED SUCCESSFULLY")

	exists, existsError := pipelineExists(details.Name)
	if existsError != nil {
		return existsError
	}

	if !exists {
		color.Yellow("AZ PIPELINES | PIPELINE %s DOES NOT EXIST. CREATING IT", details.Name)

		_, pipelineErr := createPipeline(details)

		if pipelineErr != nil {
			return pipelineErr
		}

		color.Green("AZ PIPELINES | PIPELINE %s CREATED SUCCESSFULLY", details.Name)
	}

	if exists {
		color.Yellow("AZ PIPELINES | PIPELINE %s ALREADY EXISTS. SKIPPING PIPELINE CREATION", details.Name)
	}

	return nil
}

// QueuePipeline - queues a deployment pipeline
func QueuePipeline(pipelineInfo PipelineCreate, parameters []string) (PipelineQueueRes, error) {
	var pipelineQueueRes PipelineQueueRes
	color.Cyan("AZ PIPELINES | QUEUEING PIPELINE %s", pipelineInfo.Name)

	cmd := getQueuePipelineBaseCmd(pipelineInfo.Project, pipelineInfo.DevOPSOrg, pipelineInfo.Name)

	if len(parameters) > 0 {
		cmd.Args = append(cmd.Args, "--parameters")
		for _, param := range parameters {
			cmd.Args = append(cmd.Args, param)
		}
	}

	fmt.Println(cmd.String())

	queueOut, pipelineErr := cmd.Output()

	if pipelineErr != nil {
		return pipelineQueueRes, pipelineErr
	}

	unmarshalErr := json.Unmarshal(queueOut, &pipelineQueueRes)
	if unmarshalErr != nil {
		return pipelineQueueRes, unmarshalErr
	}

	color.Green("AZ PIPELINES | PIPELINE %s QUEUED SUCCESSFULLY", pipelineInfo.Name)
	return pipelineQueueRes, nil
}

// GetPipelineStatus - WIP retrieves the status of a pipeline
func GetPipelineStatus(organization string, project string, pipelineID int) (PipelineStatus, error) {
	var pipelineStatuses []PipelineStatus
	var pipelineStatus PipelineStatus
	//query := fmt.Sprintf("'[? id == `%d`].{id:id, status:status, result:result}[0]'", pipelineID)

	// using the jmespath --query does not print the desired output. retrieve all and filter
	out, err := exec.Command("az", "pipelines", "build", "list", "--organization", organization, "--project", project).Output()

	if err != nil {
		return pipelineStatus, nil
	}

	err = json.Unmarshal(out, &pipelineStatuses)
	if err != nil {
		return pipelineStatus, err
	}

	if len(pipelineStatuses) == 0 {
		return pipelineStatus, errors.New("Failed to retrieve pipeline status")
	}

	for _, pipeline := range pipelineStatuses {
		if pipeline.ID == pipelineID {
			pipelineStatus = pipeline
			break
		}
	}

	return pipelineStatus, nil
}

func getPipelines(devopsOrg string, project string) error {
	out, pipelineErr := exec.Command("az", "pipelines", "list", "--organization", devopsOrg, "--project", project).Output()

	if pipelineErr != nil {
		return pipelineErr
	}

	unmarshalErr := json.Unmarshal(out, &pipelines)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	return nil
}

func pipelineExists(pipelineName string) (bool, error) {

	exists := false

	if pipelines == nil {
		return exists, errors.New("Pipelines not initialized yet")
	}

	for _, pipeline := range pipelines {
		if pipeline.Name == pipelineName {
			exists = true
			break
		}
	}
	return exists, nil
}

func createPipeline(details PipelineCreate) (Pipeline, error) {
	var pipeline Pipeline

	pipelineOut, pipelineErr := exec.Command("az", "pipelines", "create",
	"--name", details.Name,
	"--yaml-path", details.YamlPath,
	"--project", details.Project,
	"--repository", details.Repository,
	"--organization", details.DevOPSOrg,
	"--repository-type", "tfsgit",
	"--branch", details.Branch,
	"--skip-run").Output()

	if pipelineErr != nil {
		return pipeline, pipelineErr
	}

	unmarshalErr := json.Unmarshal(pipelineOut, &pipeline)

	if unmarshalErr != nil {
		return pipeline, unmarshalErr
	}

	return pipeline, nil
}

func getQueuePipelineBaseCmd(project string, organization string, name string) *exec.Cmd {
	return exec.Command("az", "pipelines", "run",
		"--project", project,
		"--organization", organization,
		"--name", name,
		"--verbose")
}
