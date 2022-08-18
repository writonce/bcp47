package main

import (
   "fmt"
   "golang.org/x/text/language"
)

func main() {
   if len(os.Args) != 1 {
      fmt.printf("Missing language code\n")
      os.Exit(1)
   }

   tag := language.Make(os.Args[0])
   fmt.Println(tag)
}