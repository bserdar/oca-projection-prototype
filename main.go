package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	//"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
)

type ProjectionOverlay struct {
	Project []ProjectionRule `json:"project"`
}

type ProjectionRule struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Array  bool   `json:"array"`
}

func (p ProjectionRule) process(input interface{}, output map[string]interface{}) error {
	data, err := jsonpath.Get(p.Source, input)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil || data == nil {
		return nil
	}
	if _, ok := output[p.Target]; ok {
		return fmt.Errorf("Key %s already populated", p.Target)
	}
	if arr, ok := data.([]interface{}); ok {
		if !p.Array {
			if len(arr) == 1 {
				data = arr[0]
			} else if len(arr) > 0 {
				return fmt.Errorf("Multiple values for %s", p.Source)
			} else {
				data = nil
			}
		}
	}
	if data != nil {
		output[p.Target] = data
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Args: inputFile projectionOverlay")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	var input map[string]interface{}
	if err := json.Unmarshal(data, &input); err != nil {
		panic(err)
	}

	data, err = ioutil.ReadFile(os.Args[2])
	if err != nil {
		panic(err)
	}
	var projection ProjectionOverlay
	if err := json.Unmarshal(data, &projection); err != nil {
		panic(err)
	}
	output := make(map[string]interface{})
	for _, x := range projection.Project {
		if err := x.process(input, output); err != nil {
			panic(err)
		}
	}
	out, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(out))
}
