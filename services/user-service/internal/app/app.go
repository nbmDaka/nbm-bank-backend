package app


import (
	"context"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/platform/database"
	grpcserver "github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/server/grpc"
	grpcuser "github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/delivery/grpc"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/application"
	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/infrastructure"

)

type App struct {
	db database.DB
	grpc *grpcserver.Server
}


func New(cfg config.Config) (*App,error){

	db, err := database.NewPostgres(
		cfg.Database,
	)

	if err != nil {
		return nil,err
	}

	err = database.RunMigrations(
		context.Background(),
		db,
	)

	if err != nil {
		return nil,err
	}


	userRepository := infrastructure.NewPostgresUserRepository(
		db,
	)


	userService := application.NewUserService(
		userRepository,
	)


	userHandler := grpcuser.NewHandler(
		userService,
	)


	return &App{
		db: db,
		grpc: grpcserver.NewServer(
			cfg.App.Port,
			userHandler,
		),
	},nil
}



func (a *App) Start() error {

	go func(){

		err := a.grpc.Start()
		if err != nil {
			panic(err)
		}
	}()

	return nil
}