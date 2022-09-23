package votechain

import (
	"errors"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	sampledata "github.com/fbettic/votechain-backend/internal/mock-data"
	"github.com/fbettic/votechain-backend/internal/verification"
	"github.com/fbettic/votechain-backend/pkg/dto"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

func (r *Broker) Login(userLogin dto.Login) (*dto.User, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	users := sampledata.Users
	user, found := users[userLogin.Cuit]
	if found {
		if userLogin.Password == user.Password{
			return user, nil
		}
	}

	return nil, errors.New("401 - Wrong credentials")
}

func (r *Broker) Register(newUser dto.User) error {
	users := sampledata.Users
	_, found := users[newUser.Cuit]
	if found {
		return errors.New("409 - Cuit already registered")
	}

	err := passwordvalidator.Validate(newUser.Password, 60)
	if err != nil {
		fmt.Println(err)
		return err
	}

	token, err := verification.GenerateJWT(newUser.Name, newUser.Cuit)
	if err != nil {
		return err
	}

	newUser.AccessToken = token

	users[newUser.Cuit] = &newUser

	return nil
}

func ReloadUsers() (error){
	jsonFile, err := os.OpenFile("internal/mock-data/users.json", os.O_RDWR | os.O_CREATE, os.ModeAppend)
	if err != nil {
		return errors.New("500 - Error opening user db")
	}

	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return errors.New("500 - Fallo al leer base de datos: "+err.Error())
	}

	var usersMap = make(map[int]*dto.User)

	err = json.Unmarshal(byteValue, &usersMap)

	if err != nil {
		return errors.New("500 - Fallo en el formateo de datos: "+err.Error())
	}

	sUsers := sampledata.Users

	for _, user := range sUsers {
		_, found := usersMap[user.Cuit]
		if !found {
			usersMap[user.Cuit] = user
		}
		delete(sUsers, user.Cuit)
	}

	data, err := json.Marshal(&usersMap)
	if err != nil {
		return errors.New("500 - Fallo al formatear datos de base de datos: "+err.Error())
	}

	_, err = jsonFile.WriteAt(data, 0)
	if err != nil {
		return errors.New("500 - Fallo al escribir en base de datos: "+err.Error())
	}

	sampledata.Users = usersMap

	return nil
}