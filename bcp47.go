package main

import (
   "flag"
   "fmt"
   "io"
   "os"
   "golang.org/x/text/language"
)

func main() {
   if len(os.Args) != 1 {
      fmt.Fprintf(out, "Missing language code\n")
      os.Exit(1)
   }
   checkBCP47(os.Stdout)
}

func checkBCP47(out io.Writer) int {
   fmt.Fprintf(out, language.Make(os.Args[0]))
   return 0
}