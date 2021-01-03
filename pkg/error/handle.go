package error_handler

import (
	"log"
	"runtime"
)

func Print(err error) {
	pc, _, line, _ := runtime.Caller(1)
	log.Printf("[error] in [%s:%d] - %v", runtime.FuncForPC(pc).Name(), line, err)
}
