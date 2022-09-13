package votes

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"
	"github.com/fbettic/votechain-backend/internal/verification"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func Post(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		var vote dto.Vote
		token := strings.Split(r.Header["Authorization"][0], " ")[1]

		if !verification.ValidToken(token) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(errors.New("401 - Invalid login token")))
			return
		}

		json.NewDecoder(r.Body).Decode(&vote)
		vote.AccessToken = token
		
		validationCode, err := srv.RegisterVote(vote)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(err))
			return
		}
		
		validationJson = { "code" : validationCode }

		w.WriteHeader(http.StatusOK)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(validationJson)
	}
}
