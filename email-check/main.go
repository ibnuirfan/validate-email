package main


import (
//	"fmt"
    "encoding/json"
	"log"
    "net/http"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
//	"io/ioutil"

    "github.com/gorilla/mux"
)

type User struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"f_name,omitempty"`
    Lastname  string   `json:"l_name,omitempty"`
    Email     string   `json:"email,omitempty"`
}

var people []User
var db *sql.DB
var err error

func GetUserEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	result, err := db.Query("SELECT * FROM user WHERE user_id =?", params["id"])
	if err != nil {
		panic(err.Error())
		}

	defer result.Close()

	var user User

	for result.Next() {
		err := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			panic(err.Error())
		}
		people = append(people, user)
		}
	json.NewEncoder(w).Encode(people)
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT * FROM validate_email")
	if err != nil {
		panic(err.Error())
		}

		defer result.Close()

		for result.Next() {
			var user User
			err := result.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
			if err != nil {
				panic(err.Error())
			}
			people = append(people, user)
		}
		json.NewEncoder(w).Encode(people)
}

//func CreateUserEndpoint(w http.ResponseWriter, req *http.Request) {
//    params := mux.Vars(req)
//    var user User
//    _ = json.NewDecoder(req.Body).Decode(&user)
//    user.ID = params["id"]
//    people = append(people, user)
//    json.NewEncoder(w).Encode(people)
//}
//
//func DeleteUserEndpoint(w http.ResponseWriter, req *http.Request) {
//    params := mux.Vars(req)
//    for index, item := range people {
//        if item.ID == params["id"] {
//            people = append(people[:index], people[index+1:]...)
//            break
//        }
//    }
//    json.NewEncoder(w).Encode(people)
//}

func main() {

	db, err := sql.Open("mysql", "root:Cloud#9@tcp(127.0.0.1:3306)/validate_email")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

    router := mux.NewRouter()

//    people = append(people, User{ID: "1", Firstname: "Chris", Lastname: "Brown", Email: "chrisbreezy@gmail.com"})
//    people = append(people, User{ID: "2", Firstname: "Chris", Lastname: "Ramsey", Email: "rampup@gmail.com"})
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/people/{id}", GetUserEndpoint).Methods("GET")
//    router.HandleFunc("/people/{id}", CreateUserEndpoint).Methods("POST")
//    router.HandleFunc("/people/{id}", DeleteUserEndpoint).Methods("DELETE")

    log.Fatal(http.ListenAndServe(":9090", router))
}