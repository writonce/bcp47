package main

import (
   "fmt"
   "golang.org/x/text/language"
)

func main() {
   tag := language.Make("en_us")
   fmt.Println(tag)
}