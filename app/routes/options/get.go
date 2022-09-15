package options

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"
	"github.com/fbettic/votechain-backend/internal/verification"
)

func Get(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		token := strings.Split(r.Header["Authorization"][0], " ")[1]
		fmt.Println(token)
		w.Header().Set("Content-Type", "application/json")

		if !verification.ValidToken(token) {
			err := errorHandler.GetErrorDto(errors.New("401 - Invalid login token"))
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}
		options, _ := srv.FetchOptions()
		json.NewEncoder(w).Encode(options)
	}
}
