package sections

type Renderer interface {
	Render(w, h int) string
}

type RenderFunc func(w, h int) string

func (r RenderFunc) Render(w, h int) string {
	return r(w, h)
}
