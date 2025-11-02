package main

import (
	"flag"
	"os"

	"github.com/mbcoward3/greenlight/internal/server"
)

const version = "1.0.0"

func main() {
	cfg := &server.Config{
		Version: version,
	}

	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	server := server.New(cfg)

	server.Run()
	os.Exit(1)
}
