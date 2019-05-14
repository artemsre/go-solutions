package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

var m = map[string]int{"one": 1, "two": 2, "three": 3}

func main() {
	encodeFile, err := os.Create("mapSerialize.gob")
	if err != nil {
		panic(err)
	}
	e := gob.NewEncoder(encodeFile)

	// Encoding the map
	err = e.Encode(m)
	if err != nil {
		panic(err)
	}
	encodeFile.Close()

	decodeFile, err := os.Open("mapSerialize.gob")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()
	var decodedMap map[string]int
	d := gob.NewDecoder(decodeFile)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	// Ta da! It is a map!
	fmt.Printf("%#v\n", decodedMap)
}
