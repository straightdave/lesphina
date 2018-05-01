package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/straightdave/lesphina"
	"github.com/straightdave/lesphina/item"
)

// gRPC
type Service struct {
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	Name    string `json:"name"`
	Intype  string `json:"in"`
	Outtype string `json:"out"`
}

func main() {
	if len(os.Args) != 2 {
		exit("Need one argument: pb-go file")
	}

	goFilename := os.Args[1]
	if !strings.HasSuffix(goFilename, ".pb.go") {
		exit("file should has suffix '.pb.go'")
	}

	info, err := getServiceInfo(goFilename)
	if err != nil {
		exit(err.Error())
	}

	fmt.Println(info)
}

func getServiceInfo(file string) (string, error) {
	les, err := lesphina.Read(file)
	if err != nil {
		return "", fmt.Errorf("Lesphina failed to read file: %v", err)
	}

	var grpc []Service
	for _, inf := range les.Meta.Interfaces {
		if !strings.HasSuffix(inf.Name, "Server") {
			// only consider server interface for its conciseness
			continue
		}

		svc := Service{
			Name: inf.Name,
		}

		var ends []Endpoint
		for _, met := range inf.Methods {
			ends = append(ends, Endpoint{
				Name:    met.Name,
				Intype:  item.FirstInParam(met, "~Request").RawType,
				Outtype: item.FirstOutParam(met, "~Response").RawType,
			})
		}

		svc.Endpoints = ends
		grpc = append(grpc, svc)
	}

	bytes, _ := json.MarshalIndent(grpc, "", "    ")

	return string(bytes), nil
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(0)
}
