package results

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRecordResult(t *testing.T) {
	tmpfile, err := ioutil.TempFile("", "example.*.json")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	resultData := TestResult{
		Correct:   8,
		Attempts:  10,
		Train:     "spanish",
		Timestamp: time.Now(),
	}
	RecordResult(resultData, tmpfile.Name())

	gotJson, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		log.Fatal(err)
	}
	expectedJson := "{\"testresults\":[{\"correct\":8,\"attempts\":10,\"train\":\"spanish\",\"timestamp\":"
	if !strings.HasPrefix(string(gotJson), expectedJson) {
		t.Errorf("got %q want %q", gotJson, expectedJson+"<SOME TIMESTAMP>")
	}
}
