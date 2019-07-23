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

const ResultFilePath string = "../data/results.json"

func RecordResult(this_result TestResult, result_filepath string) {

	jsonFile, err := os.Open(result_filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var all_results TestResults
	json.Unmarshal([]byte(byteValue), &all_results)
	all_results.TestResults = append(all_results.TestResults, this_result)
	json_data, _ := json.Marshal(all_results)
	ioutil.WriteFile(result_filepath, json_data, 0777)
}
