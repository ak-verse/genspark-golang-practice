package stringsops

import "strings"

// a non main package, like calc in our case is known as library package
// lib package can be used by other packages of the same project
// or if someone has imported your project then others can too use the functionality
// from the library package

// package name should be same as the folder name

// if first letter of the function in uppercase, then func would be exported
// it means other packages can call this function and use this
// Add function is exported for others to use


func reverseString(s string) string {
	return strings.ToLower(s)
}
