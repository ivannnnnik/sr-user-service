package handler

import (
	"encoding/json"
	"net/http"
	"plan/old/internal/model"
	"plan/old/internal/usecase"
	"strconv"
)

type UserHandler struct {
	users *usecase.User
}

func NewUserHandler(users *usecase.User) *UserHandler {
	return &UserHandler{
		users: users,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed pars your params", http.StatusBadRequest)
		return
	}

	user, err := h.users.Create(req.Name, req.Email)
	if err != nil {
		message := "Failed create user" + err.Error()
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid param id", http.StatusBadRequest)
		return
	}

	user, err := h.users.GetByID(idVal)
	if err != nil {
		message := "Failed get user" + err.Error()
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid param id", http.StatusBadRequest)
		return
	}

	err = h.users.Delete(idVal)
	if err != nil {
		message := "Failed delete user" + err.Error()
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid param id", http.StatusBadRequest)
		return
	}

	var userU model.UpdateUser
	if err := json.NewDecoder(r.Body).Decode(&userU); err != nil {
		http.Error(w, "Failed pars your params", http.StatusBadRequest)
		return
	}

	user, err := h.users.Update(idVal, &userU)
	if err != nil {
		message := "Failed update user: " + err.Error()
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
