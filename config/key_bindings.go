package config

import (
	"azul3d.org/keyboard.v1"
	"azul3d.org/mouse.v1"
	"encoding/yaml"
	"io/ioutil"
	"log"
	"reflect"
	"coopersoft/config/actionslist"
)

func absPath(relPath string) string {
	if len(examplesDir) == 0 {
		// Find assets directory.
		for _, path := range filepath.SplitList(build.Default.GOPATH) {
			path = filepath.Join(path, "src/coopersoft/config")
			if _, err := os.Stat(path); err == nil {
				examplesDir = path
				break
			}
		}
	}
	return filepath.Join(examplesDir, relPath)
}

func getBytesFromFile(path string) []byte {
	bytes, err := ioutil.ReadFile(absPath(path))
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

type Action struct {
	Name string
	Identifier int
	Duration float64
}

func NewAction(name string, ident int, duration float64) *Action {
	return &Action{Name: name, Identifier: ident, Duration: duration}
}

type Button interface {
	Name() string
}

type KeyButton struct {
	Name string
	Button keyboard.Key
}

type MouseButton struct {
	Name string
	Button mouse.Button
}

func(k *KeyButton) Name() string {
	return k.Name
}

func(m *MouseButton) Name() string {
	return m.Name
}



var (
	Bindings map[Action]Button

	MoveUp = NewAction("MoveUp", 0, 0)
	MoveDown = NewAction("MoveDown": 1, 0)
	MoveLeft = NewAction("MoveLeft", 2, 0) 
	MoveRight 
	Pause
	Quit
	Attack1 
	Attack2 
)

func ResetDefaultBindings(){

}

func SetBinding(a Action, b Button) {
	Bindings[a] = b
}