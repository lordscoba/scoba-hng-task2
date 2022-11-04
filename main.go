package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Person struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func CallScoba(w http.ResponseWriter, r *http.Request) {

	// respresenting the data in string form
	// data := &person{ SlackUsername: "scoba", Backend: true, Age: 26, Bio:"I am a backend developer with golang Technology"}

	//representing data in another format
	var data Person
	data.SlackUsername = "Scoba is testing"
	data.Backend = true
	data.Age = 26
	data.Bio = "I am a backend developer with golang Technology"

	// use marshal func to convert to json
	//     out, err := json.Marshal(data)
	//     if err != nil {
	//         fmt.Println(err)
	//         return
	//     }
	//
	//   fmt.Println(string(out))

	// set headers content type
	w.Header().Add("Content-Type", "application/json")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

	json.NewEncoder(w).Encode(data)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	http.HandleFunc("/", CallScoba)
	fmt.Print("Listening on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
