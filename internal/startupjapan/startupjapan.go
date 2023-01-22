package startupjapan

import (
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/config"
	"github.com/mazeyqian/go-gin-gee/internal/pkg/constants"
	"github.com/mazeyqian/go-gin-gee/internal/startupjapan/router"
)

func setConfiguration(configPath string, configType string) {
	config.Setup(configPath, configType)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func Run(configPath string, configType string) {
	// startup - begin
	cmdLines := constants.ScriptStartMsg
	// cmdLines += `cd /web/go-gin-gee;`
	cmdLines += `cd /Users/mazey/Web/Mazey/go-gin-gee;`
	cmdLines += `. ./scripts/docker-run.sh "docker.io/mazeyqian/go-gin-gee:v20230122155129-api" "WECOM_ROBOT_CHECK=b2d57746-7146-44f2-8207-86cb0ca832be";`
	cmdLines += constants.ScriptEndMsg
	// https://stackoverflow.com/questions/3985193/what-is-bin-sh-c
	cmd := exec.Command("/bin/sh", "-c", cmdLines)
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("error:", err)
	}
	log.Printf("result: %s", result)
	// startup - end
	setConfiguration(configPath, configType)
	conf := config.GetConfig()
	log.Println("conf:", conf)
	web := router.Setup()
	log.Println("go api running on port " + conf.Server.Port)
	log.Println("==================>")
	_ = web.Run(":" + conf.Server.Port)
}
