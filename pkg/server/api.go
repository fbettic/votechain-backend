package server

import (
	"encoding/json"
	"net/http"

	option "github.com/fbettic/votechain-backend/pkg"
	"github.com/gorilla/mux"
)

type api struct {
	router     http.Handler
	repository option.OptionRepository
}

type Server interface {
	Router() http.Handler
}

func New(repository option.OptionRepository) Server {
	a := &api{repository: repository}

	r := mux.NewRouter()

	r.HandleFunc("/votechainapi/options", a.fetchOptions).Methods("GET")
	r.HandleFunc("/votechainapi/vote", a.registerVote).Methods("POST")

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (a *api) fetchOptions(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	options, _ := a.repository.FetchOptions()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

func (a *api) registerVote(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var vote option.Vote
	json.NewDecoder(r.Body).Decode(&vote)
	transaction := a.repository.RegisterVote(vote)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
