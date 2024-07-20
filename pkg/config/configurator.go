package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/tkanos/gonfig"
	"path"
	"path/filepath"
	"runtime"
)

type Configurator interface {
	GetConfiguration() *Configuration
}

type Configuration struct {
	Environment string `json:"environment" envconfig:"ENVIRONMENT"`
	Maintenance string `json:"maintenance"`
	Logger      struct {
		Level string `json:"level" envconfig:"LOGGER_LEVEL"`
	} `json:"logger"`
	Server struct {
		API struct {
			Port       string `json:"port" envconfig:"SERVER_API_PORT"`
			UnixSocket string `json:"unix_socket" envconfig:"SERVER_API_UNIX_SOCKET"`
		} `json:"api"`
	} `json:"server"`
	DB struct {
		DynamoDB struct {
			Endpoint   string `json:"endpoint" envconfig:"ENDPOINT"`
			TableNames struct {
				Projects string `json:"projects" envconfig:"TABLE_NAME_PROJECTS"`
			} `json:"table_names"`
		} `json:"dynamodb"`
	} `json:"db"`
}

type JSONConfigurator struct {
	configuration Configuration
}

func (c *JSONConfigurator) GetConfiguration() *Configuration {
	return &c.configuration
}

func NewJSONConfigurator(configFile *string) Configurator {
	var configuration Configuration

	err := gonfig.GetConf(*configFile, &configuration)
	if err != nil {
		panic(err)
	}

	return &JSONConfigurator{configuration: configuration}
}

func GetConfigFullFileName(fileName string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("Unable to get the filepath.")
	}

	rootFolder := filepath.Join(filepath.Dir(file), "..", "..")

	return path.Join(rootFolder, "/", fileName)
}
