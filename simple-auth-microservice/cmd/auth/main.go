package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/everestp/simple-auth-microservice/config"
	authServerGRPC "github.com/everestp/simple-auth-microservice/internal/auth/delivery/grpc/server"
	"github.com/everestp/simple-auth-microservice/pkg/logger"
	"github.com/everestp/simple-auth-microservice/pkg/postgres"
	"github.com/everestp/simple-auth-microservice/pkg/redis"
	"github.com/everestp/simple-auth-microservice/pkg/utils"
	userService "github.com/everestp/simple-auth-microservice/proto"

	"github.com/joho/godotenv"

	

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
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

	psqlDb  , err := postgres.NewPsqlDB(cfg)
	if err !=  nil{
		appLogger.Fatalf("Postgress init %s",err)
	} else {
		appLogger.Info("Postgres connected , status :%v",psqlDb.Stats())
	}
  defer psqlDb.Close()
  redisClient := redis.NewRedisClient(cfg)
  defer redisClient.Close()
  appLogger.Info("Redis Connected")
	l, err := net.Listen("tcp", cfg.Server.Port)
	if err != nil{
		appLogger.Fatal(err)
	}
	

	server := grpc.NewServer(grpc.KeepaliveParams(
		keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
			Timeout: 15 * time.Second,
			MaxConnectionAge: 5 * time.Minute,
			
		},
	))

  if cfg.Server.Mode != "Prodiction"{
	reflection.Register(server)
  }
	
  authServerGRPC := authServerGRPC.NewAuthServerGRPC(appLogger ,cfg)
  userService.RegisterUserServiceServer(server ,authServerGRPC)
  appLogger.Info("Server is listening on port : %v",cfg.Server.Port)
  appLogger.Fatal(server.Serve(l))

}