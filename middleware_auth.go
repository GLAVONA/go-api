package main

import (
	"fmt"
	"net/http"

	"github.com/glavona/go-api/internal/auth"
	"github.com/glavona/go-api/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		handler(w, r, user)
	}
}
