package main

import (
	_ "database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/vitorallo/o365-attack-toolkit/model"
	"github.com/vitorallo/o365-attack-toolkit/server"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gcfg.v1"
)

var config_file = flag.String("c", "template.conf", "Configuration template")
var ext_server_up = flag.Bool("e", false, "Bring up the External server")
var int_server_up = flag.Bool("i", false, "Bring up the Internal server")
var debug_mode = flag.Bool("d", false, "Enable debug mode")

func main() {

	flag.Parse()

	model.GlbConfig = model.Config{}
	err := gcfg.ReadFileInto(&model.GlbConfig, *config_file)

	if err != nil {
		log.Fatal(err.Error())
	}

	//initializeRules()
	if *ext_server_up {
		log.Println("OAuth token redirect URI:", model.GlbConfig.Oauth.Redirecturi)
		go server.StartExtServer(model.GlbConfig)
	} else {
		fmt.Println("Starting up with no external server...")
	}

	if *int_server_up {
		server.StartIntServer(model.GlbConfig)
	} else {
		fmt.Println("Starting up with no internal server...")
	}

	//fmt.Println(model.GlbConfig)
}

/**
func initializeRules() {

	var ruleFiles []string
	var tempRule model.Rule

	root := "rules"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ruleFiles = append(ruleFiles, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range ruleFiles {

		ruleFile, err := os.Open(file)

		if err != nil {
			log.Println(err)
		}

		defer ruleFile.Close()

		byteValue, _ := ioutil.ReadAll(ruleFile)

		json.Unmarshal(byteValue, &tempRule)

		model.GlbRules = append(model.GlbRules, tempRule)

	}

	log.Printf("Loaded %d rules successfully.", len(model.GlbRules))
}
**/
