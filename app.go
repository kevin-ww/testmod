package testmod

// import "fmt"

// // Hi returns a friendly greeting
// func Hi(name string) string {
// 	return fmt.Sprintf("Hi, %s , this is v1.0.1,done by kevin", name)
// }

import (
	"errors"
	"fmt"
)

// Hi returns a friendly greeting in language lang
func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("unknown language")
	}
}
