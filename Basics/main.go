package main

import "fmt"

func main() {
	fmt.Println(("Hello World"));
}

// lexer's job is to break the input into tokens ( like manually putting semicolons in the input ) 
// parser's job is to build the AST
// compiler's job is to build the bytecode

// Basic Types

// 1. Strings
// 2. Integers ( uint8, uint64, int8, int64, uintptr )
// 3. Floats ( float32, float64 )
// 4. Booleans
// 5. Complex

// Advanced Types

// 1. Arrays
// 2. Slices
// 3. Maps
// 4. Structs
// 5. Pointers 