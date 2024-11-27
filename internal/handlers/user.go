package handlers

import (
	"blog/internal/models"
	"blog/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

// RegisterUser handles the HTTP POST request to register a new user.
//
// It expects a JSON request body that matches the models.User struct.
// If the user is registered successfully, it returns a 201 Created response
// with a JSON response body of the form {"message": "User registered successfully"}.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid request.
// 500 Internal Server Error: Failed to register user.
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.UserService.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("User registered successfully")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// LoginUser handles the HTTP POST request to login a user.
// It expects two JSON parameters: "email" and "password".
// If the login is successful, it returns a JSON response with a single key-value pair,
// where the key is "token" and the value is a JWT authentication token.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid request.
// 401 Unauthorized: User not found or invalid credentials.
func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.UserService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	resp := map[string]string{"token": token}
	json.NewEncoder(w).Encode(resp)
}

// VerifyEmail handles the HTTP POST request to verify a user's email address.
//
// It expects two JSON parameters: "email" and "token".
// If the verification is successful, it returns a 200 OK response.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid request.
// 401 Unauthorized: User not found or invalid verification code.
// 500 Internal Server Error: Failed to verify email.
func (h *UserHandler) VerifyEmail(w http.ResponseWriter, r *http.Request) {
	var verifyReq struct {
		Token string `json:"token"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&verifyReq); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.UserService.VerifyEmail(verifyReq.Email, verifyReq.Token)
	switch err {
	case services.ErrUNF:
		http.Error(w, "User not found", http.StatusNotFound)
		return
	case services.ErrIVC:
		http.Error(w, "Invalid verification code", http.StatusBadRequest)
		return
	case nil:
		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
