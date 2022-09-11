package verification

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const verificationChars = "1234567890ABCDEFGHIJKLMNOPQRSTUVXYZ"

func CreateVerificationCode(hash string) (string, error) {
	jsonFile, err := os.OpenFile("internal/mock-data/verification-codes.json", os.O_RDWR, os.ModeAppend)

	if err != nil {
		return "Fallo al contactar base de datos", err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return "Fallo al leer base de datos", err
	}

	var verificationMap = make(map[string]string)

	err = json.Unmarshal(byteValue, &verificationMap)

	if err != nil {
		return "Fallo en el formateo de datos", err
	}

	var isValidCode = false
	var code string

	for !isValidCode {
		code, err = createCode()

		if err != nil {
			return "Fallo al crear código alfanumérico", err
		}

		if _, isPresent := verificationMap[code]; !isPresent {
			verificationMap[code] = hash
			isValidCode = true

			data, err := json.Marshal(&verificationMap)
			if err != nil {
				return "Fallo al formatear datos de base de datos", err
			}

			_, err = jsonFile.WriteAt(data, 0)
			if err != nil {
				return "Fallo al escribir en base de datos", err
			}

		}

	}
	return code, nil
}

func createCode() (string, error) {
	buffer := make([]byte, 6)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	verificationCharsLength := len(verificationChars)
	for i := 0; i < 6; i++ {
		buffer[i] = verificationChars[int(buffer[i])%verificationCharsLength]
	}

	return string(buffer), nil
}

func GetHashCode(verificationCode string) (string, error) {
	jsonFile, err := os.OpenFile("internal/mock-data/verification-codes.json", os.O_RDWR, os.ModeAppend)

	if err != nil {
		return "Fallo al contactar base de datos", err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return "Fallo al leer base de datos", err
	}

	var verificationMap = make(map[string]string)

	err = json.Unmarshal(byteValue, &verificationMap)

	if err != nil {
		return "Fallo en el formateo de datos", err
	}

	hashValue, isPresent := verificationMap[verificationCode]
	if !isPresent {
		return "Código de verificación no encontrado en base de datos", errors.New("Código de verificación no encontrado en base de datos")
	}
	return hashValue, nil

}
