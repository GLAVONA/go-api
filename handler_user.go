package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/glavona/go-api/internal/auth"
	"github.com/glavona/go-api/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
		return
	}

	if params.Name == "" {
		respondWithError(w, 400, "Please provide a valid name for the user")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now().UTC(), UpdatedAt: time.Now().UTC(), Name: params.Name})
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't create user: ", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprint("Auth err: ", err))
		return
	}

	user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Couldn't get user: ", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
