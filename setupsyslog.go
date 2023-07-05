package gologging

import (
	"log"
	"log/syslog"
)

func SetupSyslog(appName string) {
	logWriter, err := syslog.New(syslog.LOG_SYSLOG, appName)
	if err != nil {
		log.Fatalln("Unable to set logfile:", err.Error())
	}
	log.SetOutput(logWriter)
}
