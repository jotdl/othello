package othello

// Color represents the current value of a othello stone/field
type Color int

const (
	None Color = iota
	Black
	White
)

// Opposite of the current value. Stays None if current color is None
func (p Color) Opposite() Color {
	switch p {
	case Black:
		return White
	case White:
		return Black
	}

	return None
}
