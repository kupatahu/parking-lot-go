package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/kupatahu/parking-lot-go/lot"
	"github.com/kupatahu/parking-lot-go/prompt"
)

func main() {
	p := prompt.New(bufio.NewReader(os.Stdin))

	lot := createLot(p)

	promptOptions(p, lot)
}

func createLot(p *prompt.Prompt) *lot.Lot {
	input, err := p.GetInput("Input parking lot size: ")

	if err != nil {
		fmt.Println("Let's try again")
		return createLot(p)
	}

	size, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Size must be a number")
		return createLot(p)
	}

	return lot.New(size)
}

func promptOptions(p *prompt.Prompt, lot *lot.Lot) {
	input, err := p.GetInput("Input command: ")

	if err != nil {
		fmt.Println("Let's try again")
		promptOptions(p, lot)
		return
	}

	switch input {
	case "p":
		park(p, lot)
		promptOptions(p, lot)
		return
	case "u":
		unpark(p, lot)
		promptOptions(p, lot)
		return
	case "s":
		status(lot)
		promptOptions(p, lot)
		return
	case "e":
		return
	default:
		fmt.Println("Unknown commmand")
		promptOptions(p, lot)
		return
	}
}

func park(p *prompt.Prompt, lot *lot.Lot) {
	plate, err := p.GetInput("Enter plate number: ")

	if err != nil {
		fmt.Println("Let's try again")
		promptOptions(p, lot)
		return
	}

	err = lot.Park(plate)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Car %v parked", plate)
}

func unpark(p *prompt.Prompt, lot *lot.Lot) {
	input, err := p.GetInput("Enter plate number: ")

	if err != nil {
		fmt.Println("Let's try again")
		promptOptions(p, lot)
		return
	}
	
	plate, err := lot.Unpark(input)
	
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Car %v unparked", plate)
}

func status(lot *lot.Lot) {
	fmt.Print(lot.Status())
}
