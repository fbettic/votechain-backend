package handError

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/fbettic/votechain-backend/pkg/dto"
)

func GetErrorDto(err error) dto.ErrorMessage{
	msj := err.Error()

	if !unicode.IsDigit(rune(msj[0])) || rune(msj[5]) != '-'{
		return dto.ErrorMessage{Status: 500, Message: msj}
	}

	str := strings.Split(msj," - ")

	status, errs := strconv.Atoi(str[0])
	if errs != nil{
		return dto.ErrorMessage{Status: 500, Message: "Error en la conversión a entero: "+errs.Error()}
	}

	return dto.ErrorMessage{Status: status, Message: str[1]}
}