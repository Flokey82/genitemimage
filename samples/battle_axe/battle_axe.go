package battleaxe

import (
	"image/color"

	_ "embed"

	"github.com/Flokey82/genitemimage"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

//go:embed sprites/heads_32_1x4.png
var heads_png []byte

//go:embed sprites/handles_32_1x1.png
var handles_png []byte

// #9badb7
var defaultColorHead = color.RGBA{R: 0x9b, G: 0xad, B: 0xb7, A: 0xff}

// #663931
var defaultColorHandle = color.RGBA{R: 0x66, G: 0x39, B: 0x31, A: 0xff}

func New() (*genitemimage.ItemBundle, error) {
	headSheet, err := spritesheet.New(heads_png, 32)
	if err != nil {
		return nil, err
	}
	handleSheet, err := spritesheet.New(handles_png, 32)
	if err != nil {
		return nil, err
	}

	sb := genitemimage.New("Battle Axe")
	hs := sb.AddSpritesheet(headSheet, "Head")
	hs.OptionalEffects = []genitemimage.EffectType{
		genitemimage.EffectGlow,
		genitemimage.EffectCorrosion,
	}
	sb.AddSpritesheet(handleSheet, "Handle")

	sb.AddReplaceColor(defaultColorHead, genitemimage.ColorsMetal)
	sb.AddReplaceColor(defaultColorHandle, append(genitemimage.ColorsWood, genitemimage.ColorsMetal...))

	return sb, nil
}
