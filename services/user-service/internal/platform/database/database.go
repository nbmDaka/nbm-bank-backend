package database

import "context"


type DB interface {

	Close()

	Ping(ctx context.Context) error

}