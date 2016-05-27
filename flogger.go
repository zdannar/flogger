/*
Flogger provides a simple wrapper around go's base logging library to provide
logging levels.  Further documentation can be found at
https://github.com/zdannar/flogger
*/
package flogger

import (
    "log"
    "os"
    "fmt"
)

// Returns a flogger.Flogger structure for logging.
func New(logLevel, logFormat int, levelMap map[int]string) *Flogger {
    return &Flogger{ fd: os.Stdout,
                     level: logLevel,
                     logForm: logFormat,
                     log : log.New(os.Stderr, "", FLOG_FORMAT),
                     lvlMap: levelMap }
}

// Flogger structure.
type Flogger struct {
    lpath string
    fd *os.File
    log *log.Logger
    level int
    logForm int
    lvlMap map[int]string
}

// Opens a file for flogger to log to.
func (f *Flogger) OpenFile(logPath string, openMode int, perms os.FileMode) (error) {

    f.lpath = logPath
    fd, err := os.OpenFile(f.lpath, openMode, perms)
    if err != nil {
        return err
    }
    f.log, f.fd = log.New(fd, "", f.logForm), fd
    return nil
}

// Flushes and closes file descriptor that flogger is logging to.  Should always be
// called at the end of execution.
func (f *Flogger) Close() error {
    if err := f.fd.Sync(); err != nil {
        return err
    }
    return f.fd.Close()
}

// Sets the logger level. Usually called with flogger.INFO, flogger.DEBUG, etc...
func (f *Flogger) SetLevel(level int) {
    f.level = level
}

// Sets logging level map. Provides the abilty to change level messages.
// Default logger uses flogger.FLOG_LEVELS
func (f *Flogger) SetLevelMap(lmap map[int]string) {
    //TODO: Add level checking
    f.lvlMap = lmap
}

// Provides the mapping for text of the log levels
func (f *Flogger) appLevel(level int, msg string) string {
    return fmt.Sprintf(": %s : ", f.lvlMap[level]) + msg
}

// Base method that actually applies the levels and wraps log.Print,
// log.Fatal, log.Panic
func (f *Flogger) flog(level int, args ...interface{}) {
    if level < f.level { return }
    ml := []interface{}{f.appLevel(level, "")}
    v := append(ml,args...)
    s := fmt.Sprint(v...)
    f.log.Output(4, s)
    switch {
        case level == FATAL:
            os.Exit(1)
        case level == PANIC:
            panic(s)
    }
}

// Base method that actually applies the levels and wraps log.Printf,
// log.Fatalf, log.Panicf
func (f *Flogger) flogf(level int, msg string, args ...interface{}) {
    if level < f.level { return }

    ml := f.appLevel(level,msg)
    f.log.Output(4, fmt.Sprintf(ml, args...))

    switch {
        case level == FATAL:
            os.Exit(1)
        case level == PANIC:
            panic(msg)
    }
}

// Logs to the debugging logging channel
func (f *Flogger) Debug(args ...interface{}) {
   f.flog(DEBUG, args...)
}

// Logs to the info logging channel
func (f *Flogger) Info(args ...interface{}) {
   f.flog(INFO, args...)
}

// Logs to the warning logging channel
func (f *Flogger) Warning(args ...interface{}) {
   f.flog(WARNING, args...)
}

// Logs to the error logging channel
func (f *Flogger) Error(args ...interface{}) {
   f.flog(ERROR, args...)
}

// Logs your message, then calls os.Exit(1).
func (f *Flogger) Fatal(args ...interface{}) {
   f.flog(FATAL, args...)
}

// Panics with your message.
func (f *Flogger) Panic(args ...interface{}) {
   f.flog(PANIC, args...)
}


// Debug function with string formatting
func (f *Flogger) Debugf(msg string, args ...interface{}) {
   f.flogf(DEBUG, msg, args...)
}

// Info function with string formatting
func (f *Flogger) Infof(msg string, args ...interface{}) {
   f.flogf(INFO, msg, args...)
}

// Warning function with string formatting
func (f *Flogger) Warningf(msg string, args ...interface{}) {
   f.flogf(WARNING, msg, args...)
}

// Error function with string formatting
func (f *Flogger) Errorf(msg string, args ...interface{}) {
   f.flogf(ERROR, msg, args...)
}

// Fatal function with string formatting
func (f *Flogger) Fatalf(msg string, args ...interface{}) {
   f.flogf(FATAL, msg, args...)
}

// Panic function with string formatting
func (f *Flogger) Panicf(msg string, args ...interface{}) {
   f.flogf(PANIC, msg, args...)
}
