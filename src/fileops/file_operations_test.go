package fileops

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetWords(t *testing.T) {
	content := []byte(`good, bien
cow, la vaca
buffalo, la bufala
cow, la bufala`)
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
	GetWords(tmpfn)
	EnglishToSpanishExpected := map[string][]string{
		"good":    []string{"bien"},
		"cow":     []string{"la vaca", "la bufala"},
		"buffalo": []string{"la bufala"},
	}
	if !reflect.DeepEqual(EnglishToSpanish, EnglishToSpanishExpected) {
		t.Errorf("got %v want %v", EnglishToSpanish, EnglishToSpanishExpected)
	}
	SpanishToEnglishExpected := map[string][]string{
		"bien":      []string{"good"},
		"la vaca":   []string{"cow"},
		"la bufala": []string{"buffalo", "cow"},
	}
	if !reflect.DeepEqual(SpanishToEnglish, SpanishToEnglishExpected) {
		t.Errorf("got %v want %v", SpanishToEnglish, SpanishToEnglishExpected)
	}
}
