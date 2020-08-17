package internal

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type TestCases struct {
	Tests []Endpoint `json:"tests"`
}

func ParseTestFile(filepath string) (*TestCases, error) {
	jsonFile, err := os.Open(filepath)
    if err != nil {
        return nil, fmt.Errorf("error parsing file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var testCases TestCases
	json.Unmarshal(byteValue, &testCases)
	if &testCases == nil {
		return nil, fmt.Errorf("error with json in file: %v", err)
	}
	return &testCases, nil
}