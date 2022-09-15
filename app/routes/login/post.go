package login

import (
	"encoding/json"
	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func Post(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		var user dto.Login
		json.NewDecoder(r.Body).Decode(&user)
		transaction, _ := srv.Login(user)

		w.Header().Set("Content-Type", "application/json")
		if transaction == nil {
			err := &dto.ErrorMessage{Status: http.StatusUnauthorized, Message: "Cuit y/o contraseña incorrectos"}
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(err)
		}else{
			t:=*transaction
			t.Password=""
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(t)
		}
	}
}
