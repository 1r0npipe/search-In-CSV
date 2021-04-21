package main

import (
	"fmt"
	"github.com/1r0npipe/search-In-CSV/cmd/server/config"
	"github.com/1r0npipe/search-In-CSV/internal/init"
	"github.com/chzyer/readline"
	"io"
	"log"
	"strings"
)

func main() {
	filePath := "./config.yaml"
	config, err := config.ReadNewConfig(filePath)
	if err != nil {
		log.Fatal(err)
	}
	mapa, _, _, _ := initial.FilesInit(config)
	lineRdr, err := readline.NewEx(&readline.Config{
		Prompt:            "simpleSQL > ",
		HistoryFile:       "/tmp/searchInCSV.tmp",
		InterruptPrompt:   "^C",
		HistorySearchFold: true,
	})
	if err != nil {
		log.Fatalf("can't run command line: %v", err)
	}
	for {
		str, err := lineRdr.Readline()
		if err != nil {
			if err != readline.ErrInterrupt && err != io.EOF {
				log.Fatalf("read line error: %v", err)
			}
			break
		}
		commands := strings.Split(str, " ")
		switch commands[0] {
		case "show":
			switch commands[1] {
			case "config":
				fmt.Printf("%+v\n", config)
			case "map":
				fmt.Println(mapa)
			}
		case "help":
			fmt.Println("use SELECT field_name FROM file/table WHERE {condition}")
		case "exit":
			log.Fatal("exiting..")
		case "SELECT":
			//RunParser(config)
		default:
			fmt.Println("unknown command")
		}
	}
}
