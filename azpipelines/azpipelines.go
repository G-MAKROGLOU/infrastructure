package azpipelines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// CreatePipelineFromYaml creates a pipeline from an existing YAML file in the repository.
// It is a no-op if the pipeline already exists.
func CreatePipelineFromYaml(details PipelineCreate) error {
	color.Cyan("AZ PIPELINES | CHECKING IF PIPELINE %s ALREADY EXISTS", details.Name)

	color.Cyan("AZ PIPELINES | RETRIEVING PIPELINES")
	pipelines, err := getPipelines(details.DevOPSOrg, details.Project)
	if err != nil {
		return err
	}
	color.Green("AZ PIPELINES | PIPELINES RETRIEVED SUCCESSFULLY")

	if !pipelineExists(pipelines, details.Name) {
		color.Yellow("AZ PIPELINES | PIPELINE %s DOES NOT EXIST. CREATING IT", details.Name)
		if _, err := createPipeline(details); err != nil {
			return err
		}
		color.Green("AZ PIPELINES | PIPELINE %s CREATED SUCCESSFULLY", details.Name)
	} else {
		color.Yellow("AZ PIPELINES | PIPELINE %s ALREADY EXISTS. SKIPPING PIPELINE CREATION", details.Name)
	}

	return nil
}

// QueuePipeline queues a deployment pipeline with the provided parameters.
func QueuePipeline(pipelineInfo PipelineCreate, parameters []string) (PipelineQueueRes, error) {
	var pipelineQueueRes PipelineQueueRes
	color.Cyan("AZ PIPELINES | QUEUEING PIPELINE %s", pipelineInfo.Name)

	cmd := getQueuePipelineBaseCmd(pipelineInfo.Project, pipelineInfo.DevOPSOrg, pipelineInfo.Name)

	if len(parameters) > 0 {
		cmd.Args = append(cmd.Args, "--parameters")
		cmd.Args = append(cmd.Args, parameters...)
	}

	var stderrBuf bytes.Buffer
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return pipelineQueueRes, fmt.Errorf("az pipelines run: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &pipelineQueueRes); err != nil {
		return pipelineQueueRes, fmt.Errorf("az pipelines run: failed to parse response: %w", err)
	}

	color.Green("AZ PIPELINES | PIPELINE %s QUEUED SUCCESSFULLY", pipelineInfo.Name)
	return pipelineQueueRes, nil
}

// GetPipelineStatus retrieves the current status and result of a single pipeline run.
// Bug fix: previously fetched ALL builds and filtered; now uses 'az pipelines build show --id'.
// Bug fix: previously returned nil error on az CLI failure, causing infinite polling.
func GetPipelineStatus(organization string, project string, pipelineID int) (PipelineStatus, error) {
	var pipelineStatus PipelineStatus
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "pipelines", "build", "show",
		"--id", strconv.Itoa(pipelineID),
		"--organization", organization,
		"--project", project,
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return pipelineStatus, fmt.Errorf("az pipelines build show: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &pipelineStatus); err != nil {
		return pipelineStatus, fmt.Errorf("az pipelines build show: failed to parse response: %w", err)
	}

	return pipelineStatus, nil
}

// DeletePipeline deletes an Azure DevOps pipeline by name.
// It is a no-op if the pipeline does not exist.
func DeletePipeline(org, project, name string) error {
	color.Cyan("AZ PIPELINES | DELETING PIPELINE %s", name)

	pipelines, err := getPipelines(org, project)
	if err != nil {
		return err
	}

	id := -1
	for _, p := range pipelines {
		if p.Name == name {
			id = p.ID
			break
		}
	}

	if id == -1 {
		color.Yellow("AZ PIPELINES | PIPELINE %s NOT FOUND. SKIPPING DELETION", name)
		return nil
	}

	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "pipelines", "delete",
		"--id", strconv.Itoa(id),
		"--organization", org,
		"--project", project,
		"--yes",
	)
	cmd.Stderr = &stderrBuf

	if _, err := cmd.Output(); err != nil {
		return fmt.Errorf("az pipelines delete: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	color.Green("AZ PIPELINES | PIPELINE %s DELETED SUCCESSFULLY", name)
	return nil
}

// getPipelines fetches all pipelines for the given org/project and returns a fresh slice.
// It is safe to call concurrently — no shared state is written.
func getPipelines(devopsOrg string, project string) ([]Pipeline, error) {
	var stderrBuf bytes.Buffer
	cmd := exec.Command("az", "pipelines", "list", "--organization", devopsOrg, "--project", project)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("az pipelines list: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	var pipelines []Pipeline
	if err := json.Unmarshal(out, &pipelines); err != nil {
		return nil, fmt.Errorf("az pipelines list: failed to parse response: %w", err)
	}
	return pipelines, nil
}

func pipelineExists(pipelines []Pipeline, name string) bool {
	for _, p := range pipelines {
		if p.Name == name {
			return true
		}
	}
	return false
}

func createPipeline(details PipelineCreate) (Pipeline, error) {
	var pipeline Pipeline
	var stderrBuf bytes.Buffer

	cmd := exec.Command("az", "pipelines", "create",
		"--name", details.Name,
		"--yaml-path", details.YamlPath,
		"--project", details.Project,
		"--repository", details.Repository,
		"--organization", details.DevOPSOrg,
		"--repository-type", "tfsgit",
		"--branch", details.Branch,
		"--skip-run",
	)
	cmd.Stderr = &stderrBuf

	out, err := cmd.Output()
	if err != nil {
		return pipeline, fmt.Errorf("az pipelines create: %w: %s", err, strings.TrimSpace(stderrBuf.String()))
	}

	if err := json.Unmarshal(out, &pipeline); err != nil {
		return pipeline, fmt.Errorf("az pipelines create: failed to parse response: %w", err)
	}
	return pipeline, nil
}

func getQueuePipelineBaseCmd(project string, organization string, name string) *exec.Cmd {
	return exec.Command("az", "pipelines", "run",
		"--project", project,
		"--organization", organization,
		"--name", name,
		"--verbose",
	)
}
