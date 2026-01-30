package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name given")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Elo, %v",
		"Bic boi %v",
		"Greeting 3 %v",
	}

	return formats[rand.Intn(len(formats))]

}
