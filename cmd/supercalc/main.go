package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/exfly/newlang/parser"
)

func main() {
	log.SetFlags(log.Lshortfile)

	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		parser.NewParser(line)
	}
}
