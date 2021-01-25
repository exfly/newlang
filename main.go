package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/exfly/newlang/eval"
	"github.com/exfly/newlang/object"
	"github.com/exfly/newlang/parser"
)

func main() {
	log.SetFlags(log.Lshortfile)

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	rootAstTree := parser.NewParser(string(buf))
	log.Printf("%v\n", toRaw(rootAstTree))

	env := object.NewEnvironment()
	log.Printf("%v\n", eval.Eval(rootAstTree, env))
}

func toRaw(obj interface{}) string {
	raw, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(raw)
}
