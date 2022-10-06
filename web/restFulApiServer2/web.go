package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	Id    int
	Name  string
	Age   int
	Score int
}

var students map[int]Student
var lastId int

func MakeWebHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/students", GetStudentListHandler).Methods("GET")
	mux.HandleFunc("/students/{id:[0-9]+}", GetStudentHandler).Methods("GET")

	students = make(map[int]Student)
	students[1] = Student{1, "uji", 18, 90}
	students[2] = Student{2, "bjd", 18, 99}
	lastId = 2

	return mux
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}
func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Students) Less(i, j int) bool {
	return s[i].Id < s[j].Id
}

func GetStudentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	student, ok := students[id]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func GetStudentListHandler(w http.ResponseWriter, r *http.Request) {
	list := make(Students, 0)
	for _, v := range students {
		list = append(list, v)
	}
	sort.Sort(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)

}

func main() {
	http.ListenAndServe(":3000", MakeWebHandler())
}
