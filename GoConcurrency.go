package main

import (
	"fmt"
	"sync"
)

type Course struct {
	Name string
}
type Student struct {
	Name string
}
type RegisterStudentsResults struct {
	Results []StudentRegistrationResult
}
type StudentRegistration struct {
	Student Student
	Course  Course
}
type StudentRegistrationResult struct {
	Registration StudentRegistration
	Error        error
}

func RegisterStudents(students []Student, course Course) RegisterStudentsResults {
	output := make(chan RegisterStudentsResults)
	input := make(chan StudentRegistrationResult)
	var wg sync.WaitGroup
	go handleResults(input, output, &wg)
	defer close(output)
	for _, student := range students {
		wg.Add(1)
		go ConcurrentRegisterStudent(student, course, input)
	}

	wg.Wait()
	close(input)
	return <-output
}

func handleResults(input chan StudentRegistrationResult, output chan RegisterStudentsResults, wg *sync.WaitGroup) {
	var results RegisterStudentsResults
	for result := range input {
		results.Results = append(results.Results, result)
		wg.Done()
	}
	output <- results
}

func ConcurrentRegisterStudent(student Student, course Course, output chan StudentRegistrationResult) {
	result := RegisterStudent(student, course)
	output <- result
}

func RegisterStudent(student Student, course Course) StudentRegistrationResult {
	return StudentRegistrationResult{
		Registration: StudentRegistration{
			Student: student,
			Course:  course,
		},
	}
}
func main() {
	students := []Student{
		Student{Name: "Avinash Maurya"},
		Student{Name: "Hari Joe"},
	}
	result := RegisterStudents(students, Course{Name: "Intro to Golang"})
	fmt.Print(result)
}
