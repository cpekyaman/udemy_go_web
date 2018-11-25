package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("uname:password")))

	b, err := base64.StdEncoding.DecodeString("dW5hbWU6cGFzc3dvcmQ=")
	if err != nil {
		fmt.Println("Decode failed")
	}
	fmt.Println(string(b))
}
