package backends

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/charmbracelet/log"
)

const SWWW_EXEC_BIN string = "/usr/bin/swww"

type swwwBackend struct{}

func (b *swwwBackend) Apply(buf bytes.Buffer) error {

	cmd := exec.Command(SWWW_EXEC_BIN, "img", "-")

	cmd.Stdin = &buf
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
	// cmd := exec.Command("s")
}

func (b *swwwBackend) ApplyFile(filename string) error {

	cmd := exec.Command(SWWW_EXEC_BIN, "img", filename)

	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	log.Debug(stdout)

	return nil

}
