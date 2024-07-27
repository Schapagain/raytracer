package canvas

type Color struct {
	R, G, B, A float32
}

type Canvas interface {
	Width() int
	Height() int
	PixelAt(int, int) (Color, error)
	SetPixelAt(int, int, Color) (bool, error)
}

type canvas struct {
	width, height int
	buffer        []Color
}