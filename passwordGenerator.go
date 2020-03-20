package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	numberstr  = "0123456789"
	charstr    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialstr = "~!@#$%^&*()_+"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l password_length")
	flag.StringVar(&charset, "t", "num",
		`-t password generation charset
	num: Only use number [0-9] to generate password
	char: Only use letter [a-zA-Z] to generate password
	mix: Use number and letters to generate password,
	advanced: Use number, letters and specifal character to generate password`)
	flag.Parse()
}

func generatePassword() string {
	var password []byte = make([]byte, 0, length)
	rand.Seed(time.Now().Unix())
	temp := ""
	switch {
	case charset == "num":
		temp = numberstr
	case charset == "char":
		temp = charstr
	case charset == "mix":
		temp = numberstr + charstr
	case charset == "advanced":
		temp = numberstr + charstr + specialstr
	default:
		fmt.Println("Invalid Command Line Arguments")
	}
	for i := 0; i < length; i++ {
		password = append(password, temp[rand.Intn(len(temp))])
	}
	return string(password)
}

func main() {
	parseArgs()
	password := generatePassword()
	fmt.Println(password)
}
