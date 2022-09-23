package login

import (
	"encoding/json"
	"net/http"

	"github.com/fbettic/votechain-backend/app/middleware"
	"github.com/fbettic/votechain-backend/app/webserver"
	errorHandler "github.com/fbettic/votechain-backend/internal/error-handling"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func Post(srv webserver.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		middleware.EnableCors(&w)
		var user dto.Login
		json.NewDecoder(r.Body).Decode(&user)
		transaction, err := srv.Login(user)

		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			errDto := errorHandler.GetErrorDto(err)
			if errDto.Status == 404{
				w.WriteHeader(http.StatusUnauthorized)
				errDto.Message = "Usuario y/o contraseña incorrectos"
				errDto.Status = http.StatusUnauthorized
				json.NewEncoder(w).Encode(errDto)
			}else{
				w.WriteHeader(errDto.Status)
				json.NewEncoder(w).Encode(errDto)
			}
		}else{
			t:=*transaction
			t.Password=""
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(t)
		}
	}
}
