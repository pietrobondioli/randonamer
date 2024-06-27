package util

import (
	"io"
	"log"
	"os"
)

var debug bool

func SetDebug(enabled bool) {
	debug = enabled
	if debug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(io.Discard)
	}
}

func DebugLog(format string, v ...interface{}) {
	if debug {
		log.Printf("[DEBUG] "+format, v...)
	}
}
