package backends

import (
	"navidmafi/walli/logger"
	"os"
)

func ObtainDesktopEnvironment() {
	var DE = os.Getenv("XDG_CURRENT_DESKTOP")

	if len(DE) == 0 {
		logger.Logger.Fatal("Could not detect desktop environment")
	}

	logger.Logger.Debug(DE)

}
