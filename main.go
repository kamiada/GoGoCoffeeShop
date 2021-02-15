package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func barista(coffee string) string {
	switch coffee {
	case "Latte":

	case "Americano":

	case "Macchiato":

	}
}

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
		switch userInput {
		case "Latte":
			fmt.Println("Here you go! That's going to be 1.50!")
			os.Exit(0)
		case "Americano":
			fmt.Println("Here you go! That's going to be €0.80!")
			os.Exit(0)
		case "Macchiato":
			fmt.Println("Here you go! That's going to be €1.30")
			os.Exit(0)
		default:
			fmt.Println("I don't understand your order")
		}
	}
}
