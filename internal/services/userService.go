package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web-service/internal/dto"
)

type User interface {
	GetUsers() []dto.User
}

type user struct {
	user dto.User
}

func (u user) GetUsers() []dto.User {
	fmt.Println("service start")
	var userResponse dto.UserHttpResponse
	resp, err := http.Get("https://reqres.in/api/users?page=2&&per_page=2")

	fmt.Println(resp)
	if err != nil {
		log.Fatal("error connecting to external client", err.Error())
	}

	//body, err := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK {
		err1 := json.NewDecoder(resp.Body).Decode(&userResponse)
		if err1 != nil {
			log.Fatal("cannot unmarshal the json response")
		}
		fmt.Println("service if inside")

	}

	return userResponse.Users
}

func NewUsers() User {
	return &user{}
}
