package domain

import (
	"fmt"
	"os"
	"os/user"
)

type Environment struct {
	Environment string             `yaml:"environment"`
	RpcBase     EnvironmentRpcBase `yaml:"rpc"`
	PodIP       string
	PodName     string
	NodeName    string
}

type EnvironmentLoader struct {
	Environment string      `yaml:"environment"`
	Development Environment `yaml:"development"`
}

func (e EnvironmentLoader) ToEnviroment() *Environment {

	podIP := os.Getenv("POD_IP")
	podName := os.Getenv("POD_NAME")
	nodeName := os.Getenv("NODE_NAME")
	user, _ := user.Current()
	if podIP == "" {
		if user != nil {
			podIP = fmt.Sprintf("development_%s", user.Username)
		} else {
			podIP = "localhost"
		}
	}
	if podName == "" {
		if user != nil {
			podName = fmt.Sprintf("development_%s", user.Username)
		} else {
			podName = "localhost"
		}
	}
	if nodeName == "" {
		if user != nil {
			nodeName = fmt.Sprintf("development_%s", user.Name)
		} else {
			nodeName = "localhost"
		}
	}

	e.Development.Environment = e.Environment
	e.Development.PodIP = podIP
	e.Development.PodName = podName
	e.Development.NodeName = nodeName
	return &e.Development

}

func (e *Environment) IsProduction() bool {
	if e.Environment == "development" {
		return true
	} else {
		return false
	}
}

type EnvironmentRpcBase struct {
	Enable bool `yaml:"enable"`
	Port   int  `yaml:"port"`
}
