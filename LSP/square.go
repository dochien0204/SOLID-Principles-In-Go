package lsp

type Square struct {
	Rectangle
}

func (s *Square) SetHeight(height int) {
	s.height = height
	s.width = height
}
