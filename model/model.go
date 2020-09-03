package model

import (
	"fmt"
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

func createEmptyPiece() Piece {
	return Piece{color : Empty}
}

func createNewPiece(color Color, src rc) Piece {
	return Piece{color, src, rc{-1,-1}, []rc{src}}
}

var BoardSize int 
var GameBoard [][]Piece
var CurrentPlayer *Player
var OpposingPlayer *Player
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
	CurrentPlayer = &black	
	OpposingPlayer = &white
}

func initBoard() {
	board := make([][]Piece, BoardSize)
	for i := 0; i < BoardSize; i++ {
		board[i] = make([]Piece, BoardSize)
		for j := 0; j < BoardSize; j++ {
			board[i][j] = createEmptyPiece()
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

func (color Color) toString() string {
	if color == White {
		return "White"
	} else {
		return "Black"
	}
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
}

func PrintCurrentPlayer() {
	if (*CurrentPlayer).color == White {
		fmt.Println("White")
	} else {
		fmt.Println("Black")
	}
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

func hasLiberties(piece Piece) bool {
	for _, src := range piece.contains {
		frees := getAdjacents(src)
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
	board := GameBoard
	possible := []rc{rc{0,1}, rc{1,0}, rc{-1,0}, rc{0,-1}}
	for _,d := range possible {
		newCol := src.col + d.col 
		newRow := src.row + d.row
		newRc := rc{newRow, newCol}	
		if (isInBounds(newRc) && (board[newRow][newCol]).color == Empty) {
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









func combinePieces(playerPtr *Player, a Piece, b Piece) {
	board := GameBoard
	for a.rep.row >= 0 || b.rep.row >= 0 {
		if (a.rep.row < 0 && b.rep.row < 0) {
			aHeight := abs(a.rep.row)
			bHeight := abs(b.rep.row)
			newContains := append(a.contains, b.contains...)
			var repToRemove rc
			if (aHeight >= bHeight) {
				a.contains = newContains
				a.rep = rc{a.rep.row-1, a.rep.col-1}
				b.contains = make([]rc, 0)
				b.rep = rc{a.src.row, a.src.col}
				repToRemove = b.rep	
			} else {
				b.contains = newContains
				b.rep = rc{b.rep.row-1, b.rep.col-1}
				a.contains = make([]rc, 0)
				a.rep = rc{b.src.row, b.src.col}
				repToRemove = a.rep
			}
			board[a.src.row][a.src.col] = a	
			board[b.src.row][b.src.col] = b		
			GameBoard = board
			PrintBoardAtPos(a.src.row, a.src.col)
			PrintBoardAtPos(b.src.row, b.src.col)
			// Removing the reps from the player
			player := *playerPtr
			
			if  _,found := player.reps[repToRemove]; found {
				// Remove the smaller merged rep from the player
				delete(player.reps, repToRemove)	
			}
			
			*playerPtr = player
			return	
		} 
		if(a.rep.row >= 0) {
			a = board[a.rep.row][a.rep.col]
		} 
		if(b.rep.row >= 0) {
			b = board[b.rep.row][b.rep.col]
		}
	}
}


func placePiece(playerPtr *Player, row int, col int) Result {
	board := GameBoard	
	if (board[row][col]).color != Empty {
		return Failure
	} else {
		player := *playerPtr
		newRc := rc{row, col}
		// Add to board
		newColor := player.color
		newPiece := createNewPiece(newColor, newRc)	
		board[row][col] = newPiece
		adjacents := getAdjacents(newRc)
		curRep := newPiece.rep
		for _, adj := range adjacents {
			if board[adj.row][adj.col].color == newColor {
				adjPiece := board[adj.row][adj.col]	
				if(curRep != adjPiece.rep) {
					combinePieces(playerPtr, adjPiece, newPiece)	
					newPiece = board[row][col]
				}	
			}	
		}
		// Add to player gameString
		*playerPtr = player
		return Success
	}
}

func removePieceListFromBoard(pieces []rc) {
	board := GameBoard
	for _,src := range pieces {
		board[src.row][src.col] = createEmptyPiece()
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
	// Place a piece and update gameboard
	if placePiece(CurrentPlayer, row, col) == Failure {
		return Failure
	}
	// Removed captured strings and pieces
	removeCaptures(OpposingPlayer, CurrentPlayer)
	changePlayer()
	Passes = 0
	return Success	
}

func TakePass() {
	Passes++
	changePlayer()
} 