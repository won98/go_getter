package launch

import (
	"context"
	"fmt"
	"guide_go/src/common"
	"guide_go/src/domain"
	"guide_go/src/infrastructure/grpc"
	"guide_go/src/infrastructure/grpc/proxyagent"
	"guide_go/src/infrastructure/mysql"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Launcher struct {
	Ctx   context.Context
	Env   *domain.Environment
	Mysql *mysql.Repository
}

var launcher *Launcher

func GetLauncher() *Launcher {
	return launcher
}
func (l *Launcher) LauncGuide() {
	if err := common.InitConfig(); err != nil {
		l.alert(err)
	}
	l.Mysql = mysql.InitMysql(&l.Env.Mysql)
	proxy := grpc.NewGrpcServer(l.Env)
	fmt.Println("proxy : ", proxy)
	proxyagent.Register(proxy)

}

func LoadEnvironment(f string) *Launcher {
	l := &Launcher{}
	l.Ctx = context.Background()
	l.LoadEnvironment(f)
	return l
}

func (l *Launcher) LoadEnvironment(f string) {
	yamlFile, err := os.ReadFile(f)

	if err != nil {
		log.Printf("LoadEnvironment Get Error %v ", err)
	}
	envFile := &domain.EnvironmentLoader{}
	err = yaml.Unmarshal(yamlFile, envFile)
	fmt.Println(yaml.Unmarshal(yamlFile, envFile))
	if err != nil {
		log.Printf("LoadEnvironment Unmarshal %v", err)
	}
	l.Env = envFile.ToEnviroment()
}

func (l *Launcher) alert(err error) {
	panic(err)
}
