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
	"coopersoft/player/state"
)


var (
	CurrentState *PlayerState
	LastState *PlayerState
)

type Player struct {
	Card *gfx.Object
}

func(p *Player) HandleInput(e keyboard.StateEvent) {
	CurrentState.HandleInput(e)
}

type PlayerState interface {
	HandleInput(e keyboard.StateEvent)
	ExitState(ps *PlayerState)
}

type PlayerAction struct {
	Name string
	Identifier int
	Function func()
}

type Direction interface {
	Direction() lmath.Vec3
}

type North struct {
	Vec3 lmath.Vec3
}

func(n North)Direction() lmath.Vec3 {
	return Vec3{X: 0, Y: 0, Z: -1}
}

type NorthEast struct {
	Vec3 lmath.Vec3
}

type East struct {
	Vec3 lmath.Vec3
}

type SouthEast struct {
	Vec3 lmath.Vec3
}

type South struct {
	Vec3 lmath.Vec3
}

type SouthWest struct {
	Vec3 lmath.Vec3
}

type West struct {
	Vec3 lmath.Vec3
}

type NorthWest struct {
	Vec3 lmath.Vec3
}


type Moving interface {
	MovePlayer(p *Player, d *Direction)
}

func(m Moving)HandleInput(e keyboard.StateEvent) {
	if ev.Key == keyboard.W {
		fmt.Println(ev)
		if ev.State == keyboard.Down {
			moving = true
		}
		if ev.State == keyboard.Up {
			moving = false
		}
	}
}

func(p *PlayerAction) Name() string {
	return p.Name
}

func(p *PlayerAction) Identifier() int {
	return p.Identifier
}

func(p *PlayerAction) Exec() {
	Function()
}

func NewPlayerAction(name string, ident int, f func()) *PlayerAction {
	return &PlayerAction{Name: name, Identifier: ident, Function: f}
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