package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-simple-crud/domain"
	"user-simple-crud/usecase"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(router *mux.Router, uc usecase.UserUsecase) {
	handler := &UserHandler{
		usecase: uc,
	}

	router.HandleFunc("/users", handler.GetAll).Methods("GET")
	router.HandleFunc("/users/{id}", handler.GetByID).Methods("GET")
	router.HandleFunc("/users", handler.Create).Methods("POST")
	router.HandleFunc("/users/{id}", handler.Update).Methods("PUT")
	router.HandleFunc("/users/{id}", handler.Delete).Methods("DELETE")
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.usecase.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)

	err := h.usecase.Create(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user domain.User
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = id

	err = h.usecase.Update(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.usecase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
