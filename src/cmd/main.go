package main

import (
	"guide_go/src/internal/launch"
	"guide_go/src/transport/api"
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
	guidLancher()
}

func guidLancher() {
	launcher := launch.LoadEnvironment("./environment.yaml")
	launcher.LauncGuide()
	if launcher.Env.ApiBase.Enable {
		go api.ServeHttp(launcher)
	}
	select {}
}
