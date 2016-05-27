package flogger

import (
    "log"
    "os"
)

const (
    DEBUG int = iota
    INFO
    WARNING
    ERROR
    FATAL
    PANIC
)

const (
    STDOUT = iota + 1
    STDERR
)

var (
    FLOG_APPEND int = os.O_WRONLY|os.O_APPEND|os.O_CREATE
    FLOG_FORMAT int = log.LstdFlags|log.Lshortfile

    FLOG_LEVELS = map[int] string {
        DEBUG : "DEBUG",
        INFO : "INFO",
        WARNING : "WARNING",
        ERROR : "ERROR",
        FATAL : "FATAL",
        PANIC : "PANIC",
    }
    defLogr *Flogger = New(DEBUG, FLOG_FORMAT, FLOG_LEVELS)
)

// Closes default logger's file descriptor.
func Close() error {
    return defLogr.Close()
}

// Opens default logger's file for writing.
func OpenFile(logPath string, openMode int, perms os.FileMode) (error) {
    return defLogr.OpenFile(logPath, openMode, perms)
}

// Sets the default logger's logging level.
func SetLevel(level int) {
    defLogr.SetLevel(level)
}

// Sets the default logger's level map.
func SetLevelMap(lmap map[int]string) {
    defLogr.SetLevelMap(lmap)
}

// Calls f.Debug() on base logging object.
func Debug(args ...interface{}) {
   defLogr.Debug(args...)
}

// Calls f.Info() on base logging object.
func Info(args ...interface{}) {
   defLogr.Info(args...)
}

// Calls f.Warning() on base logging object.
func Warning(args ...interface{}) {
   defLogr.Warning(args...)
}

// Calls f.Error() on base logging object.
func Error(args ...interface{}) {
   defLogr.Error(args...)
}

// Calls f.Fatal() on base logging object.
func Fatal(args ...interface{}) {
   defLogr.Fatal(args...)
}

// Calls f.Panic() on base logging object.
func Panic(args ...interface{}) {
   defLogr.Panic(args...)
}

// Calls f.Debugf() on base logging object.
func Debugf(msg string, args ...interface{}) {
   defLogr.Debugf(msg, args...)
}

// Calls f.Infof() on base logging object.
func Infof(msg string, args ...interface{}) {
   defLogr.Infof(msg, args...)
}

// Calls f.Warningf() on base logging object.
func Warningf(msg string, args ...interface{}) {
   defLogr.Warningf(msg, args...)
}

// Calls f.Errorf() on base logging object.
func Errorf(msg string, args ...interface{}) {
   defLogr.Errorf(msg, args...)
}

// Calls f.Fatalf() on base logging object.
func Fatalf(msg string, args ...interface{}) {
   defLogr.Fatalf(msg, args...)
}

// Calls f.Panicf() on base logging object.
func Panicf(msg string, args ...interface{}) {
   defLogr.Panicf(msg, args...)
}
