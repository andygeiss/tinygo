package log

import (
	"log"
	"os"
	"sync"
	"time"
)

const (
	LevelDebug = 0
	LevelInfo  = 1
	LevelWarn  = 2
	LevelError = 3
	LevelFatal = 4
)

var (
	// Level represents the current state (Info ist default)
	Level int       = LevelInfo
	last  time.Time = time.Now()
	lock  sync.Mutex
)

// Debug ...
func Debug(msg string) {
	if Level <= LevelDebug {
		handle("DEBUG", msg)
	}
}

// Info ...
func Info(msg string) {
	if Level <= LevelInfo {
		handle("INFO", msg)
	}
}

// Warn ...
func Warn(msg string) {
	if Level <= LevelWarn {
		handle("WARN", msg)
	}
}

// Error ...
func Error(msg string) {
	if Level <= LevelError {
		handle("ERROR", msg)
	}
}

// Fatal ...
func Fatal(msg string) {
	if Level <= LevelFatal {
		handle("FATAL", msg)
		os.Exit(-1)
	}
}

func handle(level, msg string) {
	lock.Lock()
	diff := time.Since(last)
	last = time.Now()
	lock.Unlock()
	log.Printf("%-5s | %-16s | %-40s", level, diff, msg)
}
