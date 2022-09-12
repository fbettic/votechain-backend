package verification

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fbettic/votechain-backend/pkg/dto"
)

func CreateHash(user dto.User) (string, error) {
	data, err := json.Marshal(user)

	if err != nil {
		return "", errors.New("500 - No se pudieron formatear los datos: "+err.Error())
	}

	return string(crypto.Keccak256(data)), nil
}
