// Package input provides a very simple file reader that just reads the file,
// and returns a slice of strings with the contents. This saves a little bit of code repetition over all days.
// Yes it's overhead, but the input files are mostly very small, and compared to the actual calculation time on some
// puzzles the overhead is trivial.
package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Read opens the file at the given path, reads it in its entirety, and returns a slice of strings with its contents
func Read(path string) []string {
	f, err := os.Open(path)

	if err != nil {
		// because input is vital to the challenge, panic on any error. We cannot continue anyway
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	buf := make([]byte, 0, 64*1024)
	s.Buffer(buf, 1024*1024*10)

	m := make([]string, 0)
	for s.Scan() {
		l := s.Text()

		m = append(m, l)
	}

	if err := s.Err(); err != nil {
		fmt.Printf("Error encountered during scanning. Cannot continue: %s\n", err)
		panic("Error during file read. Exiting.")
	}

	return m
}

// ReadInt opens the file at the given path, reads it in its entirety, and returns a slice of ints with its contents
func ReadInt(path string) []int {
	f, err := os.Open(path)

	if err != nil {
		// because input is vital to the challenge, panic on any error. We cannot continue anyway
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	m := make([]int, 0)
	for s.Scan() {
		l, _ := strconv.Atoi(s.Text())

		m = append(m, l)
	}

	return m
}
