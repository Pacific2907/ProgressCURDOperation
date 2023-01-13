package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

type Employdb struct {
	Eid        string `json:"eid,omitempty"`
	Ename      string `json:"name,omitempty"`
	Contact_No string `json:"contact_no,omitempty"`
}

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Nightpacific@07@tcp(localhost:3306)/Employdb")
	if err != nil {
		log.Fatal("err")
	}
	insert, err := db.Query("INSERT INTO Employee(Eid, Ename, Contact_NO) VALUES (103,'Govind',98765);")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	return db

}
func getEmployee(w http.ResponseWriter, r *http.Request) {
	db = getMySQLDB()
	defer db.Close()
	ee := []Employdb{}
	e := Employdb{}
	rows, err := db.Query("Select * from Employee")
	if err != nil {
		fmt.Fprintf(w, ""+err.Error())
	} else {
		for rows.Next() {
			rows.Scan(&e.Eid, &e.Ename, &e.Contact_No)
			ee = append(ee, e)
		}
	}
	json.NewEncoder(w).Encode(ee)
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST request")
}
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PUT request")
}
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DELETE request")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/employee", getEmployee).Methods("GET")
	r.HandleFunc("/employee", addEmployee).Methods("POST")
	r.HandleFunc("/employee/{Eid}", updateEmployee).Methods("PUT")
	r.HandleFunc("/employee/{Eid}", deleteEmployee).Methods("DELETE")
	http.ListenAndServe(":8081", r)
}
