package api

import (
	"errors"
	"github.com/fdev-ci/golang-plugin-sdk/log"
	"io/ioutil"
	"os"
	"strings"
	"encoding/json"
)

var gDataDir string
var gInputFile string
var gOutputFile string
var gPluginOutput *PluginOutput
var gAllPluginParam map[string]interface{}
var gPluginBaseParam *PluginBaseParam



func init() {
	gPluginOutput = NewPluginOutput()
	gDataDir = getDataDir()
	gInputFile = getInputFile()
	gOutputFile = getOutputFile()
	initPluginParam()
}

func initPluginParam() {
	err := LoadInputParam(&gAllPluginParam)
	if err != nil {
		log.Error("init plugin base param failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "init plugin base param failed", 16015100)
	}

	gPluginBaseParam = new(PluginBaseParam)
	err = LoadInputParam(gPluginBaseParam)
	if err != nil {
		log.Error("init plugin base param failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "init plugin base param failed", 16015100)
	}
}

func GetInputParam(name string) string {
	value := gAllPluginParam[name]
	if value == nil {
		return ""
	}
	strValue, ok := value.(string)
	if !ok {
		return ""
	}
	return strValue
}

func LoadInputParam(v interface{}) error {
	file := gDataDir + "/" + gInputFile
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error("load input param failed:", err.Error())
		return errors.New("load input param failed")
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		log.Error("parse input param failed:", err.Error())
		return errors.New("parse input param failed")
	}
	return nil
}


func getDataDir() string {
	dir := strings.TrimSpace(os.Getenv(DataDirEnv))
	if len(dir) == 0 {
		dir, _ = os.Getwd()
	}
	return dir
}


func getInputFile() string {
	file := strings.TrimSpace(os.Getenv(InputFileEnv))
	if len(file) == 0 {
		file = "input.json"
	}
	return file
}

func getOutputFile() string {
	file := strings.TrimSpace(os.Getenv(OutputFileEnv))
	if len(file) == 0 {
		file = "output.json"
	}
	return file
}

func GetOutputData(key string) interface{} {
	return gPluginOutput.Data[key]
}

func AddOutputData(key string, data interface{}) {
	gPluginOutput.Data[key] = data
}

func RemoveOutputData(key string) {
	delete(gPluginOutput.Data, key)
}


func WriteOutput() error {
	data, _ := json.Marshal(gPluginOutput)

	file := gDataDir + "/" + gOutputFile
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Error("write output failed: ", err.Error())
		return errors.New("write output failed")
	}
	return nil
}


func FinishBuild(status Status, msg string) {
	gPluginOutput.Message = msg
	gPluginOutput.Status = status
	WriteOutput()
	switch status {
	case StatusSuccess:
		os.Exit(0)
	case StatusFailure:
		os.Exit(1)
	case StatusError:
		os.Exit(2)
	default:
		os.Exit(0)
	}
}

func FinishBuildWithErrorCode(status Status, msg string, errorCode int) {
	gPluginOutput.Message = msg
	gPluginOutput.Status = status
	gPluginOutput.ErrorCode = errorCode
	WriteOutput()
	switch status {
	case StatusSuccess:
		os.Exit(0)
	case StatusFailure:
		os.Exit(1)
	case StatusError:
		os.Exit(2)
	default:
		os.Exit(0)
	}
}



func SetPluginOutputType(pluginOutputType string) {
	gPluginOutput.Type = pluginOutputType
}

func GetProjectName() string {
	return gPluginBaseParam.ProjectName
}

func GetProjectDisplayName() string {
	return gPluginBaseParam.ProjectNameCn
}

func GetPipelineId() string {
	return gPluginBaseParam.PipelineId
}

func GetPipelineName() string {
	return gPluginBaseParam.PipelineName
}

func GetPipelineBuildId() string {
	return gPluginBaseParam.PipelineBuildId
}

func GetPipelineBuildNumber() string {
	return gPluginBaseParam.PipelineBuildNum
}

func GetPipelineStartType() string {
	return gPluginBaseParam.PipelineStartType
}

func GetPipelineStartUserId() string {
	return gPluginBaseParam.PipelineStartUserId
}

func GetPipelineStartUserName() string {
	return gPluginBaseParam.PipelineStartUserName
}

func GetPipelineStartTimeMills() string {
	return gPluginBaseParam.PipelineStartTimeMills
}

func GetPipelineVersion() string {
	return gPluginBaseParam.PipelineVersion
}

func GetWorkspace() string {
	return gPluginBaseParam.Workspace
}

func NewStringData(value string) *StringData {
	return &StringData{
		Type:  DataTypeString,
		Value: value,
	}
}

