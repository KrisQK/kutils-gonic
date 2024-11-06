package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, errors.New("empty names")
		}
		messages[name] = message
	}
	return messages, nil
}

// func Hellos(names []string) ([]string, error) {
// 	if len(names) == 0 { // 检查 names 是否为空
// 		return nil, errors.New("empty names")
// 	}

// 	messages := []string{}
// 	for i := 0; i < len(names); i++ {
// 		// 正确使用 fmt.Sprintf 进行格式化
// 		message := fmt.Sprintf(randomFormat(), names[i])
// 		messages = append(messages, message)
// 	}
// 	return messages, nil
// }

// func Hellos(names []string) (string, error) {
// 	if names == nil {
// 		return names, errors.New("empty names")
// 	}

// 	messages := []string{}
// 	for i := 0; i < len(names); i++ {
// 		message := fmt.Sprint(randomFormat(), names[i])
// 		messages = append(messages, message)
// 	}
// 	return messages, nil
// }

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}
