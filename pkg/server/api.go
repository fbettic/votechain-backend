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
	r.HandleFunc("/votechainapi/login", a.login).Methods("POST")

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
	a.repository.RegisterVote(vote)

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vote)
}



func (a *api) login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var user option.User
	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")

	type Unauthorized struct {
		Message string `json:"message"`
	}

	type Authorized struct {
		AccessToken string `json:"access_token"`
	}

	if !a.repository.Login(user) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Unauthorized{Message: "cuit y/o contraseña incorrectos"})
	}
	
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Authorized{AccessToken: "estOesTRuCHiSiMo.Mano6785487487.79"})

}
