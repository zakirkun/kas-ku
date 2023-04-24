package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger
var debug bool

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	debug = true
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	Logger = zerolog.New(os.Stdout)
}
