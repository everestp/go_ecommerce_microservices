package service

import (
	"github.com/everestp/simple-auth-microservice/config"
	"github.com/everestp/simple-auth-microservice/pkg/logger"
	userService "github.com/everestp/simple-auth-microservice/proto"
)




type usersService struct {
	userService.UnimplementedUserServiceServer
	logger logger.Logger
	cfg *config.Config

}





// Auth server constructor
func NewAuthServerGRPC(logger logger.Logger, cfg *config.Config) *usersService {
	return &usersService{
		logger: logger,
		cfg: cfg,
		
	}
}