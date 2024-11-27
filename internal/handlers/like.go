package handlers

import (
	"blog/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type LikeHandler struct {
	LikeService *services.LikeService
}

func NewLikeHandler(likeService *services.LikeService) *LikeHandler {
	return &LikeHandler{LikeService: likeService}
}

// AddLikeHandler handles the HTTP POST request to add a like to a post.
//
// It expects post ID as a path parameter and user ID as a header parameter.
// If the like is added successfully, it returns a 201 Created response.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid post ID or user ID.
// 500 Internal Server Error: Failed to add like.
func (h *LikeHandler) AddLikeHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := mux.Vars(r)["postID"]
	userIDStr := r.Header.Get("UserID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.LikeService.AddLike(uint(postID), uint(userID)); err != nil {
		http.Error(w, "Failed to add like", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// RemoveLikeHandler handles the HTTP DELETE request to remove a like from a post.
//
// It expects post ID as a path parameter and user ID as a header parameter.
// If the like is removed successfully, it returns a 204 No Content response.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid post ID or user ID.
// 500 Internal Server Error: Failed to remove like.
func (h *LikeHandler) RemoveLikeHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := mux.Vars(r)["postID"]
	userIDStr := r.Header.Get("UserID")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.LikeService.RemoveLike(uint(postID), uint(userID)); err != nil {
		http.Error(w, "Failed to remove like", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GetLikesCounterHandler handles the HTTP GET request to retrieve the number of likes for a post.
//
// It expects post ID as a path parameter.
// If the likes count is retrieved successfully, it returns a JSON response with a single key-value pair,
// where the key is "likes" and the value is the number of likes.
// Otherwise, it returns one of the following errors:
// 400 Bad Request: Invalid post ID.
// 500 Internal Server Error: Failed to get likes count.
func (h *LikeHandler) GetLikesCounterHandler(w http.ResponseWriter, r *http.Request) {
	postIDStr := mux.Vars(r)["postID"]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	count, err := h.LikeService.GetLikesCount(uint(postID))
	if err != nil {
		http.Error(w, "Failed to get likes count", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int64{"likes": count})
}
