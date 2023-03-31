package main

import (
	"encoding/base64"
	"log"
	"os"

	"github.com/ugurkinik/utility-tools/internal"
)

func main() {
	if len(os.Args) == 2 {
		log.Println(base64decode(os.Args[1]))
	} else if len(os.Args) > 2 {
		internal.UpdateSelectedText(base64decode, os.Args[1], os.Args[2:])
	} else {
		log.Fatalf("ERROR: invalid format\nValid formats:\n\n<tool-name> <input>\n<tool-name> <file> <line-start>,<column-start>,<line-end>,<column-end> ...\n")
	}
}

func base64decode(input string) string {
	result, err := base64.StdEncoding.DecodeString(input)

	if err != nil {
		log.Fatalf("ERROR: %s\n", err)
		return input
	}

	return string(result)
}
