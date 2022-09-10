package count

import (
	"encoding/json"
	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	"github.com/fbettic/votechain-backend/pkg/dto"
	"github.com/gorilla/mux"
)

func Get(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
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
			error := &dto.ErrorMessage{
						Status: 404,
						Message: "Option not found",
					}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(error)
			return
		}
		optionCounted, _ := srv.FetchOptionCount(option)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(optionCounted)
	}
}