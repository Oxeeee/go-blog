package main

import (
	"blog/config"
	"blog/internal/handlers"
	"blog/internal/models"
	"blog/internal/repositories"
	"blog/internal/services"
	"blog/pkg/db"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from file
	config.LoadConfig()

	// Connect to the database
	database, err := db.Connect(config.AppConfig.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database, %s", err)
	}

	// Migrate the database
	if err := database.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}); err != nil {
		log.Fatalf("Failed to migrate database, %s", err)
	}

	log.Println("Connected to database")

	// Create a new router
	r := mux.NewRouter()

	// Add a health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is running!"))
	}).Methods("GET")

	// Create a user handler
	userRepo := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/verify", userHandler.VerifyEmail).Methods("POST")

	// Create a post handler
	postRepo := repositories.NewPostRepository(database)
	postService := services.NewPostService(postRepo)
	postHandler := handlers.NewPostHandler(postService)
	r.HandleFunc("/posts/{userID}", postHandler.GetPostsByUserIDHandler).Methods("GET")
	r.HandleFunc("/posts", postHandler.CreatePostHandler).Methods("POST")
	r.HandleFunc("/posts/{postID}", postHandler.DeletePostHandler).Methods("DELETE")

	// Create a like handler
	likeRepo := repositories.NewLikeRepository(database)
	likeService := services.NewLikeService(likeRepo)
	likeHandler := handlers.NewLikeHandler(likeService)
	r.HandleFunc("/posts/{postID}/like", likeHandler.AddLikeHandler).Methods("POST")
	r.HandleFunc("/posts/{postID}/like", likeHandler.RemoveLikeHandler).Methods("DELETE")
	r.HandleFunc("/posts/{postID}/likes", likeHandler.GetLikesCounterHandler).Methods("GET")

	// Add a login endpoint
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	log.Println("Server is running on port " + config.AppConfig.Server.Port)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start server, %s", err)
	}

}
