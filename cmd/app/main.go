package main

import (
	"ai-hr-backend/internal/app"
	"ai-hr-backend/internal/config"
	"flag"
	"log"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "/etc/ai-hr-backend/config.yaml", "path to config file")
	flag.Parse()

	cfg := config.NewConfig()
	err := cfg.Load(cfgPath)
	if err != nil {
		log.Fatalf("Error load config: %v", err)
		return
	}

	a := app.NewApp()
	a.Run(*cfg)
}
