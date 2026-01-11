package main

import (
	"encoding/base64"
	"os"
)

func decoder() {
	base64Str := "PASTE_BASE64_HERE"

	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		panic("Invalid Base64")
	}

	os.WriteFile("test.jpg", data, 0644)
}
