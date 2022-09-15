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
		/*
		token := strings.Split(r.Header["Authorization"][0], " ")[1]

		if !verification.ValidToken(token) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(errors.New("401 - Invalid login token")))
			return
		}
*/
w.Header().Set("Content-Type", "application/json")
		code := vars["code"]
		option, err := srv.GetVote(code)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(err))
			return
		}


		json.NewEncoder(w).Encode(option)
	}
}