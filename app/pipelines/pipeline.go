package pipelines

import (
	"github.com/fbettic/votechain-backend/app/routes/options"
	"github.com/fbettic/votechain-backend/app/routes/votes"
	"github.com/fbettic/votechain-backend/app/routes/count"
	"github.com/fbettic/votechain-backend/app/routes/login"
	"github.com/fbettic/votechain-backend/app/webserver"
	"github.com/gorilla/mux"
)

// BuildPipeline builds the HTTP pipeline
func BuildPipeline(srv webserver.Server, r *mux.Router) {

	r.HandleFunc("/votechainapi/options", options.Get(srv)).Methods("GET")
	r.HandleFunc("/votechainapi/vote", votes.Post(srv)).Methods("POST")
	r.HandleFunc("/votechainapi/count/{id}",count.Get(srv)).Methods("GET")
	r.HandleFunc("/votechainapi/login",login.Post(srv)).Methods("POST")
}
