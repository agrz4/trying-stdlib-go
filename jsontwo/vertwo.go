package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// --- Using json.Marshal and json.Unmarshal ---
	fmt.Println("--- Using json.Marshal and json.Unmarshal ---")

	// Marshal John to JSON bytes
	john := Person{Name: "John", Age: 23}
	b, err := json.Marshal(john)
	if err != nil {
		fmt.Println("Error marshalling John:", err)
		return
	}
	fmt.Println("Marshalled John:", string(b))

	// Unmarshal JSON bytes back into a Person struct
	var johnUnmarshal Person
	err = json.Unmarshal(b, &johnUnmarshal)
	if err != nil {
		fmt.Println("Error unmarshalling John:", err)
		return
	}
	fmt.Println("Unmarshalled John:", johnUnmarshal)

	// --- Using json.NewEncoder and json.NewDecoder (for streaming) ---
	fmt.Println("\n--- Using json.NewEncoder and json.NewDecoder ---")

	// Marshal Alice using json.NewEncoder
	alice := Person{Name: "Alice", Age: 25}
	out := new(strings.Builder) // io.Writer to write JSON to
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ") // Makes the output human-readable (optional)

	err = enc.Encode(alice)
	if err != nil {
		fmt.Println("Error encoding Alice:", err)
		return
	}
	fmt.Println("Encoded Alice:\n", out.String())

	// Unmarshal Bob using json.NewDecoder
	in := strings.NewReader(`{"Name":"Bob","Age":30}`) // io.Reader to read JSON from
	dec := json.NewDecoder(in)
	var bob Person
	err = dec.Decode(&bob)
	if err != nil {
		fmt.Println("Error decoding Bob:", err)
		return
	}
	fmt.Println("Decoded Bob:", bob)
}
