package main

import (
	"encoding/base64"
	"flag"
	"fmt"
)

func main() {
	var (
		cant    string
		command string
		input   string
	)

	flag.StringVar(&cant, "cant", "123123", "cant")
	flag.StringVar(&command, "command", "d", "command")
	flag.StringVar(&input, "input", "", "input")

	flag.Parse()

	if input == "" {
		flag.PrintDefaults()
		return
	}

	switch command {
	case "d":
		output := encoding(cant, input)
		fmt.Println(string(output))

	case "e":
		output := decoding(cant, input)
		fmt.Println(string(output))

	}

}

func encoding(cant, context string) (output []byte) {
	mask := []byte(cant)
	temp := []byte(context)

	cache := make([]byte, len(temp))

	maskLen := len(mask)

	for index, value := range temp {
		cache[index] = byte(int8(value) + int8(mask[index%maskLen]))
	}

	output = make([]byte, base64.StdEncoding.EncodedLen(len(temp)))
	base64.StdEncoding.Encode(output, cache)

	return output
}

func decoding(cant, context string) (output []byte) {
	mask := []byte(cant)
	temp := []byte(context)

	cache := make([]byte, base64.StdEncoding.DecodedLen(len(temp)))
	base64.StdEncoding.Decode(cache, temp)

	maskLen := len(mask)

	output = make([]byte, len(cache))
	for index, value := range cache {
		output[index] = byte(int8(value) - int8(mask[index%maskLen]))
	}

	return output
}
