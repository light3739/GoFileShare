package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	Log = zerolog.New(os.Stderr).With().Timestamp().Logger()
}
func LogError(err error, msg string) {
	Log.Error().Err(err).Msg(msg)
}

func LogInfo(msg string) {
	Log.Info().Msg(msg)
}
