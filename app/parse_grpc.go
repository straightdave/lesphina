package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/straightdave/lesphina"
)

// gRPC endpoint
type Endpoint struct {
	Name string `json:"name"`
}

func main() {
	if len(os.Args) != 2 {
		exit("Need one argument: pb-go file")
	}

	goFilename := os.Args[1]
	if !strings.HasSuffix(goFilename, ".pb.go") {
		exit("file should has suffix '.pb.go'")
	}

	jsonStr, err := getMeta(goFilename)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println(jsonStr)
}

func getMeta(file string) (string, error) {
	les, err := lesphina.Read(file)
	if err != nil {
		return "", fmt.Errorf("Lesphina failed to read file: %v", err)
	}

	bytes, err := json.MarshalIndent(les.Meta, "", "    ")
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
