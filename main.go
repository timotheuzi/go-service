package main

import (
	"go-service/controllers"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	controller := controllers.NewUserController()
	router.GET("/", index)
	router.POST("/users", controller.PostUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.PATCH("/users/:id", controller.PatchUserByID)
	router.DELETE("/users/:id", controller.DeleteUserByID)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe(":"+port, router))
}

func index(respWr http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	respWr.WriteHeader(http.StatusOK)
	respWr.Header().Set("Content-Type", "text/html; charset=utf-8")
	respWr.Write([]byte("User Database API"))
}
