package results

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type TestResults struct {
	TestResults []TestResult `json:"testresults"`
}

type TestResult struct {
	Correct   int       `json:"correct"`
	Attempts  int       `json:"attempts"`
	Train     string    `json:"train"`
	Timestamp time.Time `json:"timestamp"`
}

const ResultJsonPath string = "../data/results.json"

func RecordResult(this_result TestResult, result_filepath string) error {

	jsonFile, err := os.Open(result_filepath)
	if err != nil {
		log.Print(err)
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var allResults TestResults
	json.Unmarshal([]byte(byteValue), &allResults)
	allResults.TestResults = append(allResults.TestResults, this_result)
	jsonData, _ := json.Marshal(allResults)
	ioutil.WriteFile(result_filepath, jsonData, 0777)
	return nil
}
