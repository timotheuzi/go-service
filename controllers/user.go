package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-service/models"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	mongo *mongo.Client
}

func NewUserController() *UserController {
	return &UserController{getMongoClient()}
}

func (controller *UserController) PostUser(respWr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}
	printErrIfAny(json.NewDecoder(req.Body).Decode(&u))

	collection := controller.mongo.Database("goServer").Collection("users")
	res, err := collection.InsertOne(context.Background(), models.User{
		Email: u.Email,
		ID:    u.ID,
		Name:  u.Name})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)

	uj, err := json.Marshal(u)
	printErrIfAny(err)

	respWr.WriteHeader(http.StatusOK)
	respWr.Header().Set("Content-Type", "application/json")
	respWr.Write(uj)
}

func (controller *UserController) GetUserByID(respWr http.ResponseWriter, req *http.Request, p httprouter.Params) {
	user := models.User{}
	collection := controller.mongo.Database("goServer").Collection("users")
	printErrIfAny(collection.FindOne(context.TODO(), bson.M{"id": p.ByName("id")}).Decode(&user))

	userJSON, err := json.Marshal(user)
	printErrIfAny(err)

	respWr.WriteHeader(http.StatusOK)
	respWr.Header().Set("Content-Type", "application/json")
	respWr.Write(userJSON)
}

func (controller *UserController) PatchUserByID(respWr http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{}
	printErrIfAny(json.NewDecoder(req.Body).Decode(&u))
	email := u.Email
	name := u.Name

	filter := bson.M{"id": p.ByName("id")}

	collection := controller.mongo.Database("goServer").Collection("users")
	if email != "" {
		update := bson.M{"$set": bson.M{"email": email}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		printErrIfAny(err)
	}

	if name != "" {
		update := bson.M{"$set": bson.M{"name": name}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		printErrIfAny(err)
	}

	respWr.WriteHeader(http.StatusNoContent)
}

func (controller *UserController) DeleteUserByID(respWr http.ResponseWriter, req *http.Request, p httprouter.Params) {
	collection := controller.mongo.Database("goServer").Collection("users")
	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": p.ByName("id")})
	printErrIfAny(err)

	respWr.WriteHeader(http.StatusNoContent)
}

func checkFatality(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func printErrIfAny(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
