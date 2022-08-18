package main

import (
   "fmt"
   "os"
   "golang.org/x/text/language"
)

func main() {
   if len(os.Args) != 1 {
      fmt.Printf("Missing language code\n")
      os.Exit(1)
   }

   tag := language.Make(os.Args[0])
   fmt.Println(tag)
}