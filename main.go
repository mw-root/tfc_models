package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/jsonschema"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var schemaFolder string = "./jsonschemas"

func ensureSchemaFolder() {
	if _, err := os.Stat(schemaFolder); os.IsNotExist(err) {
		err := os.Mkdir(schemaFolder, 0755)
		check(err)
	}
}

func main() {

	m := getStructs()

	ensureSchemaFolder()

	for file, iface := range m {

		schema := jsonschema.Reflect(iface)
		jschema, _ := json.MarshalIndent(schema, "", "    ")
		filename := schemaFolder + "/" + file + ".json"

		fmt.Printf("Writing %s\n", filename)
		err := ioutil.WriteFile(filename, jschema, 0644)
		check(err)
	}

}
