package othello

import "testing"

func TestBoard_findValidDirectionsOfMove(t *testing.T) {
	board := newBoard(8)

	// this should be a valid move
	dir, err := board.findValidDirectionsOfMove(2, 4, Black)

	if err != nil {
		t.Error(err)
	}

	if len(dir) != 1 {
		t.Errorf("There should be 1 valid direction, but were %v", len(dir))
	}

	if len(dir) < 1 {
		return
	}

	if dir[0].X != 0 && dir[0].Y != 1 {
		t.Errorf("There should be the valid direction (X=0,Y=1) downwards, but found (X=%v,Y=%v)", dir[0].X, dir[0].Y)
	}
}
