package genitemimage

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

type ItemBundle struct {
	Name          string
	Spritesheets  []*Component
	ReplaceColors []*ReplaceColor
}

type ReplaceColor struct {
	From color.RGBA
	To   []color.RGBA
}

// New returns a new ItemBundle with the given name.
func New(name string) *ItemBundle {
	return &ItemBundle{Name: name}
}

func (i *ItemBundle) AddSpritesheet(s *spritesheet.Spritesheet, name string, optional bool) {
	i.Spritesheets = append(i.Spritesheets, &Component{s: s, Name: name, Optional: optional})
}

func (i *ItemBundle) AddReplaceColor(from color.RGBA, to []color.RGBA) {
	i.ReplaceColors = append(i.ReplaceColors, &ReplaceColor{From: from, To: to})
}

func (i *ItemBundle) Generate() image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 32, 32))
	for _, s := range i.Spritesheets {
		if s.Optional && rand.Intn(2) == 0 {
			continue
		}
		sImg := s.s.TileImage(rand.Intn(s.s.NumTiles()))
		draw.Draw(img, img.Bounds(), sImg, image.Point{0, 0}, draw.Over)
	}

	for _, r := range i.ReplaceColors {
		img = spritesheet.ReplaceColor(img, r.From, r.To[rand.Intn(len(r.To))])
	}

	return img
}

type Component struct {
	s        *spritesheet.Spritesheet
	Name     string
	Optional bool
}
