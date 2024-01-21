// Copyright (c) 2024 Mattia Cabrini
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

//go:embed res/version.txt
var version string

//go:embed res/welcome-message.txt
var welcomeMessage string

//go:embed res/about.txt
var about string

func runeMapFromString(str string) (runeMap map[rune]int) {
	runeMap = make(map[rune]int)

	for _, rx := range strings.ToLower(str) {
		if !unicode.IsDigit(rx) && !unicode.IsLetter(rx) {
			continue
		}

		n, b := runeMap[rx]

		if !b {
			n = 0
		}

		runeMap[rx] = n + 1
	}

	return
}

func compareRuneMaps(runeMapA, runeMapB map[rune]int) bool {
	if len(runeMapA) != len(runeMapB) {
		return false
	}

	for rx, aNX := range runeMapA {
		bNX, found := runeMapB[rx]

		if !found {
			return false
		}

		if aNX != bNX {
			return false
		}
	}

	return true
}

func readLineP(k *bufio.Scanner) string {
	if !k.Scan() {
		log.Printf("could not read input: %v", k.Err())
		os.Exit(1)
	}

	return k.Text()
}

func checker(k *bufio.Scanner) {
	var line1, line2 string

	fmt.Print("String 1: ")
	line1 = readLineP(k)

	fmt.Print("String 2: ")
	line2 = readLineP(k)

	runeMap1 := runeMapFromString(line1)
	runeMap2 := runeMapFromString(line2)

	ok := compareRuneMaps(runeMap1, runeMap2)

	if ok {
		fmt.Println("Anagram is OK!")
	} else {
		fmt.Println("Anagram is KO!")
	}
}

func showAbout() {
	var found bool = false

	for _, ax := range os.Args {
		ax = strings.ToLower(ax)
		found = (ax == "-l" || ax == "--license")

		if found {
			print(about)
			os.Exit(0)
			break
		}
	}
}

func main() {
	fmt.Printf("check-anagrams v%s\n", version)
	fmt.Printf(welcomeMessage)
	fmt.Printf("\n")

	var sdtinR = bufio.NewReader(os.Stdin)
	var k = bufio.NewScanner(sdtinR)

	showAbout()

	for {
		checker(k)
	}
}
