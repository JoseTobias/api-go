package routes

import (
	"net/http"
	"encoding/json"
)

type responseHello struct {
	Message string `json:"message"`
	Name string `json:"name"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
    
	w.Header().Set("Content-Type", "application/json")

	name := ""
	message := "Hi"
	if ok {
		name = keys[0]
		message += " " + name
	}

	res := responseHello{message + " welcome to the api!", name}

	resJson, _ := json.Marshal(res)

	w.Write(resJson)
}