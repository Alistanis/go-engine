package player 

import (
	"go/build"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"azul3d.org/gfx.v1"
	"azul3d.org/keyboard.v1"
	"azul3d.org/lmath.v1"
)

type Player struct {
	Card *gfx.Object
}

func NewPlayer() *Player {
	// Create a simple shader.
	shader := gfx.NewShader("SimpleShader")
	shader.GLSLVert = glslVert
	shader.GLSLFrag = glslFrag

	// Load the picture.
	f, err := os.Open(absPath("assets/textures/texture_coords_1024x1024.png"))
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	// Create new texture.
	tex := gfx.NewTexture()
	tex.Source = img
	tex.MinFilter = gfx.LinearMipmapLinear
	tex.MagFilter = gfx.Linear
	tex.Format = gfx.DXT1RGBA

	// Create a card mesh.
	cardMesh := gfx.NewMesh()
	cardMesh.Vertices = []gfx.Vec3{
		// Bottom-left triangle.
		{-1, 0, -1},
		{1, 0, -1},
		{-1, 0, 1},

		// Top-right triangle.
		{-1, 0, 1},
		{1, 0, -1},
		{1, 0, 1},

	}
	cardMesh.TexCoords = []gfx.TexCoordSet{
		{
			Slice: []gfx.TexCoord{
				{0, 1},
				{1, 1},
				{0, 0},

				{0, 0},
				{1, 1},
				{1, 0},
			},
		},
	}
	// Create a card object.
	card := gfx.NewObject()
	card.AlphaMode = gfx.AlphaToCoverage
	card.Shader = shader
	card.Textures = []*gfx.Texture{tex}
	card.Meshes = []*gfx.Mesh{cardMesh}
	card.SetPos(lmath.Vec3{400, 400, 400})
	return &Player{Card: card}
}

func(p *Player) X() float64 {
	return Card.Pos().X
}

func(p *Player) Y() float64 {
	return Card.Pos().Y
}

func(p *Player) Z() float64 {
	return Card.Pos().Z
}