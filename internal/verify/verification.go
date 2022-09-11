package verify

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const verificationChars = "1234567890ABCDEFGHIJKLMNOPQRSTUVXYZ"

func CreateVerificationCode(hash string) (string, error) {
	jsonFile, err := os.OpenFile("internal/mock-data/verification-codes.json", os.O_RDWR, os.ModeAppend)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var verificationMap = make(map[string]string)

	json.Unmarshal(byteValue, &verificationMap)

	var isValidCode = false
	var code string

	for !isValidCode {
		code, err = createCode()

		if err != nil {
			fmt.Println(err)
		}

		if _, isPresent := verificationMap[code]; !isPresent {
			verificationMap[code] = hash
			isValidCode = true
			json.Marshal(&verificationMap)

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
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var verificationMap = make(map[string]string)

	json.Unmarshal(byteValue, &verificationMap)

}
