package main

import (
	"flag"
	"fmt"
	"github.com/hmbilal/gofiber-start/internal/sample"
	"github.com/hmbilal/gofiber-start/pkg/checker"
	"github.com/hmbilal/gofiber-start/pkg/config"
	"github.com/hmbilal/gofiber-start/pkg/container"
	"github.com/hmbilal/gofiber-start/pkg/fiberLib"
	"github.com/hmbilal/gofiber-start/pkg/fiberLib/api"
	"github.com/hmbilal/gofiber-start/pkg/fiberLib/health"
	"log"
)

var configFile = flag.String("config", config.GetConfigFullFileName("config/settings.json"), "Path to config file")

func main() {
	app := fiberLib.New()

	// container
	con := container.NewContainer(
		configFile,
	)

	// handlers
	healthHandler := health.NewHandler(checker.NewCheckerPool())
	sampleHandler := sample.NewHandler(con)

	// routers
	routers := []api.Router{
		health.NewRouter(healthHandler),
		sample.NewRouter(sampleHandler),
	}

	for _, r := range routers {
		r.RegisterRoutes(app)
	}

	// Recover from panic
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", con.Cfg.Server.API.Port)))
}
