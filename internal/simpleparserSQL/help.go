package simpleparserSQL

import (
	"fmt"

	"github.com/1r0npipe/search-In-CSV/cmd/server/config"
)
const HELP = `To use that tool, you can perform those commands below:
		"show config" will pring current configuration with info about CSV file and log files
		"show map" (debug) will print current content of the uploaded file
		"WHERE {conditions} you can type any relevant condition against the CSV content, example:
		"WHERE cont == "Asia" AND country == "China" OR (country == "China" AND people > 10000"`

func PrintConfig(conf  *config.Config) {
	fmt.Printf("Server config - ip: %s, port: %s\n", conf.Server.Host, conf.Server.Port)
	fmt.Printf("Timeout for searcing in sec: %v\n", conf.Timeout)
	fmt.Printf("Loaded CSV file: %s\n", conf.DefaultFile)
}