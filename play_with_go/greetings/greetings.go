package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// RandomGreet pics a random greet from the list
func RandomGreet() string{
	rand.Seed(time.Now().UnixNano())
	greets := []string{
		"Welcome",
		"Hello",
		"Howdy",
		"Sup",
	}
	randInt := rand.Intn(len(greets))
	return fmt.Sprintf("%v", greets[randInt])
}

// Hellos greets a list of people
func Hellos(names []string) ([]string, []error) {
	greeted := make([]string, 0)
	errs := make([]error, 0)
	for _, name := range names {
		message, err := Hello(name)
		greeted = append(greeted, message)
		errs = append(errs, err)
	}
	return greeted, errs
}

// Hello takes a name as a parameter and returns a greeting
// for that name
func Hello(name string) (string, error) {
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("%v, %v",RandomGreet(), name)
	return message, nil
}