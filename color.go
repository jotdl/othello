package othello

type Color int

const (
	None Color = iota
	Black
	White
)

func (p Color) Opposite() Color {
	switch p {
	case Black:
		return White
	case White:
		return Black
	}

	return None
}
