package internal

import (
	"testing"
)

func TestFileParsing(t *testing.T) {
	testFilepath := "../test_files/sample.json"
	testCases, err := ParseTestFile(testFilepath)
	if err != nil {
		t.Errorf("error: %s could not be read", testFilepath)
	}
	expectedNumTests := 1
	actualNumTests := len(testCases.Tests)
	if expectedNumTests != actualNumTests {
		t.Errorf("error: expected %d got %d test cases when parsing file %s", expectedNumTests, actualNumTests, testFilepath)
	}
	idx0 := testCases.Tests[0]
	if idx0.Uri != "1.1.1.1" {
		t.Errorf("error: expected endpoint to be %s, got %s", idx0.Uri, "1.1.1.1")
	}
	idx0Port := idx0.ConnectionParams[0].Port
	if idx0Port != 53 {
		t.Errorf("error: expected 53, got %d", idx0Port)
	}
	idx0Protocol := idx0.ConnectionParams[0].Protocol
	if idx0Protocol != "tcp" {
		t.Errorf("error: expected tcp, got %s", idx0Protocol)
	}
}

func TestErrorFileJsonFormat(t *testing.T) {
	testFilepath := "../test_files/bad-format.json"
	_, err := ParseTestFile(testFilepath)
	if err != nil {
		t.Errorf("error: bad json formatting %x", testFilepath)
	}
}