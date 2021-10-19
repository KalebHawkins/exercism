package twofer

import "fmt"

// ShareWith returns the string of the person you are sharing with.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %v, one for me.", name)
}
