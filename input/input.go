// Package input provides a very simple file reader that just reads the file,
// and returns a slice of strings with the contents. This saves a little bit of code repetition over all days.
// Yes it's overhead, but the input files are mostly very small, and compared to the actual calculation time on some
// puzzles the overhead is trivial.
package input

import (
	"bufio"
	"os"
)

// Read opens the file at the given path, reads it in its entirity, and returns a slice of strings with its contents
func Read(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		// because input is vital to the challenge, panic on any error. We cannot continue anyway
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	m := make([]string, 0)
	for s.Scan() {
		l := s.Text()

		m = append(m, l)
	}

	return m
}
