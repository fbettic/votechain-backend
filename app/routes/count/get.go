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
		w.Header().Set("Content-Type", "application/json")

		if !verification.ValidToken(token) {
			err := errorHandler.GetErrorDto(errors.New("401 - Invalid login token"))
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}
		vars := mux.Vars(r)
		option := new(dto.Option)
		option = nil
		options, err := srv.FetchOptions()
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			w.WriteHeader(errDto.Status)
			json.NewEncoder(w).Encode(errDto)
			return
		}
		for _, o := range options{
			if o.ID == vars["id"]{
				option = o
				break
			}
		}
		if option == nil{
			err := errorHandler.GetErrorDto(errors.New("404 - Option not found"))
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}
		optionCounted, err := srv.FetchOptionCount(option)
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			w.WriteHeader(errDto.Status)
			json.NewEncoder(w).Encode(errDto)
			return
		}

		json.NewEncoder(w).Encode(optionCounted)
	}
}