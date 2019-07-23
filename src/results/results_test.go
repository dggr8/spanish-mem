package results

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestRecordResult(t *testing.T) {

	t.Run("File exists", func(t *testing.T) {
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

		gotJSON, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
		expectedJSON := "{\"testresults\":[{\"correct\":8,\"attempts\":10,\"train\":\"spanish\",\"timestamp\":"
		if !strings.HasPrefix(string(gotJSON), expectedJSON) {
			t.Errorf("got %q want %q", gotJSON, expectedJSON+"<SOME TIMESTAMP>")
		}
	})

	t.Run("File doesn't exist", func(t *testing.T) {
		dir, err := ioutil.TempDir("", "example")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(dir)

		fakeFile := dir + "fake.txt"
		resultData := TestResult{
			Correct:   8,
			Attempts:  10,
			Train:     "spanish",
			Timestamp: time.Now(),
		}
		previousFileInfo, _ := os.Stat(fakeFile)
		gotErr := RecordResult(resultData, fakeFile)
		afterFileInfo, _ := os.Stat(fakeFile)

		if !reflect.DeepEqual(previousFileInfo, afterFileInfo) {
			t.Errorf("before %v after %v", previousFileInfo, afterFileInfo)
		}
		if gotErr == nil {
			t.Error("Wanted error but got none.")
		}
		expectedSuffix := "no such file or directory"
		if !strings.HasSuffix(gotErr.Error(), expectedSuffix) {
			t.Errorf("got %q want <Something>: %q", gotErr.Error(), expectedSuffix)
		}
	})
}
