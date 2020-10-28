package modules

import (
	"log"
	//"os"
)

type Parser struct {
	Module
}

func ConfigureParser(path string) bool {
	log.Printf("Running Parser Module")
	return true
}

func (Parser *Parser) Configure(path string) bool {
	log.Printf("Running Parser Module")
	return true
}

func (parser *Parser) GetName() string {
	return parser.Module.Name
}