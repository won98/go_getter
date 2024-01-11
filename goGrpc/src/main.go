package main

import (
	"grpc_go/src/interfaces"
	"grpc_go/src/launch"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	os.Setenv("TZ", "UTC")
	time.Local = time.UTC
}

func main() {
	launcher := launch.LoadEnvironment("./environment.yaml")
	launcher.LaunchApplication()
	if launcher.Env.RpcBase.Enable {
		go interfaces.ServeRpc(launcher)
	}
	select {}
}
