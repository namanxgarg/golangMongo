package controllers

import (
	"context"
	"encoding/json"

	"golangMongo/models"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// Helper function to send JSON response
func sendJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// GetUser retrieves a user by ID
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Validate ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Connect to the collection
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find user by ID
	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	sendJSONResponse(w, http.StatusOK, user)
}

// CreateUser creates a new user
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	// Decode the request body into the user struct
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		return
	}

	// Assign a new ObjectId to the user
	user.Id = primitive.NewObjectID()

	// Connect to the collection
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insert the user into the database
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
		return
	}

	sendJSONResponse(w, http.StatusCreated, user)
}

// DeleteUser deletes a user by ID
func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Validate ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "Invalid user ID"})
		return
	}

	// Connect to the collection
	collection := uc.client.Database("mongo-golang").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Remove the user from the database
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		sendJSONResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	sendJSONResponse(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}
