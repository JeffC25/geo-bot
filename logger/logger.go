package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func Logger(level int) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	var output io.Writer

	output = zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	return zerolog.New(output).With().Caller().Timestamp().Logger().Level(zerolog.Level(level))
}
