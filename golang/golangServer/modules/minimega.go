package modules

import (
	"log"
	//"os"
)

type Minimega struct {
	Module
}

func ConfigureMinimega(path string) bool {
	log.Printf("Running Minimega Module")
	return true
}

func (minimega *Minimega) Configure(path string) bool {
	log.Printf("Running Minimega Module")
	return true
}

func (minimega *Minimega) GetName() string {
	return minimega.Module.Name
}