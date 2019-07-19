package cli

import (
	"reflect"
	"strings"
	"testing"
)

const write = "write"
const read = "read"

type CliSpy struct {
	reader *strings.Reader
	Calls  []string
	Prints string
}

func (s *CliSpy) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	s.Prints = s.Prints + string(b)
	return
}

func (s *CliSpy) Read(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, read)
	return s.reader.Read(b)
}

func TestGetCommand(t *testing.T) {

	t.Run("Command is in the list", func(t *testing.T) {
		Spy := CliSpy{}
		Spy.reader = strings.NewReader("list files\n")
		got := GetCommand(&Spy, &Spy)
		want := "list files"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		got1 := Spy.Calls
		want1 := []string{
			write,
			write,
			read,
		}
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("wanted calls %v got %v", want1, got1)
		}
	})

	t.Run("Command not in the list", func(t *testing.T) {
		Spy := CliSpy{}
		Spy.reader = strings.NewReader("nonsense\n")
		got := GetCommand(&Spy, &Spy)
		want := ""
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		got1 := Spy.Calls
		want1 := []string{
			write,
			write,
			read,
			write,
		}
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("wanted calls %v got %v", want1, got1)
		}
	})
}

func TestGetAnswer(t *testing.T) {
	Spy := CliSpy{}
	Spy.reader = strings.NewReader("my answer\n")
	got := GetAnswer(&Spy, &Spy, "Gimme an answer")

	want := "my answer"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	gotCalls := Spy.Calls
	wantCalls := []string{
		write,
		read,
	}
	if !reflect.DeepEqual(gotCalls, wantCalls) {
		t.Errorf("got calls %v want %v", gotCalls, wantCalls)
	}
}

func TestGetInt(t *testing.T) {

	t.Run("input is not a number", func(t *testing.T) {
		Spy := CliSpy{}
		Spy.reader = strings.NewReader("bs\n")
		got, err := GetInt(&Spy, &Spy, "Gimme a number")

		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

		errExpected := "expected integer"
		if err.Error() != errExpected {
			t.Errorf("got error %q wanted %q", err.Error(), errExpected)
		}

		gotCalls := Spy.Calls
		wantCalls := []string{
			write,
			read,
		}
		if !reflect.DeepEqual(gotCalls, wantCalls) {
			t.Errorf("wanted calls %v got %v", wantCalls, gotCalls)
		}

		gotPrint := Spy.Prints
		expectedPrint := "Gimme a number->"
		if gotPrint != expectedPrint {
			t.Errorf("got %q wanted %q", gotPrint, expectedPrint)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		Spy := CliSpy{}
		Spy.reader = strings.NewReader("23\n")
		got, err := GetInt(&Spy, &Spy, "Gimme a number")
		want := 23
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
		if err != nil {
			t.Errorf("expected no error but got %q", err.Error())
		}
	})
}
