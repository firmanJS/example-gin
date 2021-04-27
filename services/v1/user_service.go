package v1

import (
	"context"
	helper "firmanjs/example-gin/helpers"
	"firmanjs/example-gin/repository/models"
	requestModel "firmanjs/example-gin/repository/request/v1"
	responseModel "firmanjs/example-gin/repository/response/v1"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	uuid "github.com/google/uuid"
)

var collection *mongo.Collection

type UserServices struct {
	User models.Users
}

func UsersCollections(m *mongo.Database) {
	collection = m.Collection("users")
}

func GetAll() (int, string, string, interface{}) {
	filter := bson.M{"base.deletedby": ""}
	// userdb := us.User
	usersresponse := []responseModel.UserResponse{}
	//get all user from db
	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Printf("Error when getting all user %v\n", err)
		return http.StatusInternalServerError, "Error", "Someting wrong", nil
	}

	for cursor.Next(context.TODO()) {
		var user *models.Users
		cursor.Decode(&user)
		userRes := responseModel.UserResponse{
			ID:       user.ID,
			Username: user.Username,
		}

		usersresponse = append(usersresponse, userRes)
	}

	return http.StatusOK, "Succes", "Get all data sucessfully", usersresponse
}

func Create(id string, req *requestModel.CreateUserReq) (int, string, string, interface{}) {

	if req == nil {
		return http.StatusBadRequest, "Error", "Someting wrong", nil
	}

	newUser := models.Users{
		ID:       uuid.New().String(),
		Username: req.Username,
		Password: req.Password,
		Base: models.Base{
			CreatedAt: time.Now(),
			CreatedBy: id,
		},
	}

	_, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error when inserting new users : %v\n", err)
		return http.StatusInternalServerError, "Error", err.Error(), nil
	}

	return http.StatusCreated, "Succes", "Succesfull create user", newUser
}

func Update(id string, user *requestModel.EditUserReq) map[string]interface{} {

	filter := bson.M{"$and": []bson.M{
		{"id": user.ID},
		{"base.deletedby": ""},
	}}

	newData := bson.M{
		"$set": bson.M{
			"username":         user.Username,
			"base.updatedtime": time.Now(),
			"base.updatedby":   id,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, newData)

	if err != nil {
		log.Printf("Error when updating users : %v\n", err)
		response := helper.Message(http.StatusInternalServerError, "Someting wrong")
		response["data"] = nil
		return response
	}

	if result.MatchedCount == 0 {
		response := helper.Message(http.StatusNotFound, "Not found Document")
		response["data"] = nil
		return response
	}

	reponse := helper.Message(http.StatusOK, "Succesfull Edit user")
	reponse["data"] = newData
	return reponse
}

func GetByParam(id string) (int, string, string, interface{}) {
	filter := bson.M{"$and": []bson.M{
		{"id": id},
	}}

	user := models.Users{}

	err := collection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return http.StatusNotFound, "Not found document", err.Error(), nil
		}

		log.Printf("Error when get users : %v\n", err)
		return http.StatusInternalServerError, "Error", err.Error(), nil
	}

	userResponse := responseModel.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return http.StatusOK, "Succes", "Get data by param sucessfully", userResponse
}

func DeleteByID(id string) (int, string, string, interface{}) {
	filter := bson.M{"$and": []bson.M{
		{"id": id},
	}}

	newData := bson.M{
		"$set": bson.M{
			"base.deletedtime": time.Now(),
			"base.deletedby":   id,
		},
	}

	result, err := collection.UpdateOne(context.TODO(), filter, newData)

	if err != nil {
		log.Printf("Error when delete users : %v\n", err)
		return http.StatusInternalServerError, "Error", err.Error(), nil
	}

	if result.MatchedCount == 0 {
		return http.StatusNotFound, "Error", "Not found Document", nil
	}

	return http.StatusNotFound, "Success", "Delete Uses Succesfully", result
}
