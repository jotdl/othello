package othello

import "errors"

import "fmt"

type Player int

const (
	None Player = iota
	Player1
	Player2
)

type Field struct {
	Column int
	Row    int

	Value Player
}

type Board struct {
	Dimension int
	Fields [][]Player
}

type Score [2]uint8

func (b *Board) Score() Score {
	score := Score{}

	for i := 0; i < len(b.Fields); i++ {
		for j := 0; j < len(b.Fields[i]); j++ {
			field := b.Fields[i][j]

			switch field {
			case Player1:
				score[0]++
			case Player2:
				score[1]++
			}
		}
	}

	return score
}

func (b *Board) Field(row, column int) Field {
	if row > len(b.Fields) || column > len(b.Fields[row]) {
		return Field{}
	}
	return Field{Column: column, Row: row, Value: b.Fields[row][column]}
}

func (b *Board) Clone() Board {
	clone := Board{
		Fields: make([][]Player, len(b.Fields)),
	}

	for i, column := range b.Fields {
		clonedColumn := make([]Player, len(column))
		for i, val := range column {
			clonedColumn[i] = val
		}
		clone.Fields[i] = clonedColumn
	}

	return clone
}

func (b *Board) isMoveOnBoard(row, column int) error {
	if(row >= len(b.Fields) || column >= len(b.Fields[row])) {
		return fmt.Errorf("move needs to have valid row (0 < row < %v) and valid colum (0 < column < %v), but had row = %v, column = %v", len(b.Fields), len(b.Fields[row]))
	}

	return nil
}

func (b *Board) IsValidMove(row, column int, player Player) error {
	// check horizontally within the same row
	for i := 0; i < 
}

func (b *Board) MakeTurn(row, column int, player Player) {

}
