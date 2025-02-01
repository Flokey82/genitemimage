package staff

import (
	"image/color"

	_ "embed"

	"github.com/Flokey82/genitemimage"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

//go:embed sprites/heads_32_1x4.png
var heads_png []byte

//go:embed sprites/handles_32_1x3.png
var handles_png []byte

// #fbf236
var defaultColorHead = color.RGBA{R: 0xfb, G: 0xf2, B: 0x36, A: 0xff}

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

	sb := genitemimage.New("Staff")

	hs := sb.AddSpritesheet(handleSheet, "Handle")
	hs.OptionalEffects = []genitemimage.EffectType{
		genitemimage.EffectFlame,
		genitemimage.EffectDrip,
		genitemimage.EffectGlow,
	}
	sb.AddSpritesheet(headSheet, "Head")

	sb.AddReplaceColor(defaultColorHead, genitemimage.ColorsGem)
	sb.AddReplaceColor(defaultColorHandle, genitemimage.ColorsWood)

	return sb, nil
}
