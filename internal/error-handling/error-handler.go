package handError

import (
	"strconv"
	"strings"

	"github.com/fbettic/votechain-backend/pkg/dto"
)

func GetErrorDto(err error) dto.ErrorMessage{
	str := strings.Split(err.Error()," - ")

	status, errs := strconv.Atoi(str[0])
	if errs != nil{
		return dto.ErrorMessage{Status: 500, Message: "Error en la conversión a entero: "+errs.Error()}
	}

	return dto.ErrorMessage{Status: status, Message: str[1]}
}