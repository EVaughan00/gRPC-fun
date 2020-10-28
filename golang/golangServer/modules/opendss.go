package modules

import (
	"log"
)

type OpenDSS struct {
	Module
}

func ConfigureOpenDSS(path string) bool {
	log.Printf("Running OpenDSS Module")
	return true
}

func (opendss *OpenDSS) Configure(path string) bool {
	log.Printf("Running OpendSS Module")
	return false
}

func (opendss *OpenDSS) GetName() string {
	return opendss.Module.Name
}