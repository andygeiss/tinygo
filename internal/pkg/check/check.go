// package check handles error handling by logging the error message to stderr.
package check

import "github.com/andygeiss/tinygo/internal/pkg/log"

// Debug ...
func Debug(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

// Error ...
func Error(err error) {
	if err != nil {
		log.Error(err.Error())
	}
}

// Fatal ...
func Fatal(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// Info ...
func Info(err error) {
	if err != nil {
		log.Info(err.Error())
	}
}

// Warn ...
func Warn(err error) {
	if err != nil {
		log.Warn(err.Error())
	}
}
