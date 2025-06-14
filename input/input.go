package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Do not use inToken and inLine together
var (
	inToken *bufio.Scanner
	inLine  *bufio.Scanner
)

// init() initialize var in package - behave as constructor
func init() {
	inToken = bufio.NewScanner(os.Stdin)
	inLine = bufio.NewScanner(os.Stdin)
	inToken.Split(bufio.ScanWords)
	inLine.Split(bufio.ScanLines)
}

// NextInt return single token of integer
func NextInt() int {
	if !inToken.Scan() {
		panic("NextInt: failed to scan token (eof or input error)")
	}
	token := inToken.Text()
	num, err := strconv.Atoi(token)
	if err != nil {
		panic("NextInt: invalid integer token (" + token + ")")
	}
	return num
}

// NextStr return single token of string
func NextStr() string {
	if !inToken.Scan() {
		panic("NextStr: failed to scan token (eof or input error)")
	}
	return inToken.Text()
}

// NextIntLine reads full line and return integers parsed from it
func NextIntLine() []int {
	if !inLine.Scan() {
		panic("NextIntLine: failed to scan line (eof or input error)")
	}
	line := strings.Fields(inLine.Text())
	ret := make([]int, len(line))
	for ind, val := range line {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic("NextIntLine: invalid integer token (" + val + ")")
		}
		ret[ind] = num
	}
	return ret
}

// NextStrLine reads full line ans return strings parsed from it
func NextStrLine() []string {
	if !inLine.Scan() {
		panic("NextStrLine: failed to scan token (eof or input error)")
	}
	return strings.Fields(inLine.Text())
}

// InitForTesting initialize inToken/inLine to read input as string
func InitForTesting(s string) {
	inToken = bufio.NewScanner(strings.NewReader(s))
	inLine = bufio.NewScanner(strings.NewReader(s))
	inToken.Split(bufio.ScanWords)
	inLine.Split(bufio.ScanLines)
}
