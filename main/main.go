package main

import (
	"fmt"
	"strconv"
	model "GoWithBros/model"
)



func main() {
	
	model.InitGame(19)
	model.PrintBoard()

	for model.Passes != 2 {
		var rowString string
		var colString string
		fmt.Print("Current player is:")
		model.PrintCurrentPlayer()
		fmt.Println("White info: ")
		model.PrintWhite()
		fmt.Println("Black info: ")
		model.PrintBlack()

		for true {
			fmt.Println("Type in a row (type p if pass turn)")
			fmt.Scanln(&rowString)
			if (rowString == "p") {
				model.TakePass()
				fmt.Println("Player changed")
				break
			}
			fmt.Println("Type in a col (type p if pass turn)")
			fmt.Scanln(&colString)
			if (colString == "p") {
				model.TakePass()
				fmt.Println("Player changed")
			} else {
				row, _ := strconv.Atoi(rowString) 
				col, _ := strconv.Atoi(colString)

				res := model.TakeTurn(row, col)
				if res == model.Failure {
					fmt.Println("Invalid row and col!")
				} else {
					model.PrintBoard()
					break
				}
			}
		}
	}
	
	winner := model.EndGame()

	fmt.Println("")
	fmt.Println("Final White info: ")
	model.PrintWhite()
	fmt.Println("Final Black info: ")
	model.PrintBlack()

	fmt.Print(winner.ToString())
	fmt.Println(" is the winner!")
}