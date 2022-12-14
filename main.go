package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
  "github.com/gorilla/mux"
	"github.com/joho/godotenv"
  "strings"
)

type Input struct {
	OperationType string `json:"operation_type"`
	X       int   `json:"x"`
	Y      int    `json:"y"`
}

type Output struct {
	SlackUsername string `json:"slackUsername"`
	Result       int   `json:"result"`
	OperationType  string    `json:"operation_type"`
}


func Calculate(w http.ResponseWriter, r *http.Request) {

fmt.Println("Calculating.... \n")
var input Input
json.NewDecoder(r.Body).Decode(&input)

// var OperationType = struct {
//   Addition string
//   Subtraction  string
//   Multiplication  string
// }{
//   Addition: "addition",
//   Subtraction:  "subtraction",
//   Multiplication:     "multiplication",
// }

var z int
var OperationTypeReturn string
OperationTypeLowerCase := strings.ToLower(input.OperationType)
if OperationTypeLowerCase == "addition" {
z = input.X + input.Y
OperationTypeReturn = "addition"
}else if  OperationTypeLowerCase == "subtraction"{
z = input.X - input.Y
OperationTypeReturn = "subtraction"
}else if  OperationTypeLowerCase == "multiplication"{
z = input.X * input.Y
OperationTypeReturn = "multiplication"
}else {
  z = 0
  OperationTypeReturn = "Not Supported"
}


fmt.Println(z)
fmt.Println(input)

var output Output

output.SlackUsername = "Scoba"
output.Result = z
output.OperationType = OperationTypeReturn

//set header tag
w.Header().Set("Content-Type", "application/json")
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")

  json.NewEncoder(w).Encode(output)
	w.WriteHeader(http.StatusOK)
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

  r := mux.NewRouter()
  r.HandleFunc("/", Calculate).Methods("POST")
  http.Handle("/", r)

	fmt.Print("Listening on :" + port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
