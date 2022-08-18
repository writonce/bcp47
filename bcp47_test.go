package main
import (

    "bytes"
    "flag"
    "os"
    "testing"
)
func TestCheckBCP47(T *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()
    cases := []struct {
        Name           string
        Args           []string
        ExpectedExit   int
        ExpectedOutput string
    }{
        {"ok1", []string{"fr-fr"}, 0, "fr-FR\n"},
        {"ok2", []string{"fr-FR"}, 0, "fr-FR\n"},
        {"ok3", []string{"FR-FR"}, 0, "fr-FR\n"},
        {"ko1", []string{""}, 1, "Missing language code\n"},
    }
    for _, tc := range cases {
        // this call is required because otherwise flags panics,
        // if args are set between flag.Parse call
        flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
        // we need a value to set Args[0] to cause flag begins parsing at Args[1]
        os.Args = append([]string{tc.Name}, tc.Args...)
        var buf bytes.Buffer
        actualExit := checkBCP47(&buf)
        if tc.ExpectedExit != actualExit {
            T.Errorf("Wrong exit code for args: %v, expected: %v, got: %v",
                tc.Args, tc.ExpectedExit, actualExit)
        }
        actualOutput := buf.String()
        if tc.ExpectedOutput != actualOutput {
            T.Errorf("Wrong output for args: %v, expected %v, got: %v",
                tc.Args, tc.ExpectedOutput, actualOutput)
        }
    }
}