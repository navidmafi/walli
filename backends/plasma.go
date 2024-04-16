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

const BASH_EXEC_BIN string = "/usr/bin/bash"

type plasmaBackend struct{}

func (b *plasmaBackend) Apply(buf bytes.Buffer) error {

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

func (b *plasmaBackend) ApplyFile(filename string) error {
	bashCmd := fmt.Sprintf(`#!/bin/sh
qdbus org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript "
    var allDesktops = desktops();
    print (allDesktops);
    for (i=0;i<allDesktops.length;i++) {
        d = allDesktops[i];
        d.wallpaperPlugin = 'org.kde.image';
        d.currentConfigGroup = Array('Wallpaper',
                                    'org.kde.image',
                                    'General');
        d.writeConfig('Image', 'file://%s')
    }"`, filename)

	cmd := exec.Command(BASH_EXEC_BIN, "-c", bashCmd)

	stdout, err := cmd.Output()

	if err != nil {
		return err
	}

	log.Debug(stdout)

	return nil

}
