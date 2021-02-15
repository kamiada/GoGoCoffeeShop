package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	menu, err := ioutil.ReadFile("./menu.txt")
	if err != nil {
		log.Fatal(err)
	}
	bashMenu := string(menu)
	fmt.Println("Hello customer! Here is our menu: " + bashMenu + "...What would you like to order?")
	//for loops in Go are a bit strange
	for {
		userInput, _ := reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
		if strings.Compare("Latte", userInput) == 0 {
			fmt.Println("Here you go! That's going to be 1.50!")
		}
	}
}
