package log

import (
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	//设置日志格式为json格式
	log.Set
}

func Logger(level string, payload error) {
	//log.Loger_level
	switch level {
	case "trace":
		
		
	}
}