package apilogger

import (
	log "github.com/google/logger"
	"io/ioutil"
)

var Logger *log.Logger

func init() {
	Logger = log.Init("API Demo", true, false, ioutil.Discard)
}
