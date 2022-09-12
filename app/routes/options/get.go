package options

import (
	"encoding/json"
	"net/http"
	"strings"
	"errors"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"
	"github.com/fbettic/votechain-backend/internal/verification"
)

func Get(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		token := strings.Split(r.Header["Authorization"][0], " ")[1]

		if !verification.ValidToken(token) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(errors.New("401 - Invalid login token")))
			return
		}
		options, _ := srv.FetchOptions()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(options)
	}
}
