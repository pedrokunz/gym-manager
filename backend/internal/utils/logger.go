package utils

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "GYM-APP: ", log.LstdFlags|log.Lshortfile)
