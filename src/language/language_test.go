package language

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestTestSpanish(t *testing.T) {
	spy := CliSpy{}
	spy.reader = strings.NewReader("0\n")
	TestSpanish(&spy, &spy, "")
	expectedCalls := []string{
		write,
		read,
		read,
		write,
	}
	if !reflect.DeepEqual(spy.Calls, expectedCalls) {
		t.Errorf("got calls %v want %v", spy.Calls, expectedCalls)
	}

	expectedPrints := "How many words do you want to train->Translate these words to spanish:\n"
	if spy.Prints != expectedPrints {
		t.Errorf("printed %q but wanted to print %q", spy.Prints, expectedPrints)
	}
}

func TestTestEnglish(t *testing.T) {
	spy := CliSpy{}
	spy.reader = strings.NewReader("0\n")
	TestEnglish(&spy, &spy, "")
	expectedCalls := []string{
		write,
		read,
		read,
		write,
	}
	if !reflect.DeepEqual(spy.Calls, expectedCalls) {
		t.Errorf("got calls %v want %v", spy.Calls, expectedCalls)
	}

	expectedPrints := "How many words do you want to train->Translate these words to english:\n"
	if spy.Prints != expectedPrints {
		t.Errorf("printed %q but wanted to print %q", spy.Prints, expectedPrints)
	}
}

func TestTestLanguage(t *testing.T) {

	t.Run("trainCount is zero", func(t *testing.T) {
		spy := CliSpy{}
		spy.reader = strings.NewReader("")
		TestLanguage(&spy, &spy, nil, 0, "", "")

		if len(spy.Calls) != 0 {
			t.Errorf("Wanted no calls but got %v", spy.Calls)
		}
	})

	t.Run("correct answers", func(t *testing.T) {
		tmpfile, err := ioutil.TempFile("", "example.*.json")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())
		spy := CliSpy{}
		spy.reader = strings.NewReader("home\n")

		languageMap := map[string][]string{
			"la casa": []string{"house", "home"},
		}

		TestLanguage(&spy, &spy, languageMap, 1, "english", tmpfile.Name())

		gotJSON, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
		expectedJSON := "{\"testresults\":[{\"correct\":1,\"attempts\":1,\"train\":\"english\",\"timestamp\":"
		if !strings.HasPrefix(string(gotJSON), expectedJSON) {
			t.Errorf("got %q want %q", gotJSON, expectedJSON+"<SOME TIMESTAMP>")
		}

		expectedCalls := []string{
			write,
			read,
			write,
			write,
		}
		if !reflect.DeepEqual(spy.Calls, expectedCalls) {
			t.Errorf("got calls %v want %v", spy.Calls, expectedCalls)
		}

		expectedPrints := "la casa->Correct!\n1 correct out of 1\n"
		if spy.Prints != expectedPrints {
			t.Errorf("printed %q but wanted to print %q", spy.Prints, expectedPrints)
		}
	})

	t.Run("skipping", func(t *testing.T) {
		tmpfile, err := ioutil.TempFile("", "example.*.json")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())
		spy := CliSpy{}
		spy.reader = strings.NewReader("\nhome\n")

		languageMap := map[string][]string{
			"aaaaaa": []string{"aaaaaa"},
		}

		TestLanguage(&spy, &spy, languageMap, 1, "english", tmpfile.Name())

		gotJSON, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
		expectedJSON := "{\"testresults\":[{\"correct\":0,\"attempts\":1,\"train\":\"english\",\"timestamp\":"
		if !strings.HasPrefix(string(gotJSON), expectedJSON) {
			t.Errorf("got %q want %q", gotJSON, expectedJSON+"<SOME TIMESTAMP>")
		}

		expectedCalls := []string{
			write,
			read,
			write,
			write,
		}
		if !reflect.DeepEqual(spy.Calls, expectedCalls) {
			t.Errorf("got calls %v want %v", spy.Calls, expectedCalls)
		}

		expectedPrints := "aaaaaa->It is one of \"[aaaaaa]\".\n0 correct out of 1\n"
		if spy.Prints != expectedPrints {
			t.Errorf("printed %q but wanted to print %q", spy.Prints, expectedPrints)
		}
	})

	t.Run("wrong answer", func(t *testing.T) {
		tmpfile, err := ioutil.TempFile("", "example.*.json")
		if err != nil {
			log.Fatal(err)
		}
		defer os.Remove(tmpfile.Name())
		spy := CliSpy{}
		spy.reader = strings.NewReader("home\n")

		languageMap := map[string][]string{
			"aaaaaa": []string{"aaaaaa"},
		}

		TestLanguage(&spy, &spy, languageMap, 1, "english", tmpfile.Name())

		gotJSON, err := ioutil.ReadAll(tmpfile)
		if err != nil {
			log.Fatal(err)
		}
		expectedJSON := "{\"testresults\":[{\"correct\":0,\"attempts\":1,\"train\":\"english\",\"timestamp\":"
		if !strings.HasPrefix(string(gotJSON), expectedJSON) {
			t.Errorf("got %q want %q", gotJSON, expectedJSON+"<SOME TIMESTAMP>")
		}

		expectedCalls := []string{
			write,
			read,
			write,
			write,
		}
		if !reflect.DeepEqual(spy.Calls, expectedCalls) {
			t.Errorf("got calls %v want %v", spy.Calls, expectedCalls)
		}

		expectedPrints := "aaaaaa->Nah! It is one of \"[aaaaaa]\".\n0 correct out of 1\n"
		if spy.Prints != expectedPrints {
			t.Errorf("printed %q but wanted to print %q", spy.Prints, expectedPrints)
		}
	})
}

func TestMinDistance(t *testing.T) {

	t.Run("Max distance", func(t *testing.T) {
		want := 9
		got := MinDistance("something", []string{"bbbbbbbbbbbbbb"})

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("perfect match", func(t *testing.T) {
		want := 0
		got := MinDistance("something", []string{"bbbbbbbbbbbbbb", "something"})

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
