package main

import (
	"errors"
	"sync"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var (
	students = make(map[int]Student)
	mu       sync.RWMutex
)

func validateStudent(student Student) error {
	if student.ID <= 0 {
		return errors.New("ID must be positive")
	}
	if student.Name == "" {
		return errors.New("Name is required")
	}
	if student.Age <= 0 {
		return errors.New("Age must be positive")
	}
	if student.Email == "" {
		return errors.New("Email is required")
	}
	return nil
}
