package config

import (
	"azul3d.org/keyboard.v1"
	"azul3d.org/mouse.v1"
	"encoding/yaml"
	"io/ioutil"
	"log"
	"coopersoft/player"
)

const (
	MoveUp = iota
	MoveDown 
	MoveLeft 
	MoveRight 
	Pause
	Quit
	Attack1 
	Attack2 
)

const (

)

var (
	
	Actions = [...]*Action { 
		NewAction("MoveUp", MoveUp), 
		NewAction("MoveDown", MoveDown), 
		NewAction("MoveLeft", MoveLeft),
		NewAction("MoveRight", MoveRight),
		NewAction("Pause", Pause),
		NewAction("Quit", Quit),
		NewAction("Attack1", Attack1),
		NewAction("Attack2", Attack2),
		}

	DefaultBindings := map[*Button]*Action{
		NewKeyButton("MoveUp")
	}
	
	bindingsDir string


)


func init() {
	Bindings = make(map[*Button]*Action)
}

func absPath(relPath string) string {
	if len(bindingsDir) == 0 {
		// Find assets directory.
		for _, path := range filepath.SplitList(build.Default.GOPATH) {
			path = filepath.Join(path, "src/coopersoft/config")
			if _, err := os.Stat(path); err == nil {
				examplesDir = path
				break
			}
		}
	}
	return filepath.Join(bindingsDir, relPath)
}

func getBytesFromFile(path string) []byte {
	bytes, err := ioutil.ReadFile(absPath(path))
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

type Action interface {
	Name() string
	Identifier() int
	Exec func()
}

type Button interface {
	Name() string
	Action() *Action
}

type KeyButton struct {
	Name string
	Binding keyboard.Key
	Act *Action
}

func NewKeyButton(s string, k keyboard.Key) *KeyButton {
	return &KeyButton{Name: s, Button: k}
}

type MouseButton struct {
	Name string
	Binding mouse.Button
	Act *Action
}

func NewMouseButton(s string, m mouse.Button) *MouseButton {
	return &MouseButton{Name: s, Button: m}
}

func(k *KeyButton) Name() string {
	return k.Name
}

func(m *MouseButton) Name() string {
	return m.Name
}

func ResetDefaultBindings(){
	for i := range Actions; i < len(Actions); i++ {

	}
}

func SetBinding(b *Button, a *Action) {
	Bindings[b] = a
}