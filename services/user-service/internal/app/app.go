package app


import (

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/config"

	"github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/platform/database"

	grpcserver "github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/server/grpc"
	grpcuser "github.com/nbmDaka/nbm-bank-backend/services/user-service/internal/user/delivery/grpc"

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

	userHandler := grpcuser.NewHandler()

	return &App{
		db: db,
		grpc: grpcserver.NewServer(cfg.App.Port, userHandler),
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