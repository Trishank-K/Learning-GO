package main

import "fmt"

/*
Compiled
- Go tool can run file directly, without VM.
(Due to the fact that the runtime is baked into the final executable file).
- Executable are different for OS

Where is it used?
System apps to web apps - Cloud

Object Oriented
- Yes and No
- What you see on the screen is the code

Missing
- No try catch
(lexer does a lot of work)

*/
func main() {
	fmt.Println("Hello World")
}

/*
	Lexer's job is to check that if we're following the grammar of the language and the syntax is correct, variables are properly defined, and any other pre-computation stuff.
*/

/*
	Types
	- Case insensitive; almost
	- Variable type should be known in advance
	- Everything is a Type

	String
	Bool
	Integer - uint8, uint64, int8, int64, uintptr
				[aliases are involved too]
	Floating - float32, float64
	Complex

	Advance Types
	Arrays
	Slices
	Maps
	Structs
	Pointers
	Functions, Channels, Mutexes
*/
