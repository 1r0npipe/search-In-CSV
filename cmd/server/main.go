package main

import (
	"fmt"
	"github.com/1r0npipe/search-In-CSV/cmd/server/config"
	"github.com/1r0npipe/search-In-CSV/internal/init"
	"github.com/1r0npipe/search-In-CSV/internal/simpleparserSQL"
	"github.com/chzyer/readline"
	"io"
	"log"
	"strings"
	"go/ast"
    "go/parser"
    "go/token"
)

func main() {
	filePath := "./config.yaml"
	config, err := config.ReadNewConfig(filePath)
	if err != nil {
		log.Fatal(err)
	}
	mapa, _, _, _ := initial.FilesInit(config)
	fmt.Println("This is simple CLI query tool, use 'help' command to get more information.")
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
				simpleparserSQL.PrintConfig(config)
			case "map":
				fmt.Println(mapa)
			}
		case "help":
			fmt.Println(simpleparserSQL.HELP)
		case "exit":
			log.Fatal("exiting..")
		case "WHERE":
			condition := strings.Join(commands[1:], " ")
			condition = strings.Replace(condition, "AND", "&&", -1)
			condition = strings.Replace(condition, "OR", "||", -1)
			fmt.Println(condition)
			fs := token.NewFileSet()
			tr, _ := parser.ParseExpr(condition)
			ast.Print(fs, tr)
		default:
			fmt.Println("unknown command")
		}
	}
}
