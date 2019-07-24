package results

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// TestResults is a struct of an array of TestResult.
type TestResults struct {
	TestResults []TestResult `json:"testresults"`
}

// TestResult is a struct to save results of a user session.
// Contains the time, string tag, correct attempts and total attempts.
type TestResult struct {
	Correct   int       `json:"correct"`
	Attempts  int       `json:"attempts"`
	Train     string    `json:"train"`
	Timestamp time.Time `json:"timestamp"`
}

// ResultJSONPath is a hard-coded constant path to the json file to save results on.
const ResultJSONPath string = "../data/results.json"

// RecordResult takes in an instance of TestResult and saves it to the file.
func RecordResult(thisResult TestResult, resultFilePath string) error {

	jsonFile, err := os.Open(resultFilePath)
	if err != nil {
		log.Print(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	allResults := TestResults{}
	json.Unmarshal([]byte(byteValue), &allResults)
	allResults.TestResults = append(allResults.TestResults, thisResult)
	jsonData, _ := json.Marshal(allResults)
	ioutil.WriteFile(resultFilePath, jsonData, 0777)
	return nil
}
