package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/LucasKeley/CRUD_Go/src/configuration/logger"
	"github.com/LucasKeley/CRUD_Go/src/configuration/rest_err"
	"github.com/LucasKeley/CRUD_Go/src/model"
	"github.com/LucasKeley/CRUD_Go/src/model/repository/entity"
	"github.com/LucasKeley/CRUD_Go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser repository", zap.String("Journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(UserEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with email %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := ("Error trying to find user by email")
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", UserEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*UserEntity), nil

}

func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser repository", zap.String("Journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}

	ObjectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: ObjectID}}
	err := collection.FindOne(context.Background(), filter).Decode(UserEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with id %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := ("Error trying to find user by email")
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByID repository executed successfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", UserEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*UserEntity), nil

}

func (ur *userRepository) FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUser repository", zap.String("Journey", "createUser"))

	collection_name := os.Getenv(MONGODB_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	UserEntity := &entity.UserEntity{}
	filter := bson.D{
		{Key: "email", Value: email},
		{Key: "password", Value: password},
	}
	err := collection.FindOne(context.Background(), filter).Decode(UserEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := ("User or Password invalid")
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailandPassword"))
			return nil, rest_err.NewForbiddenError(errorMessage)
		}

		errorMessage := ("Error trying to find user by email")
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailandPassword"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmailAndPassWord"),
		zap.String("email", email),
		zap.String("userId", UserEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*UserEntity), nil

}
