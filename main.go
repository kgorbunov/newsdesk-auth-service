package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"errors"
)

type Response struct {
	Status	string `json:"status"`
	Data	User	`json:"data"`
	Message	string	`json:"message"`
}

type User struct {
	Uid	string			`json:"uid"`
	Name	string			`json:"name"`
	Password string			`json:"password"`
	Roles	map[string]string	`json:"roles"`
}

func SearchUser(username, password string, collection []User, callback func(User)) error {
	for i:= 0; i < len(collection); i++ {
		if collection[i].Name == username && collection[i].Password == password {
			callback(collection[i])
			return nil
		}
	}
	return errors.New("username or password invalid")
}

func handler(w http.ResponseWriter, r *http.Request){
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	roles := map[string]string{
		"2": "authenticated user",
		"32": "bz_manager",
		"38": "bz_data_calendar",
	}

	userCollection := []User{
		User{
			Uid: "90759",
			Name: "kgorbunov",
			Password: "secret",
			Roles:roles,
		},
		User{
			Uid: "90759",
			Name: "dsemenov",
			Password: "secret",
			Roles:roles,
		},
		User{
			Uid: "90759",
			Name: "demo",
			Password: "demo",
			Roles:roles,
		},
	}

	CreateResponseCallback := func(user User) {
		response, err := json.Marshal(Response{
			Status: "ok",
			Data: user,
			Message:"access granted",
		})
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		fmt.Fprint(w,string(response))
	}

	err := SearchUser(username, password, userCollection, CreateResponseCallback)
	if err != nil {
		fmt.Fprintf(w, `{"status":"error","data":"","message": "%s"}`, err)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:8001", nil)
}