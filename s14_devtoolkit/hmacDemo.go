package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	fmt.Println(generateHmac("cenk"))
	fmt.Println(generateHmac("pekyaman"))
}

func generateHmac(s string) string {
	h := hmac.New(sha256.New, []byte("mykey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
