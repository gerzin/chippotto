package renderer

type Drawable interface {
	Draw(func([]uint8))
}
