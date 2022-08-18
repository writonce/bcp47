package main

import (
   "fmt"
   "os"
   "io"
   "golang.org/x/text/language"
)

func main() {
   os.Exit(checkBCP47(os.Stdout))
}

func checkBCP47(out io.Writer) int {
   if len(os.Args) != 1 {
      fmt.Fprintf(out,"Missing language code\n")
      return 1
   }

   tag := language.Make(os.Args[0])
   fmt.Fprintln(out,tag)
   return 0
}