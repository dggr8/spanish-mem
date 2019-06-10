package results

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type TestResults struct {
	TestResults []TestResult `json:"testresults"`
}

type TestResult struct {
	Correct  int    `json:"correct"`
	Attempts int    `json:"attempts"`
	Train    string `json:"train"`
	Timestamp time.Time `json:"timestamp"`
}

const result_filepath string = "../data/results.json"

func RecordResult(correct int, attempts int, train string) {
	this_result := TestResult{
		Correct: correct,
		Attempts: attempts,
		Train: train,
		Timestamp: time.Now(),
	}

	jsonFile, err := os.Open(result_filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var all_results TestResults
	json.Unmarshal([]byte(byteValue), &all_results)
	all_results.TestResults = append(all_results.TestResults, this_result)
	json_data, _ := json.Marshal(all_results)
	ioutil.WriteFile(result_filepath, json_data, 0777)
}