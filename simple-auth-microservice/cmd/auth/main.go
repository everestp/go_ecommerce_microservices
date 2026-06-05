package main

import (
	"log"
	"os"

	"github.com/everestp/simple-auth-microservice/config"
	"github.com/everestp/simple-auth-microservice/pkg/logger"
	"github.com/everestp/simple-auth-microservice/pkg/utils"
	"github.com/joho/godotenv"
)

func main(){
	log.Println("Starting the application")
	 err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

		configPath := utils.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}



	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)
	appLogger.InitLogger()

	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s, SSL: %v", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode, cfg.Server.SSL)
	appLogger.Infof("Success parsed config: %v", cfg.Server.AppVersion)

}