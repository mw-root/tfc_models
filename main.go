package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/alecthomas/jsonschema"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	m := getStructs()

	for file, iface := range m {
		schema := jsonschema.Reflect(iface)
		jschema, _ := json.MarshalIndent(schema, "", "    ")
		filename := file + ".json"
		fmt.Printf("Writing %s\n", filename)
		err := ioutil.WriteFile(filename, jschema, 0644)
		check(err)
	}

}
