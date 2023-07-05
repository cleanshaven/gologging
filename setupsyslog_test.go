package gologging

import (
	"log"
	"testing"
)

func TestSetupSyslog(t *testing.T) {
	SetupSyslog("TestSetupSyslog")
	log.Println("Testing SetupSyslog")
}
