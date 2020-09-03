package model
/*
import (
	"math"
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

type edge struct {
	a rc
	b rc
}

func createEmptyGameString() GameString {
	return GameString{ make([]edge, 0)}
}

func addToGameString (gs GameString, newRc rc) GameString {
	pieces := gs.pieces	
	pieces = append(pieces, newRc)	
	newGs := createEmptyGameString()
	newGs.pieces = pieces
	return newGs
}

type Player struct {
	color Color
	pieces []rc
	captured int

}

func createNewPlayer(color Color) Player {
	return Player{color, make([]rc, 0), make([]GameString,0), 0 } 
} 

type rc struct {
	row int
	col int
}

type Piece struct {
	color Color
	contains []rc
	rep rc
}

var BoardSize int 
var GameBoard [][]Piece
var CurrentPlayer *Player
var white Player
var black Player

//Initiliaze
func InitGame(boardSize int) {
	BoardSize = boardSize
	initBoard()
	white = createNewPlayer(White)
	black = createNewPlayer(Black)
	CurrentPlayer = &black	
}

func initBoard() {
	board := make([][]Piece, BoardSize)
	for i := 0; i < BoardSize; i++ {
		board[i] = make([]Piece, BoardSize)
		for j := 0; j < BoardSize; j++ {
			board[i][j] = Piece{Empty, make([]rc, 0), rc{-i, -j}}
		}
	} 
	GameBoard = board
}

//Utilities

func hasLibertiesGs(gs GameString) bool {
	for key, _ := range gs.pieces {
		frees := getAdjacents(key)
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
		if(isInBounds(newRc) && GameBoard[newRow][newCol] == Empty) {
			res = append(res, newRc)
		}
	}
	return res
}

func isAdjacent(src rc, vis rc) bool {
	dx := math.Abs(float64(src.row - vis.row))
	dy := math.Abs(float64(src.col - vis.col))
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}














func addPieceToGameString(playerPtr *Player, newRc rc) {
	player := *playerPtr
	for _, orig := range player.pieces {
		if(isAdjacent(orig, newRc)) {
			newGs := addToGameString(*player.strings[orig], newRc)	
			// Get the address of the newly created gamestring
			player.strings[orig] = &newGs
			*playerPtr = player
			return 
		} 
	}
	newGs := createEmptyGameString()
	player.strings[newRc] = &newGs
	*playerPtr = player
}

func placePiece(playerPtr *Player, row int, col int) Result {
	
	if (GameBoard[row][col]).color != Empty {
		return Failure
	} else {
		player := *playerPtr
		piece := rc{row, col}

		// Add to board
		GameBoard[row][col] = (*playerPtr).color	

		// Add to player struct
		pieces := player.pieces
		pieces = append(pieces, piece)
		player.pieces = pieces

		// Add to player gameString
		addPieceToGameString(playerPtr, piece)
		*playerPtr = player
		return Success
	}
}

func removeCaptures(playerPtr *Player) {
	for src,gs := range (*playerPtr).strings {
		if !hasLibertiesGs(*gs) {
			
		}
	}
}










func changePlayer() {
	if(*CurrentPlayer).color == White {
		CurrentPlayer = &black
	} else {
		CurrentPlayer = &white
	}
}

func TakeTurn(row int, col int) Result {
	// Place a piece and update gameboard
	if placePiece(CurrentPlayer, row, col) == Failure {
		return Failure
	}
	// Change the current player
	changePlayer()
	// Removed captured strings and pieces
	return Success	
}
*/