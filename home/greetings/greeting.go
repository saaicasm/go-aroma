package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("invalid String")
	}
	message := fmt.Sprintf(RandomGenerator(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func RandomGenerator() string {
	formats := []string{
		"Even in darkest time there is always %v",
		"You will never walk alone %v",
		"%v, Get on the dance floor",
	}

	return formats[rand.Intn(len(formats))]
}
