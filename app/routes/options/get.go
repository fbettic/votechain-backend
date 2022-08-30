package options

import (
	"encoding/json"
	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
)

func Get(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		options, _ := srv.FetchOptions()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(options)
	}
}
