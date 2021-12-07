# Advent of Code 2021

This is my attempt at solving all puzzles in the [Advent of Code 2021](https://adventofcode.com/2021). Everything is written in Golang 1.17.

Since all puzzle inputs are personally generated, I've added the input files for each day as well.

## Packages
During this event I also created a few packages to isolate useful code so they can be used in later solutions as well.

### Input
This package provides a quick way of parsing the puzzle input by reading the entire file in memory, and returning either a slice of strings or slice of ints (if the input is formatted this way, must be checked by the user).

Yes, this package introduces overhead, because for some days it's also possible to calculate the solution _whilst_ reading the input. But I'd rather move all input related code out of the way, so that each solution has a clean way of receiving the input file in a slice.

### Mathutil
This package contains a few mathematical equations that are always handy. Currently supports _mean_, and _median_ on a slice of ints.