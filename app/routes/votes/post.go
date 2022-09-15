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
		w.Header().Set("Content-Type", "application/json")

		if !verification.ValidToken(token) {
			err := errorHandler.GetErrorDto(errors.New("401 - Invalid login token"))
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}

		json.NewDecoder(r.Body).Decode(&vote)
		vote.AccessToken = token

		validationCode, err := srv.RegisterVote(vote)
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			w.WriteHeader(errDto.Status)
			json.NewEncoder(w).Encode(errDto)
			return
		}

		validationJson := struct {
			Code string `json:"code"`
		}{Code: validationCode}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(validationJson)
	}
}
