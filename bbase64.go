package main

import (
	"fmt"
	"os"
	b64 "encoding/base64"
	"strings"
  )

func main() {
    
	input := os.Args[2]
	op := os.Args[1]
	fmt.Println(input)
	if strings.Compare(op, "e") == 0 {
	sEnc := b64.StdEncoding.EncodeToString([]byte(input))
    fmt.Println(sEnc)
	} else if strings.Compare(op, "d") == 0 {
	sDec, _ := b64.StdEncoding.DecodeString(input)
    fmt.Println(string(sDec))
	} else{
		fmt.Println("Base64 encrypt/decrypt")
		fmt.Println("usage :bbase64 e/d input")
	}
}