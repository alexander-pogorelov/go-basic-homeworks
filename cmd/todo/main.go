// main package
package main

import "fmt"

func main() {
	var name string
	fmt.Print("Как вас зовут? ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Printf("Привет, %s!\n", name)
}
