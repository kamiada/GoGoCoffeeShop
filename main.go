package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	menu, err := ioutil.ReadFile("./menu.txt")
	if err != nil {
		log.Fatal(err)
	}
	bashMenu := string(menu)
	fmt.Println("Hello customer! Here is our menu: " + bashMenu + "...What would you like to order?")
}
