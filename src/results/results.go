package results

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
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

// PrintableResult stores the values of a row in the table printed.
type PrintableResult struct {
	Train     string
	Accuracy  float64
	TestCount int
}

// PrintResults displays the historical results for the user.
func PrintResults(wr io.Writer, resultFilePath string) error {

	jsonFile, err := os.Open(resultFilePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	allResults := TestResults{}
	json.Unmarshal([]byte(byteValue), &allResults)

	var englishAccurate, spanishAccurate, englishTotal, spanishTotal int64
	for _, result := range allResults.TestResults {
		if result.Train == "english" {
			englishAccurate += int64(result.Correct)
			englishTotal += int64(result.Attempts)
		}
		if result.Train == "spanish" {
			spanishAccurate += int64(result.Correct)
			spanishTotal += int64(result.Attempts)
		}
	}

	table := [][]string{
		[]string{"spanish", strconv.FormatInt(spanishTotal, 10),
			strconv.FormatFloat(float64(spanishAccurate)/(float64(spanishTotal)+0.001), 'g', -1, 32)},
		[]string{"english", strconv.FormatInt(englishTotal, 10),
			strconv.FormatFloat(float64(englishAccurate)/(float64(englishTotal)+0.001), 'g', -1, 32)},
	}
	header := []string{
		"Language", "Total tested", "Accuracy",
	}
	printTable(wr, header, table)
	return nil
}

func printTable(wr io.Writer, header []string, table [][]string) {
	newTableWriter := tablewriter.NewWriter(wr)
	newTableWriter.SetHeader(header)
	newTableWriter.AppendBulk(table)
	newTableWriter.Render()
}
