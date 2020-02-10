package npwut

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Flags struct {
	Dir      string
	Help     bool
	Verboose bool
	All      bool
}

func parseFlags() Flags {
	f := Flags{}
	flag.StringVar(&f.Dir, "d", "", "Return npm run commands in specific directory")
	flag.BoolVar(&f.All, "a", false, "Return all npm run commands")
	flag.BoolVar(&f.Help, "h", false, "Instructions for use")
	flag.BoolVar(&f.Verboose, "v", false, "Enable logging")

	flag.Parse()
	return f
}

func parseCommand() string {
	var output string
	for _, arg := range os.Args {
		if !strings.Contains(arg, "-") {
			output = arg
		}
	}
	return output
}

func usage() {
	fmt.Fprintf(os.Stdout, `Usage:	npmwut [-adv] build
				npmwut <command> - Return npm run command with the given command string
				npmwut -d - Return npm run commands in specific directory
				npmwut -a - Return all npm run commands
				npmwut -v - Verboose mode
	`)
}

func findPackageJSON(dir string) []byte {
	data, err := ioutil.ReadFile(dir)
	if err != nil {
		log.Fatal("Error reading package.json file or no file here: ", err)
	}
	return data
}

func readPackageJSON(data []byte) (scripts map[string]string) {
	var f interface{}
	err := json.Unmarshal(data, &f)
	if err != nil {
		log.Fatal("Error Unmarshalling package.json file: ", err)
	}
	m := f.(map[string]interface{})

	scripts = make(map[string]string)

	for k, v := range m {
		if k == "scripts" {
			nested, ok := v.(map[string]interface{})
			if !ok {
				continue
			}
			for nestedK, v := range nested {
				s, ok := v.(string)
				if !ok {
					continue
				}
				scripts[nestedK] = s
			}
		}
	}
	return scripts
}

func main() {
	var command string
	targetFile := "package.json"
	pwd, _ := os.Getwd()

	if len(os.Args) > 1 {
		command = parseCommand()
	}

	if command == "" {
		log.Fatalf("No npm run command in args")
	}

	// fmt.Println(verboose)
	// fmt.Println(all)
	// fmt.Println(help)
	// fmt.Println(command)
	Flags := parseFlags()

	if Flags.Help {
		usage()
		os.Exit(1)
	}

	filePath := pwd + "/" + Flags.Dir + "/" + targetFile

	data := findPackageJSON(filePath)
	if data == nil {
		// scan file directories
		//TODO: update to look for file first before stepping through files in the directory.
		files, err := ioutil.ReadDir(".")
		if err != nil {
			log.Fatal("Error reading directory: ", err)
		}
		fmt.Println(files)
	}

	scripts := readPackageJSON(data)

	if scripts[command] == "" || Flags.All {
		fmt.Printf("%#v\n", scripts)
	} else {
		fmt.Println(scripts[command])
	}

}
