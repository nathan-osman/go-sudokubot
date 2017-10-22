package solver

// Square represents one individual square and the information gleaned from it.
// Internally, an array of booleans is used to track impossible values for the
// square. As possibilities are eliminated, the array values are set to true
// until one remains, which must be the correct answer.
type Square struct {
	value  int
	values [9]bool
}

// GetValue returns 0 if the square is unsolved or its value if solved.
func (s *Square) GetValue() int {
	return s.value
}

// SetValue indicates that the value for the square is known.
func (s *Square) SetValue(value int) {
	s.value = value
}

// Eliminate indicates that a square cannot be a specific value.
func (s *Square) Eliminate(value int) {
	s.values[value] = true
	falseIdx := 0
	for i, v := range s.values {
		if !v {
			if falseIdx != 0 {
				return
			}
			falseIdx = i + 1
		}
	}
	s.value = falseIdx
}
