package configures

import (
	"flag"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type IndexConfig struct {
	Port int `yaml:"port"`

	Log struct {
		LogPath string `yaml:"logPath"`
		LogName string `yaml:"logName"`
	} `ymal:"log"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DbName   string `yaml:"name"`
		Debug    bool   `yaml:"debug"`
	} `yaml:"mysql"`
}

var Config IndexConfig

func InitConfigures() error {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter()
	configFile := flag.String("config", "conf/config.yml", "Path to the configuration file")
	flag.Parse()
	cfBytes, err := os.ReadFile(*configFile)
	if err == nil {
		var conf IndexConfig
		yaml.Unmarshal(cfBytes, &conf)
		Config = conf
		return nil
	} else {
		return err
	}
}
