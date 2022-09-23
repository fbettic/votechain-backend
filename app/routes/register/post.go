package register

import (
	"encoding/json"
	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func Post(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)

		var user dto.User
		json.NewDecoder(r.Body).Decode(&user)

		user.HasVoted = false

		w.Header().Set("Content-Type", "application/json")

		err := srv.Register(user)
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			w.WriteHeader(errDto.Status)
			json.NewEncoder(w).Encode(errDto)
		}else{
			w.WriteHeader(http.StatusOK)

			result := struct {
				Result string `json:"result"`
			}{Result: "User registered correctly. You can login now"}

			json.NewEncoder(w).Encode(result)
		}
	}
}