package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	write   io.Writer
}

func NewLogger(prefix string) *Logger {
	write := io.Writer(os.Stdout)
	logger := log.New(write, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(write, "[DEBUG] ", logger.Flags()),
		info:    log.New(write, "[INGO] ", logger.Flags()),
		warning: log.New(write, "[WARNING] ", logger.Flags()),
		err:     log.New(write, "[ERROR] ", logger.Flags()),
		write:   write,
	}
}

func (log *Logger) Debug(v ...interface{}) {
	log.debug.Println(v...)
}

func (log *Logger) Info(v ...interface{}) {
	log.info.Println(v...)
}

func (log *Logger) Warning(v ...interface{}) {
	log.warning.Println(v...)
}

func (log *Logger) Error(v ...interface{}) {
	log.err.Println(v...)
}

//----

func (log *Logger) Debugf(format string, v ...interface{}) {
	log.debug.Printf(format, v...)
}
func (log *Logger) Infof(format string, v ...interface{}) {
	log.info.Printf(format, v...)
}
func (log *Logger) Warningf(format string, v ...interface{}) {
	log.warning.Printf(format, v...)
}
func (log *Logger) Errorf(format string, v ...interface{}) {
	log.err.Printf(format, v...)
}
