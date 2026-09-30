package utils

import (
	"io"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	SetupLogger(os.Stderr, "info")
}

// SetupLogger configures the package logger to write human readable output to w.
func SetupLogger(w io.Writer, level string) {
	zerolog.TimeFieldFormat = time.RFC3339
	log = zerolog.New(zerolog.ConsoleWriter{
		Out:        w,
		TimeFormat: time.RFC3339,
		NoColor:    !isTerminal(w),
	}).With().Timestamp().Logger().Level(ParseLevel(level))
}

func isTerminal(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	return isatty.IsTerminal(f.Fd()) || isatty.IsCygwinTerminal(f.Fd())
}

// ParseLevel maps a log level name to a zerolog level, defaulting to info.
func ParseLevel(level string) zerolog.Level {
	parsed, err := zerolog.ParseLevel(level)
	if err != nil || parsed == zerolog.NoLevel {
		return zerolog.InfoLevel
	}
	return parsed
}

func Debug(msg string) {
	log.Debug().Msg(msg)
}

func Info(msg string) {
	log.Info().Msg(msg)
}

func Warn(msg string) {
	log.Warn().Msg(msg)
}

func Error(err error) {
	log.Error().Err(err).Msg("error")
}
