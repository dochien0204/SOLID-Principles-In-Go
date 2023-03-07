package lsp

type Rectangle struct {
	height int
	width  int
}

func (r *Rectangle) GetArea() int {
	return r.height * r.width
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetWidth() int {
	return r.width
}
