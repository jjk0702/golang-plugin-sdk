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
var gAtomOutput *AtomOutput
var gAllAtomParam map[string]interface{}
var gAtomBaseParam *AtomBaseParam



func init() {
	gAtomOutput = NewAtomOutput()
	gDataDir = getDataDir()
	gInputFile = getInputFile()
	gOutputFile = getOutputFile()
	initAtomParam()
}

func initAtomParam() {
	err := LoadInputParam(&gAllAtomParam)
	if err != nil {
		log.Error("init atom base param failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "init atom base param failed", 16015100)
	}

	gAtomBaseParam = new(AtomBaseParam)
	err = LoadInputParam(gAtomBaseParam)
	if err != nil {
		log.Error("init atom base param failed: ", err.Error())
		FinishBuildWithErrorCode(StatusError, "init atom base param failed", 16015100)
	}
}

func GetInputParam(name string) string {
	value := gAllAtomParam[name]
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
	return gAtomOutput.Data[key]
}

func AddOutputData(key string, data interface{}) {
	gAtomOutput.Data[key] = data
}

func RemoveOutputData(key string) {
	delete(gAtomOutput.Data, key)
}


func WriteOutput() error {
	data, _ := json.Marshal(gAtomOutput)

	file := gDataDir + "/" + gOutputFile
	err := ioutil.WriteFile(file, data, 0644)
	if err != nil {
		log.Error("write output failed: ", err.Error())
		return errors.New("write output failed")
	}
	return nil
}


func FinishBuild(status Status, msg string) {
	gAtomOutput.Message = msg
	gAtomOutput.Status = status
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
	gAtomOutput.Message = msg
	gAtomOutput.Status = status
	gAtomOutput.ErrorCode = errorCode
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



func SetAtomOutputType(atomOutputType string) {
	gAtomOutput.Type = atomOutputType
}

func GetProjectName() string {
	return gAtomBaseParam.ProjectName
}

func GetProjectDisplayName() string {
	return gAtomBaseParam.ProjectNameCn
}

func GetPipelineId() string {
	return gAtomBaseParam.PipelineId
}

func GetPipelineName() string {
	return gAtomBaseParam.PipelineName
}

func GetPipelineBuildId() string {
	return gAtomBaseParam.PipelineBuildId
}

func GetPipelineBuildNumber() string {
	return gAtomBaseParam.PipelineBuildNum
}

func GetPipelineStartType() string {
	return gAtomBaseParam.PipelineStartType
}

func GetPipelineStartUserId() string {
	return gAtomBaseParam.PipelineStartUserId
}

func GetPipelineStartUserName() string {
	return gAtomBaseParam.PipelineStartUserName
}

func GetPipelineStartTimeMills() string {
	return gAtomBaseParam.PipelineStartTimeMills
}

func GetPipelineVersion() string {
	return gAtomBaseParam.PipelineVersion
}

func GetWorkspace() string {
	return gAtomBaseParam.Workspace
}

