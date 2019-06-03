package testmod

import "fmt"

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s , this is v1.0.1,done by kevin", name)
}
