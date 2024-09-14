package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"user-simple-crud/delivery"
	"user-simple-crud/infra"
	"user-simple-crud/repository"
	"user-simple-crud/usecase"
)

func main() {
	db := infra.InitDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	router := mux.NewRouter()
	delivery.NewUserHandler(router, userUsecase)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
