package votes

import (
	"encoding/json"

	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"

	"github.com/gorilla/mux"
)

func Get(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		vars := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		
		code := vars["code"]
		option, err := srv.GetVote(code)
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			w.WriteHeader(errDto.Status)
			json.NewEncoder(w).Encode(errDto)
			return
		}


		json.NewEncoder(w).Encode(option)
	}
}