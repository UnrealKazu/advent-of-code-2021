package main

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/advent-of-code-2021/input"
)

func main() {
	in := input.Read("input.txt")

	vSum := processTransmission(in[0])

	fmt.Printf("Sum of all version numbers is %d", vSum)
}

func processTransmission(s string) int {
	bits := StupidHexToBitString(s)

	vSum := 0

	i := 0
	for i < len(bits) {
		curI := i
		v, _ := strconv.ParseInt(bits[i:i+3], 2, 32)
		vSum += int(v)

		t, _ := strconv.ParseInt(bits[i+3:i+6], 2, 32)

		i = GetNextStartIndex(s, int(t), i+6)

		diffI := i - curI

		// increment i by the remainder of diff modulo 4, because
		// the base is 16 (hex), we need to take into account the padding of
		// zeroes to reach the base
		i += diffI
	}

	return vSum
}

func GetNextStartIndex(s string, v, i int) int {
	if v == 4 {
		// literal value, so keep on skipping 4 bits until we reach the last group
		cur := string(s[i])

		for cur != "0" {
			// skip ahead 5 bits
			i = i + 5
			cur = string(s[i])
		}

		// final group reached, skip 5 again
		return i + 5
	} else {
		// operator packet
		lT := string(s[i])

		if lT == "0" {
			// next 15 bits is the total length of the next sub-packets
			i++
			subL := s[i : i+15]
			skip, _ := strconv.ParseInt(subL, 2, 32)

			return i + 15 + int(skip)
		} else {
			// next 11 bits are the number of sub-packets
			i++
			subL := s[i : i+11]
			skipNum, _ := strconv.ParseInt(subL, 2, 32)

			return i + 11 + (int(skipNum) * 11)
		}
	}
}

func StupidHexToBitString(s string) string {
	ns := ""

	for _, r := range s {
		i, err := strconv.ParseInt(string(r), 16, 64)

		if err != nil {
			panic(err)
		}

		ns += fmt.Sprintf("%01b", i)
	}

	return ns
}

func HexToBitString(s string) string {
	// first get the length of the byte array
	// this is important for us to later pad the string with
	// any leading zeroes
	b, err := hex.DecodeString(s)

	if err != nil {
		panic(err)
	}

	len := len(b) * 8

	i, err := strconv.ParseUint(s, 16, 64)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%0*b", len, i)
}
