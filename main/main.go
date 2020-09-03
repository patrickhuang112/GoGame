package main

import(
	"fmt"
	"strconv"
	model "GoWithBros/model"
)



func main() {
	model.InitGame(19)
	// fmt.Printf("%v", model.BoardSize)
	model.PrintBoard()
	for model.Passes == 2 {
		var rowString string
		var colString string
		var passOrNo string
		fmt.Println("Pass or not: type y or n")
		fmt.Scanln(&passOrNo)
		if passOrNo == "y" {
			model.TakePass()
		} else {
			for true {
				fmt.Println("Type in a row")
				fmt.Scanln(&rowString)
				fmt.Println("Type in a col")
				fmt.Scanln(&colString)
			
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
}