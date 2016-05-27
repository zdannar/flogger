// +build !windows

package flogger

import "syscall"

// Redirects file descriptor to log file.
// Example: flogger.RedirectStream(flogger.STDERR)
func (f *Flogger) RedirectStream(fd int) (error) {
    err := syscall.Dup2(int(f.fd.Fd()), fd)
    return err
}

// Redirects stdout and stderr to log file.  Very handy to catch the output of
// panics in your log file.
func (f *Flogger) RedirectStreams() (error) {
    err := f.RedirectStream(STDOUT)
    if err != nil {
        return err
    }
    err = f.RedirectStream(STDERR)
    return err
}

// Calls f.RedirectStreams() on base logging object.
func RedirectStreams() (error) {
    err := defLogr.RedirectStreams()
    return err
}
