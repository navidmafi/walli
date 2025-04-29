package backends

import "bytes"

type BackendName string

const (
	Swww     BackendName = "swww"
	Gnome    BackendName = "gnome"
	Plasma   BackendName = "plasma"
	Feh      BackendName = "feh"
	MvpPaper BackendName = "mvppaper"
)

type Backend interface {
	ApplyFile(filename string) error

	Apply(image bytes.Buffer) error
}

var Backends = map[BackendName]Backend{
	Swww:   &swwwBackend{},
	Gnome:  &gnomeBackend{},
	Plasma: &plasmaBackend{},
	Feh:    &fehBackend{},
}

func GetAvailable() []string {
	var result []string

	for key := range Backends {
		result = append(result, string(key))
	}

	return result
}
