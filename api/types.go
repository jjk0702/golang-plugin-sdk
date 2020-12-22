package api


const (
	DataDirEnv    = "ci_data_dir"
	InputFileEnv  = "ci_data_input"
	OutputFileEnv = "ci_data_output"
)

type Status string

const (
	StatusSuccess Status = "success"
	StatusFailure Status = "failure"
	StatusError   Status = "error"
)

type DataType string

const (
	DataTypeString   DataType = "string"
)

type PluginBaseParam struct {
	PipelineVersion        string `json:"pipeline.version"`
	ProjectName            string `json:"project.name"`
	ProjectNameCn          string `json:"project.name.chinese"`
	PipelineId             string `json:"pipeline.id"`
	PipelineBuildNum       string `json:"pipeline.build.num"`
	PipelineBuildId        string `json:"pipeline.build.id"`
	PipelineName           string `json:"pipeline.name"`
	PipelineStartTimeMills string `json:"pipeline.time.start"`
	PipelineStartType      string `json:"pipeline.start.type"`
	PipelineStartUserId    string `json:"pipeline.start.user.id"`
	PipelineStartUserName  string `json:"pipeline.start.user.name"`
	Workspace            string `json:"workspace"`
}

type PluginOutput struct {
	Status    Status                 `json:"status"`
	Message   string                 `json:"message"`
	ErrorCode int                    `json:"errorCode"`
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
}

func NewPluginOutput() *PluginOutput {
	output := new(PluginOutput)
	output.Status = StatusSuccess
	output.Message = "success"
	output.Type = "default"
	output.Data = make(map[string]interface{})
	return output
}
type StringData struct {
	Type  DataType `json:"type"`
	Value string   `json:"value"`
}