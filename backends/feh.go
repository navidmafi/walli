package backends

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/log"
)

const FEH_EXEC_BIN = "/usr/bin/feh"

type fehBackend struct{}

func (b *fehBackend) Apply(buf bytes.Buffer) error {
	name := fmt.Sprintf("walli-%d.jpg", time.Now().UnixMilli())
	path := "/tmp/" + name

	if err := os.WriteFile(path, buf.Bytes(), 0644); err != nil {
		return fmt.Errorf("feh: could not write buffer to %s: %w", path, err)
	}

	return b.ApplyFile(path)
}

func (b *fehBackend) ApplyFile(filename string) error {
	cmd := exec.Command(FEH_EXEC_BIN, "--bg-scale", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("feh: failed to set wallpaper: %w", err)
	}

	log.Debug("feh set wallpaper to", filename)
	return nil
}
