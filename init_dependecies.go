package main

import (
	"github.com/LucasKeley/CRUD_Go/src/controller"
	"github.com/LucasKeley/CRUD_Go/src/model/repository"
	"github.com/LucasKeley/CRUD_Go/src/model/service"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func initDependecies(database *mongo.Database) controller.UserControllerInterface {
	//Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
