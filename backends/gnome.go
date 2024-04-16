package backends

import (
	"bytes"
	"fmt"
	"navidmafi/walli/logger"
	"os"
	"os/exec"
	"time"

	"github.com/charmbracelet/log"
)

const GNOME_EXEC_BIN string = "/usr/bin/gsettings"

type gnomeBackend struct{}

func (b *gnomeBackend) Apply(buf bytes.Buffer) error {

	fileName := fmt.Sprintf("walli%d", time.Now().UnixMilli())
	filePath := fmt.Sprintf("/tmp/%s", fileName)
	writeErr := os.WriteFile(filePath, buf.Bytes(), 0777)

	if writeErr != nil {
		logger.Logger.Fatal("Could not write buffer to disk %s", writeErr)
	}

	logger.Logger.Debug("Wrote buffer to disk")

	logger.Logger.Debug(filePath)
	err := b.ApplyFile(filePath)

	if err != nil {
		return err
	}

	return nil
}

func (b *gnomeBackend) ApplyFile(filename string) error {

	fileURI := fmt.Sprintf("file://%s", filename)
	logger.Logger.Debug(fileURI)
	cmdLight := exec.Command(GNOME_EXEC_BIN, "set", "org.gnome.desktop.background", "picture-uri", fileURI)
	cmdDark := exec.Command(GNOME_EXEC_BIN, "set", "org.gnome.desktop.background", "picture-uri-dark", fileURI)

	stdoutLight, errLight := cmdLight.Output()
	stdoutDark, errDark := cmdDark.Output()

	if errDark != nil {
		return errDark
	}
	if errLight != nil {
		return errLight
	}

	log.Debug(stdoutLight)
	log.Debug(stdoutDark)

	return nil

}
