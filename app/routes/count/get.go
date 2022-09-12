package count

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
	"github.com/gorilla/mux"
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
		vars := mux.Vars(r)
		option := new(dto.Option)
		option = nil
		options, _ := srv.FetchOptions()
		for _, o := range options{
			if o.ID == vars["id"]{
				option = o
			}
		}
		if option == nil{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(errors.New("404 - Option not found")))
			return
		}
		optionCounted, err := srv.FetchOptionCount(option)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errorHandler.GetErrorDto(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(optionCounted)
	}
}