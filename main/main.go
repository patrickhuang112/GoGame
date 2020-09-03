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
	for true {
		var rowString string
		var colString string
		fmt.Println("Type in a row")
		fmt.Scanln(&rowString)
		fmt.Println("Type in a col")
		fmt.Scanln(&colString)
		
		row, _ := strconv.Atoi(rowString) 
		col, _ := strconv.Atoi(colString)
		model.TakeTurn(row, col)
		model.PrintBoard()
		model.PrintCurrentPlayer()
	}
}