package verification

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	//"github.com/fbettic/votechain-backend/pkg/dto"
	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateHash(info any) (string, error) {
	data, err := json.Marshal(info)

	
	if err != nil {
		return "", errors.New("500 - No se pudieron formatear los datos: "+err.Error())
	}

	return string(crypto.Keccak256Hash(data).String()), nil
}

func GenerateJWT(name string, cuit int) (string, error) {
	var mySigningKey = []byte("V0T3CH41N")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["name"] = name
	claims["cuit"] = cuit

	tokenString, err := token.SignedString(mySigningKey)



	if err != nil {
		fmt.Println(err)
		return "", err
	}
	
	return tokenString, nil
}