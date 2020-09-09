package model

import (
	"fmt"
	"strconv"
)

type Result int
const (
	Success Result = iota
	Failure
)

type Color int
const (
	White Color = iota
	Black
	Empty
)

type Player struct {
	color Color
	reps map[rc]bool
	captured int

}

func createNewPlayer(color Color) Player {
	return Player{color, make(map[rc]bool), 0 } 
} 

type rc struct {
	row int
	col int
}

type Piece struct {
	color Color
	src rc
	rep rc
	contains []rc
}

func createNewPiece(color Color, src rc) Piece {
	return Piece{color, src, rc{-1,-1}, []rc{src}}
}

var BoardSize int 
var GameBoard [][]Piece
var CurrentPlayer *Player
var OpposingPlayer *Player
var EmptyPlayer *Player
var white Player
var black Player
var Passes int

//Initiliaze
func InitGame(boardSize int) {
	Passes = 0
	BoardSize = boardSize
	initBoard()
	white = createNewPlayer(White)
	black = createNewPlayer(Black)
	emptyPlr := createNewPlayer(Empty)
	EmptyPlayer = &emptyPlr
	CurrentPlayer = &black	
	OpposingPlayer = &white
}

func initBoard() {
	board := make([][]Piece, BoardSize)
	for i := 0; i < BoardSize; i++ {
		board[i] = make([]Piece, BoardSize)
		for j := 0; j < BoardSize; j++ {
			board[i][j] = createNewPiece(Empty, rc{i,j})
		}
	} 
	GameBoard = board
}

func PrintBoard() {
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			piece := GameBoard[i][j]
			if piece.color == Empty {
				fmt.Print("-")
			} else if piece.color == White {
				fmt.Print("W")
			} else {
				fmt.Print("B")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}
//Utilities

func copyRepMap(reps map[rc]bool) map[rc]bool{
	newMap := make(map[rc]bool)
	for src,_ := range reps {
		newMap[src] = true
	}
	return newMap
}

func printRowCol(row int, col int) {
	fmt.Print("Row: ")
	fmt.Println(strconv.Itoa(row))
	fmt.Print("Col: ")
	fmt.Println(strconv.Itoa(col))
	fmt.Println("")
}

func (color Color) ToString() string {
	if color == White {
		return "White"
	} else if color == Black{
		return "Black"
	}
	return "No one"
}

func PrintBoardAtPos(row int, col int) {
	board := GameBoard
	piece := board[row][col]
	if piece.color == Empty {
		fmt.Print("Empty")
	} else if piece.color == White {
		fmt.Print("W")
	} else {
		fmt.Print("B")
	}	
	fmt.Println("")
	fmt.Println("Rep :")
	printRowCol(piece.rep.row, piece.rep.col)
}

func PrintCurrentPlayer() {
	if (*CurrentPlayer).color == White {
		fmt.Println("White")
		
	} else {
		fmt.Println("Black")
	}
	fmt.Println("")
}

func PrintWhite() {
	printPlayer(&white)
}

func PrintBlack() {
	printPlayer(&black)
}

func printPlayer(playerPtr *Player) {
	reps := (*playerPtr).reps
	fmt.Println("Reps: ")
	for rp,_ := range reps {
		fmt.Print("Row: ")
		fmt.Print(strconv.Itoa(rp.row))
		fmt.Print("    Col: ")
		fmt.Println(strconv.Itoa(rp.col))
	}
	fmt.Print("Captured: ")
	fmt.Print(strconv.Itoa((*playerPtr).captured))
	fmt.Println("")
	fmt.Println("")
}


func abs(num int) int {
	if (num < 0) {
		return -num
	} else {
		return num
	}
}

func remove(s []rc, i int) []rc {
	return append(s[:i], s[i+1:]...)
}

func territoryForWhichColor(piece Piece) Color {
	color := Empty
	for _, src := range piece.contains {
		adjs := getAdjacents(src)
		for _, adj := range adjs {
			adjPiece := GameBoard[adj.row][adj.col]
			if (color == White && adjPiece.color == Black) || 
			   (color == Black && adjPiece.color == White) {
				return Empty
			}
			color = adjPiece.color	
		}
	}
	return color
}

func hasLiberties(piece Piece) bool {
	for _, src := range piece.contains {
		frees := getEmptyAdjacents(src)
		if len(frees) != 0 {
			return true
		}
	}
	return false
}

func isInBounds (src rc) bool {
	return src.row >= 0 && src.row < BoardSize && src.col >= 0 && src.col < BoardSize
}

func getAdjacents(src rc) []rc {
	res := make([]rc, 0)
	possible := []rc{rc{0,1}, rc{1,0}, rc{-1,0}, rc{0,-1}}
	for _,d := range possible {
		newCol := src.col + d.col 
		newRow := src.row + d.row
		newRc := rc{newRow, newCol}	
		if (isInBounds(newRc)) {
			res = append(res, newRc)
		}
	}
	return res
}

func getEmptyAdjacents(src rc) []rc {
	res := make([]rc, 0)
	possible := []rc{rc{0,1}, rc{1,0}, rc{-1,0}, rc{0,-1}}
	for _,d := range possible {
		newCol := src.col + d.col 
		newRow := src.row + d.row
		newRc := rc{newRow, newCol}	
		if (isInBounds(newRc) && GameBoard[newRow][newCol].color == Empty) {
			res = append(res, newRc)
		}
	}
	return res
}

func isAdjacent(src rc, vis rc) bool {
	dx := abs(src.row - vis.row)
	dy := abs(src.col - vis.col)
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}





func EndGame() Color {
	board := make([][]Piece, BoardSize)	
	for i := 0; i < BoardSize; i++ {
		board[i] = make([]Piece, BoardSize)
		for j := 0; j < BoardSize; j++ {
			board[i][j] = GameBoard[i][j]
		}
	}

	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			piece := GameBoard[i][j]
			// fmt.Println("We are here: ")
			// printRowCol(i,j)
			if piece.color == Empty {
				src := rc{i,j}
				adjacents := getEmptyAdjacents(src)
				combined := false
				
				for _,adj := range adjacents {
					
					adjPiece := GameBoard[adj.row][adj.col]	
					// printRowCol(adjPiece.rep.row, adjPiece.rep.col)
					// printRowCol(piece.rep.row, piece.rep.col)
					combinePieces(EmptyPlayer, adjPiece, piece)
					piece = GameBoard[src.row][src.col]
					combined = true	
				}

				if !combined {
					(*EmptyPlayer).reps[src] = true
				}
			}
		}
	}
	countTerritory(CurrentPlayer)
	countTerritory(OpposingPlayer)

	if white.captured > black.captured {
		return White
	}
	if black.captured > white.captured {
		return Black
	}
	return Empty
}

func countTerritory(playerPtr *Player) {
	player := *playerPtr
	empty := *EmptyPlayer
	for src,_ := range empty.reps {
		piece := GameBoard[src.row][src.col]
		color := territoryForWhichColor(piece)
		if color == player.color {
			player.captured += len(piece.contains)
		} 
	}
	*playerPtr = player
}


func combinePieces(playerPtr *Player, a Piece, b Piece) {
	board := GameBoard
	
	/*
	fmt.Println("SRCS")
	printRowCol(a.src.row, a.src.col)
	printRowCol(b.src.row, b.src.col)
	fmt.Println("REPS")
	printRowCol(a.rep.row, a.rep.col)
	printRowCol(b.rep.row, b.rep.col)
	*/

	for a.rep.row >= 0 || b.rep.row >= 0 {
		if (a.rep.row >= 0) {
			a = board[a.rep.row][a.rep.col]
		} 
		if (b.rep.row >= 0) {
			b = board[b.rep.row][b.rep.col]
		}
	}
	
	// They are already combined
	if (a.rep.row != -1 && a.rep.row == b.rep.row) {
		return
	}
	//printRowCol(a.rep.row, a.rep.col)
	//printRowCol(b.rep.row, b.rep.col)
	aHeight := abs(a.rep.row)
	bHeight := abs(b.rep.row)
	newContains := append(a.contains, b.contains...)
	var repToRemove rc
	var repToAdd rc
	if (aHeight >= bHeight) {
		a.contains = newContains
		a.rep = rc{a.rep.row-1, a.rep.col-1}
		b.contains = make([]rc, 0)
		b.rep = rc{a.src.row, a.src.col}
		repToRemove = b.src	
		repToAdd = a.src
	} else {
		b.contains = newContains
		b.rep = rc{b.rep.row-1, b.rep.col-1}
		a.contains = make([]rc, 0)
		a.rep = rc{b.src.row, b.src.col}
		repToRemove = a.src
		repToAdd = b.src
	}

	board[a.src.row][a.src.col] = a	
	board[b.src.row][b.src.col] = b		
	GameBoard = board

	// Removing the reps from the player
	player := *playerPtr
	player.reps[repToAdd] = true
	if  _,found := player.reps[repToRemove]; found {
		// Remove the smaller merged rep from the player
		delete(player.reps, repToRemove)	
	}
	// Update player
	*playerPtr = player
}


func placePiece(playerPtr *Player, row int, col int) Result {
	if (GameBoard[row][col]).color != Empty {
		return Failure
	} else {
		player := *playerPtr
		newRc := rc{row, col}
		// Add to board
		newColor := player.color
		newPiece := createNewPiece(newColor, newRc)	
		GameBoard[row][col] = newPiece
		adjacents := getAdjacents(newRc)
		combined := false
		
		// Combining pieces if needed
		for _, adj := range adjacents {
			if GameBoard[adj.row][adj.col].color == newColor {
				adjPiece := GameBoard[adj.row][adj.col]	
				combinePieces(playerPtr, adjPiece, newPiece)
				newPiece = GameBoard[row][col]
				combined = true
			}	
		}
			
		if !combined {
			player.reps[newRc] = true
		}
		// Add to player gameString
		*playerPtr = player
		
		return Success
	}
}

func removePieceListFromBoard(pieces []rc) {
	board := GameBoard
	for _,src := range pieces {
		board[src.row][src.col] = createNewPiece(Empty, src)
	}
	GameBoard = board
}

func removeCaptures(remPlayerPtr *Player, opPlayerPtr *Player) {
	// Stuff removed from remPlayer, score added to opPlayer
	board := GameBoard
	remPlayer := *remPlayerPtr
	opPlayer := *opPlayerPtr
	newReps := remPlayer.reps
	for src,_ := range newReps {
		piece := board[src.row][src.col]
		if !hasLiberties(piece)	 {
			delete(newReps, src)
			opPlayer.captured += len(piece.contains)	
			removePieceListFromBoard(piece.contains)
		} 
	}
	remPlayer.reps = newReps
	*remPlayerPtr = remPlayer
	*opPlayerPtr = opPlayer
}










func changePlayer() {
	if (*CurrentPlayer).color == White {
		CurrentPlayer = &black
		OpposingPlayer = &white
	} else {
		CurrentPlayer = &white
		OpposingPlayer = &black
	}
}

func TakeTurn(row int, col int) Result {
	// Have to use this to copy a map (they are reference types, not primitive)
	
	oldRepMap := copyRepMap((*CurrentPlayer).reps)

	// Place a piece and update gameboard
	if placePiece(CurrentPlayer, row, col) == Failure {
		return Failure
	}

	// Removed captured strings and pieces
	removeCaptures(OpposingPlayer, CurrentPlayer)

	// If the piece we just put down at row, col has no liberties after 
	// removing opposing player captured strings, that means the move is invalid,
	// so we remove the piece in the gameboard and revert the player to what it
	// was like before the move was made
	if !hasLiberties(GameBoard[row][col]) {
		GameBoard[row][col] = createNewPiece(Empty, rc{row,col})
		(*CurrentPlayer).reps = oldRepMap
		return Failure
	}

	changePlayer()
	Passes = 0
	return Success	
}

func TakePass() {
	Passes++
	changePlayer()
} 