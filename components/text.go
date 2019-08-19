package components

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

type Text struct {
	txt   string
	color color.RGBA
}

func CreateTextLine(txt string, color color.RGBA) Text {
	return Text{
		txt:   txt,
		color: color,
	}
}

func WriteText(txt []Text, color color.RGBA, win *pixelgl.Window, mat pixel.Matrix) {
	face, err := loadTTF("./assets/font.ttf", 14)
	if err != nil {
		panic(err)
	}

	atlas := text.NewAtlas(face, text.ASCII)

	text := text.New(win.Bounds().Center(), atlas)
	text.LineHeight = atlas.LineHeight() * 1.5

	for _, t := range txt {
		text.Color = t.color
		text.Dot.X -= text.BoundsOf(t.txt).W() / 2
		fmt.Fprintln(text, t.txt)
	}

	text.Draw(win, mat)
}

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}