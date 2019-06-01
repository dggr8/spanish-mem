package file_operations

import (
  "encoding/csv"
  "path/filepath"
  "os"
)

func check(e error) {
	if e != nil {
	  panic(e)
	}
}


func LoadFiles(glob_path string) int64 {
	// Get all the data files from the glob path.
	all_files, err := filepath.Glob(glob_path)
	check(err)
	var number_of_words int64
	for _, filename := range all_files {
		file, err := os.Open(filename)
		check(err)
		defer file.Close()

		lines, err := csv.NewReader(file).ReadAll()
		check(err)
		number_of_words += int64(len(lines))
	}

	return number_of_words
}