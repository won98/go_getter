package domain

import (
	"fmt"
	"os"
	"os/user"
)

type Environment struct {
	Environment string                 `yaml:"environment"`
	ApiBase     EnvironmentDefaultBase `yaml:"api"`
	Mysql       EnvironmentMysql       `yaml:"mysql"`
	Redis       EnvironmentClient      `yaml:"redis"`
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

type EnvironmentDefaultBase struct {
	Enable bool `yaml:"enable"`
	Port   int  `yaml:"port"`
}

type EnvironmentClient struct {
	Enable bool   `yaml:"enable"`
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
}

type EnvironmentMysql struct {
	Enable   bool   `yaml:"enable"`
	Project  string `yaml:"project"`
	Instance string `yaml:"instance"`
	Database string `yaml:"database"`
}
