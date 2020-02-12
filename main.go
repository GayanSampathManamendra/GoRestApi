package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start the main function....")
	addStudent()
	fmt.Println(std)
	// initialize the go router ...
	r := mux.NewRouter()

	// currently not connecting the db . Therfor moke the data...

	// route handerlers (End points ...)

	r.HandleFunc("/api/students", getStudents).Methods("GET")
	r.HandleFunc("/api/student/{id}", getStudent).Methods("GET")
	r.HandleFunc("/api/student", createStudent).Methods("POST")
	r.HandleFunc("/api/student/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/api/student/{id}", deleteStudent).Methods("DELETE")

	// server listen the http request to the particular port...
	log.Fatal(http.ListenAndServe(":8000", r))
}

/*
	golang is not an object oriented programming language .So class and objects can't be created ..
	Beacause of that In go , create a stuct insted of classes ..
*/

// create a student struct..
type student struct {
	stdId      string `json:"stdId"`
	stdName    string `json:"stdName"`
	stdAge     int    `json:"stdAge"`
	stdAddress string `json:"stdAddress"`

	// define an another stuct as a property...
	stdDpt *department `json:"stdDpt"`
}

// create a department stuct..

type department struct {
	dptId   string `json:"dptId"`
	dptName string `json:"dptName"`
}

// get all the students..
func getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This function for getting the all the students ...")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(std)
}

// get a specific student ..
func getStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This function for getting the particular student..")
	w.Header().Set("Content-Type", "application/json")
	parameters := mux.Vars(r) // get the all the parameters which are come from front end..

	for _, stud := range std {
		if stud.stdId == parameters["stdId"] {
			json.NewEncoder(w).Encode(stud)
			return
		}
	}
	json.NewEncoder(w).Encode(&student{})
}

// create a student ..
func createStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This function for creating a new student ..")

	w.Header().Set("Content-type", "application/json")
	var newStudent student
	_ = json.NewDecoder(r.Body).Decode(&newStudent)

	// to generate a id
	// using strconv.itoa funciton , converts the generated random number to a string ...
	newStudent.stdId = strconv.Itoa(rand.Intn(10000000))
	std = append(std, newStudent)
	json.NewEncoder(w).Encode(newStudent)
}

// update a  student record
func updateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This function for updating a student record")

	w.Header().Set("Content-type", "application/json")
	parameters := mux.Vars(r)

	for index, stud := range std {
		if stud.stdId == parameters["stdId"] {
			std = append(std[:index], std[index+1:]...)
			break
			var updateStudent student
			_ = json.NewDecoder(r.Body).Decode(&updateStudent)

			// to generate a id
			// using strconv.itoa funciton , converts the generated random number to a string ...
			updateStudent.stdId = strconv.Itoa(rand.Intn(10000000))
			std = append(std, updateStudent)
			json.NewEncoder(w).Encode(updateStudent)
			return
		}
	}
	json.NewEncoder(w).Encode(std)
}

// delete a student..
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This function for deleting a student ...")

	w.Header().Set("Content-type", "application/json")
	parameters := mux.Vars(r)

	for index, stud := range std {
		if stud.stdId == parameters["stdId"] {
			std = append(std[:index], std[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(std)
}

// initialized the slice
var std []student

// add the students to the student slice...

func addStudent() {
	std = append(std, student{stdId: "1", stdName: "Gayan", stdAge: 25, stdAddress: "Kandy", stdDpt: &department{dptId: "DPT001", dptName: "Mathematics"}})
	std = append(std, student{stdId: "2", stdName: "Sampath", stdAge: 23, stdAddress: "Colombo", stdDpt: &department{dptId: "DPT002", dptName: "Science"}})

}
