package launch

import (
	"context"
	"fmt"
	"grpc_go/src/domain"
	"os"
	"runtime"
	"strconv"

	"gopkg.in/yaml.v2"
)

type AppLauncher struct {
	Ctx context.Context
	Env *domain.Environment
}

var appLauncher *AppLauncher

func GetAppLauncher() *AppLauncher {
	return appLauncher
}

func LoadEnvironment(f string) *AppLauncher {
	l := &AppLauncher{}
	l.Ctx = context.Background()
	l.LoadEnvironment(f)
	return l
}

func (a *AppLauncher) LoadEnvironment(f string) {
	yamlFile, err := os.ReadFile(f)
	if err != nil {
		fmt.Println("LoadEnvironment Get Error %v", err)
	}
	envLoader := &domain.EnvironmentLoader{}
	err = yaml.Unmarshal(yamlFile, envLoader)
	if err != nil {
		fmt.Println("LoadEnvironment Unmarshal %v", err)
	}
	a.Env = envLoader.ToEnviroment()
}

func (l *AppLauncher) LaunchApplication() {
	fmt.Println("This project is dependent on", runtime.Version(), "version.")
	os.Setenv("IsProduction", strconv.FormatBool(l.Env.IsProduction()))
	var err error
	// l.Authentication, err = application.RallyShellStart()
	if err != nil {
		panic(err)
	}
	appLauncher = l
}
