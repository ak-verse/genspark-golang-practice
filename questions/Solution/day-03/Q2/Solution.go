package main

import "fmt"

func main() {
	users := []string{"a", "b"}

	emp := []string{"p", "q"}

	users = append(users, "d")
	inspect("users", users)

	emp = append(emp, "s")
	inspect("emp", emp)

}

func inspect(name string, slice []string) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()

}
