package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
)


func main() {
	r := mux.NewRouter()

	r.Handle("/find", getMsg()).Methods("GET")

	appPort := fmt.Sprintf(":9001")
	log.Println("Starting service on http://localhost", appPort)
	err := http.ListenAndServe(appPort, r)
	log.Println(err)
}

func getMsg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Yayayyyy!!! you found me..."
		jsonResponse, err := json.Marshal(msg)
		if err != nil {
			fmt.Sprintf("error marshalling response: %s", err)
			return
		}
		w.Write(jsonResponse)
	}
}


