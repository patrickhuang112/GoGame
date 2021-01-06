# GoGame
The board game Go, currently only in console form, but with plans to implement a GUI in Go using Fyne.io

## How to Build
First navigate to the main directory
> cd main
>
> go build
  
This creates a clickable executable in the main folder

## How to Run
To run the project directly without building
> cd main
>
> go run main.go

### How to Play:
Black goes first. Type in a row and col to indicate where you want to place the piece on the board. The board is 0 indexed, so 0 <= row <= 18 and 0 <= col <= 18 where row,col are the numbers typed in. Information will printed on the console about the current state of the game

### Current Features:
The basic rules of the board game Go, including liberties, capturing strings, and counting territories at the end of the game.

### Future features: 
UI with Fyne
