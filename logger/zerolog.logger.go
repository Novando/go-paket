package logger

import (
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {

	// ConsoleLoggingEnabled console logging
	ConsoleLoggingEnabled bool

	// EncodeLogsAsJson makes the log framework log JSON
	EncodeLogsAsJson bool

	// FileLoggingEnabled makes the framework log to a file.go
	// the fields below can be skipped if this value is false!
	FileLoggingEnabled bool

	// Directory to log to when filelogging is enabled
	Directory string

	// Filename is the name of the logfile which will be placed inside the directory
	Filename string

	// MaxSize the max size in MB of the logfile before it's rolled
	MaxSize int

	// MaxBackups the max number of rolled files to keep
	MaxBackups int

	// MaxAge the max age in days to keep a logfile
	MaxAge int

	// CallerSkip the number of directory hierarchy to be skipped
	CallerSkip int

	// LowestLevel the lowest log type that will be printed
	LowestLevel zerolog.Level
}

type Logger *zerolog.Logger

// InitZerolog call it during cmd main.go
func InitZeroLog(config ...Config) *zerolog.Logger {
	if len(config) == 0 {
		config = append(config, Config{
			ConsoleLoggingEnabled: true,
			LowestLevel:           zerolog.TraceLevel,
		})
	}
	var writers []io.Writer

	if config[0].ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}
	if config[0].FileLoggingEnabled {
		writers = append(writers, newRollingFile(config[0]))
	}
	if config[0].CallerSkip == 0 {
		config[0].CallerSkip = 3
	}

	zerolog.CallerSkipFrameCount = config[0].CallerSkip
	mw := io.MultiWriter(writers...)
	zerolog.SetGlobalLevel(config[0].LowestLevel)

	logger := zerolog.New(mw).With().Timestamp().Caller().Logger()

	logger.Info().
		Bool("fileLogging", config[0].FileLoggingEnabled).
		Bool("jsonLogOutput", config[0].EncodeLogsAsJson).
		Str("logDirectory", config[0].Directory).
		Str("fileName", config[0].Filename).
		Int("maxSizeMB", config[0].MaxSize).
		Int("maxBackups", config[0].MaxBackups).
		Int("maxAgeInDays", config[0].MaxAge)

	return &logger
}

func newRollingFile(config Config) io.Writer {
	l := &lumberjack.Logger{
		Filename:   path.Join(config.Directory, config.Filename),
		MaxBackups: config.MaxBackups, // files
		MaxSize:    config.MaxSize,    // megabytes
		MaxAge:     config.MaxAge,     // days
	}

	return l
}
