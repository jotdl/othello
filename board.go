package othello

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

// direction reflects a simple direction on the othello boards
// used for easier navigation
type direction struct {
	X int
	Y int
}

// all possible directions
var directions = [8]direction{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, -1},
	{-1, 0},
	{-1, 1},
}

// Board represents the current game state
type Board struct {
	Dimension int       // Dimension of this board, defaults to 8
	Fields    [][]Color // Slice of Slice of fields. Outer slice represents rows while inner slice represents column
}

// newBoard creates a new board with the given dimension
// E.g. for a board with the dimension of 8 initial state would look as below:
//       0   1   2   3   4   5   6   7
//     ---------------------------------
//   0 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   1 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   2 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   3 |   |   |   | ● | ○ |   |   |   |
//     ---------------------------------
//   4 |   |   |   | ○ | ● |   |   |   |
//     ---------------------------------
//   5 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   6 |   |   |   |   |   |   |   |   |
//     ---------------------------------
//   7 |   |   |   |   |   |   |   |   |
//     ---------------------------------
func newBoard(dim int) *Board {
	b := Board{
		Dimension: dim,
		Fields:    make([][]Color, dim),
	}

	for i := 0; i < dim; i++ {
		b.Fields[i] = make([]Color, dim)
	}

	lower := int(math.Floor(float64(dim-1) / 2.0))
	upper := int(math.Ceil(float64(dim-1) / 2.0))

	b.Fields[lower][lower] = Black
	b.Fields[upper][upper] = Black
	b.Fields[upper][lower] = White
	b.Fields[lower][upper] = White

	return &b
}

// Score returns the current score for the given board
func (b Board) Score() Score {
	score := Score{}

	for i := 0; i < len(b.Fields); i++ {
		for j := 0; j < len(b.Fields[i]); j++ {
			field := b.Fields[i][j]

			switch field {
			case Black:
				score[0]++
			case White:
				score[1]++
			}
		}
	}

	return score
}

// ScoreOf returns only the score of the given player
func (b Board) ScoreOf(player Color) uint8 {
	scores := b.Score()

	switch player {
	case Black:
		return scores[0]
	case White:
		return scores[1]
	}
	return 0
}

// Field returns the Field at the given location. Returns an empty Field if requested location is invalid
func (b Board) Field(row, column int) Field {
	if !b.IsValidPos(row, column) {
		return Field{}
	}
	return Field{Column: column, Row: row, Value: b.Fields[row][column]}
}

// Clone the current board so it's save to modify
func (b Board) Clone() *Board {
	clone := Board{
		Dimension: b.Dimension,
		Fields:    make([][]Color, b.Dimension),
	}

	for i, column := range b.Fields {
		clonedColumn := make([]Color, b.Dimension)
		copy(clonedColumn, column)
		clone.Fields[i] = clonedColumn
	}

	return &clone
}

func (b Board) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(strings.Repeat("-", b.Dimension*4+1))
	buf.WriteRune('\n')
	for row := 0; row < b.Dimension; row++ {
		for column := 0; column < b.Dimension; column++ {
			value := b.Fields[row][column]

			buf.WriteRune('|')
			buf.WriteRune(' ')
			var valRune rune
			switch value {
			case None:
				valRune = ' '
			case Black:
				valRune = '●'
			case White:
				valRune = '○'
			}
			buf.WriteRune(valRune)
			buf.WriteRune(' ')

		}

		buf.WriteRune('|')
		buf.WriteRune('\n')
		buf.WriteString(strings.Repeat("-", b.Dimension*4+1))
		buf.WriteRune('\n')
	}

	return buf.String()
}

// IsValidPos checks if the specified location is on this board or outside. Returns false if outside
func (b Board) IsValidPos(row, column int) bool {
	return row >= 0 && row < b.Dimension && column >= 0 && column < b.Dimension
}

func (b Board) findValidDirectionsOfMove(row, column int, player Color) ([]direction, error) {
	if !b.IsValidPos(row, column) {
		return nil, fmt.Errorf("move needs to have valid row (0 < row < %v) and valid colum (0 < column < %v), but had row = %v, column = %v", b.Dimension, b.Dimension, row, column)
	}

	// 1. First condition is that current move is still empty
	if b.Fields[row][column] != None {
		return nil, fmt.Errorf("field in row %v, column %v already taken by player", row, column)
	}

	enemy := player.Opposite()

	validDirections := make([]direction, 0)
	// 2. First field next to current move needs to be owned by enemy
	for _, dir := range directions {
		nextRow := row + dir.Y
		nextColumn := column + dir.X

		if !b.IsValidPos(nextRow, nextColumn) || // check that next field is still valid
			b.Fields[nextRow][nextColumn] != enemy { // if adjacent field doesn't belong to other player this is no valid move (at least in this direction)
			continue
		}

		// check that at some point there is field by the own player
		for b.IsValidPos(nextRow, nextColumn) {
			value := b.Fields[nextRow][nextColumn]
			if value == None {
				break
			}

			if value == player {
				validDirections = append(validDirections, dir)
				break
			}

			// go to next field
			nextRow += dir.Y
			nextColumn += dir.X
		}
	}

	return validDirections, nil
}

// IsValidTurn checks if a move for a given player/color is valid. Returns an error if not
func (b Board) IsValidTurn(row, column int, player Color) error {
	directions, err := b.findValidDirectionsOfMove(row, column, player)
	if err != nil {
		return err
	}

	if len(directions) == 0 {
		return fmt.Errorf("Move in row %v, column %v is no valid move in any direction", row, column)
	}

	return nil
}

// MakeTurn makes the given turn with the specified location and color. Will return an error if move is not valid
func (b *Board) MakeTurn(row, column int, player Color) error {
	directions, err := b.findValidDirectionsOfMove(row, column, player)
	if err != nil {
		return err
	}

	if len(directions) == 0 {
		return fmt.Errorf("Move in row %v, column %v is no valid move in any direction", row, column)
	}

	b.Fields[row][column] = player

	enemy := player.Opposite()

	for _, dir := range directions {
		nextRow := row + dir.Y
		nextCol := column + dir.X

		for b.IsValidPos(nextRow, nextCol) {
			value := b.Fields[nextRow][nextCol]
			if value != enemy {
				break
			}

			b.Fields[nextRow][nextCol] = player

			// go to next field
			nextRow += dir.Y
			nextCol += dir.X
		}
	}

	return nil
}
