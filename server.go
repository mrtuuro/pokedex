package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Server struct {
    Config *Config
}

func NewServer(cfg *Config) *Server {
	return &Server{
        Config: cfg,
    }
}

func (s *Server) Start() {
    commandMap := getCommandMap()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("pokedex >>> ")
		scanned := scanner.Scan()
		if !scanned {
			continue
		}
		scannedText := scanner.Text()

		time.Sleep(200 * time.Millisecond)

		scannedText = strings.ToLower(scannedText)
		fields := strings.Fields(scannedText)
		if len(fields) == 0 {
			continue
		}

		command, ok := commandMap[fields[0]]
		if !ok {
			continue
		}

        err := command.ExecuteFunc(s.Config, fields)
        if err != nil {
            log.Fatal(err)
        }
        continue

	}
}

func (s *Server) Stop() {

}
