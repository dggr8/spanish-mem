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
		CommandLineSpy := CliSpy{}
		CommandLineSpy.reader = strings.NewReader("list files\n")
		got := GetCommand(&CommandLineSpy, &CommandLineSpy)
		want := "list files"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		got1 := CommandLineSpy.Calls
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
		CommandLineSpy := CliSpy{}
		CommandLineSpy.reader = strings.NewReader("nonsense\n")
		got := GetCommand(&CommandLineSpy, &CommandLineSpy)
		want := ""
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		got1 := CommandLineSpy.Calls
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
