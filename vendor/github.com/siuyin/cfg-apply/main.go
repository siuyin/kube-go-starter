// cfg-apply takes a configuration yaml, a configuration template and
// renders on os.Stdout.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

func main() {
	m := make(map[interface{}]interface{})
	if len(os.Args) < 3 {
		fmt.Printf(`Usage %s {config.yaml} {template.yaml}
Eg:
    %s cfg_local.yaml deploy.template.yaml > deploy.local.yaml
`, os.Args[0], os.Args[0])
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	dat, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(dat, &m)

	t := template.Must(template.ParseFiles(os.Args[2]))
	if err := t.Execute(os.Stdout, m); err != nil {
		log.Fatal(err)
	}
}
