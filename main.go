package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createStudent(w http.ResponseWriter, r *http.Request) {
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateStudent(student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if _, exists := students[student.ID]; exists {
		http.Error(w, "Student with this ID already exists", http.StatusBadRequest)
		return
	}
	students[student.ID] = student
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func getStudents(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	list := make([]Student, 0, len(students))
	for _, student := range students {
		list = append(list, student)
	}
	json.NewEncoder(w).Encode(list)
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	mu.RLock()
	student, ok := students[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	var student Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validateStudent(student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := students[id]; !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	student.ID = id // Ensure ID remains unchanged
	students[id] = student
	json.NewEncoder(w).Encode(student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if _, ok := students[id]; !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	delete(students, id)
	w.WriteHeader(http.StatusNoContent)
}

func getStudentSummary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}
	mu.RLock()
	student, ok := students[id]
	mu.RUnlock()
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	summary, err := generateSummary(student)
	if err != nil {
		http.Error(w, "Failed to generate summary: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/students/{id}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{id}", deleteStudent).Methods("DELETE")
	r.HandleFunc("/students/{id}/summary", getStudentSummary).Methods("GET")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
