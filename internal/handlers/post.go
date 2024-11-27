package handlers

import (
	"blog/internal/models"
	"blog/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	PostService *services.PostService
}

func NewPostHandler(postService *services.PostService) *PostHandler {
	return &PostHandler{PostService: postService}
}

// CreatePostHandler handles the HTTP POST request to create a new post.
//
// It expects a JSON body with the post details. If the post is created successfully,
// it returns a 201 Created response with the created post in JSON format.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: If the request body is invalid.
// 500 Internal Server Error: If the server fails to create the post.
func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.PostService.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) GetPostsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userIDStr := mux.Vars(r)["userID"]
	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	posts, err := h.PostService.GetPostsByUserID(uint(userID))
	if err != nil {
		http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(posts)
}

// DeletePostHandler handles the HTTP DELETE request to delete a post.
//
// It expects a post ID as a path parameter.
// If the post is deleted successfully, it returns a 204 No Content response.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid post ID.
// 500 Internal Server Error: Failed to delete post.

func (h *PostHandler) DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := mux.Vars(r)["postID"]
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	if err := h.PostService.DeletePost(uint(postID)); err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
