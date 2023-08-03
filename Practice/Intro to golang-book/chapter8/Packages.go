package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//To search for a smaller string in a bigger string, use the Contains function
	fmt.Println("The Contains() function:")
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("tomato", "oma"), "\n")
	// => true

	//To count the number of times a smaller string occurs in a bigger string, use the Count function:
	fmt.Println("The Cont() function:")
	// func Count(s, sep string) int
	fmt.Println(strings.Count("tomato", "t"), "\n")
	// => 2

	//To determine if a bigger string starts with a smaller string, use the HasPrefix function:
	fmt.Println("The HasPrefix() function:")
	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("tomato", "te"), "\n")
	// => false

	//To determine if a bigger string ends with a smaller string, use the HasSuffix function:
	fmt.Println("The HasSuffix() function:")
	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("tomato", "to"), "\n")
	// => true

	//To find the position of a smaller string in a bigger string, use the Index function (it returns -1 if not found):
	fmt.Println("The Index() function:")
	// func Index(s, sep string) int
	fmt.Println(strings.Index("tomato", "a"), "\n")
	// => 1

	//To take a list of strings and join them together in a single string separated by another string (e.g., a comma), use the Join function:
	fmt.Println("The Join() function:")
	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"Mad", "Tomato"}, "_"), "\n")
	// => "Mad_Tomato"

	//To repeat a string, use the Repeat function:
	fmt.Println("The Repeat() function:")
	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("o", 5), "\n")
	// => "ooooo"

	//To replace a smaller string in a bigger string with some other string, use the Replace
	//function. In Go, Replace also takes a number indicating how many times to do the
	//replacement (pass -1 to do it as many times as possible):
	fmt.Println("The Replace() function:")
	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2), "\n")
	// => "bbaa"

	//To split a string into a list of strings by a separating string (e.g., a comma), use the Split function (Split is the reverse of Join):
	fmt.Println("The Split() function:")
	// func Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-"), "\n")
	// => []string{"a","b","c","d","e"}

	//To convert a string to all lowercase letters, use the ToLower function:
	fmt.Println("The ToLower() function:")
	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TOMATO"), "\n")
	// => "tomato"

	//To convert a string to all uppercase letters, use the ToUpper function:
	fmt.Println("The ToUpper() function:")
	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("tomato"), "\n")
	// => "TOMATO"

	//Sometimes we need to work with strings as binary data. To convert a string to a slice of bytes (and vice versa), do this:
	// arr := []byte("test")
	// str := string([]byte{'t', 'e', 's', 't'})

	//To read or write to a []byte or a string, you can use the Buffer struct found in the bytes package:
	var buf bytes.Buffer
	buf.Write([]byte("test"))
}
