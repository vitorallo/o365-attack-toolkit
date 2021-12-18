package main

import (
	_ "database/sql"
	"flag"
	"fmt"

	//"github.com/sirupsen/logrus"

	"github.com/vitorallo/o365-attack-toolkit/logging"
	"github.com/vitorallo/o365-attack-toolkit/model"
	"github.com/vitorallo/o365-attack-toolkit/server"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gcfg.v1"
)

var (
	configFile  = flag.String("c", "template.conf", "Configuration template")
	LogOutput   = flag.String("o", "", "Write session and logs to a file")
	IntServerUp = flag.Bool("i", true, "Enable internal HTTP GUI server <default: true>")
	ApiServerUp = flag.Bool("a", false, "Enable REST API GUI server <default: false>")
	debug_mode  = flag.String("d", "error", "Set debug level <error: default, debug, trace>")
)

func main() {

	flag.Parse()

	model.GlbConfig = model.Config{}
	err := gcfg.ReadFileInto(&model.GlbConfig, *configFile)

	if err != nil {
		fmt.Printf("Fatal error reding config file: %v", err.Error())
		return
	}

	//setting logging environment
	Log := logging.NewLogger(*debug_mode)
	Log.Trace(model.GlbConfig)

	//initializeRules()

	//staring external server
	Log.Debug("Starting with oauth token redirect URI: ", model.GlbConfig.Oauth.Redirecturi)
	go server.StartExtServer(model.GlbConfig, logging.GetLogger())

	if *ApiServerUp {
		Log.Debug("Starting REST API GUI")
		server.StartAPIServer(model.GlbConfig, logging.GetLogger())
	} else if !*IntServerUp {
		Log.Fatal("Internal or API GUI, one of the two must be enabled!")
	} else {
		Log.Debug("Starting internal HTTP GUI")
		server.StartIntServer(model.GlbConfig)
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
