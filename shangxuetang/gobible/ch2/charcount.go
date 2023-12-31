package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Charcount computes counts of Unicode characters.

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax+1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)

	var (
		letters, numbers, puncts, symbles int
	)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letters++
		}
		if unicode.IsNumber(r) {
			numbers++
		}
		if unicode.IsPunct(r) {
			puncts++
		}
		if unicode.IsSymbol(r) {
			symbles++
		}
		counts[r]++
		utflen[n]++
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Printf("len\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n")
		
	}
	fmt.Printf("letters: %d\n", letters)
	fmt.Printf("numbers: %d\n", numbers)
	fmt.Printf("puncts: %d\n", puncts)
	fmt.Printf("symbles: %d\n", symbles)
}