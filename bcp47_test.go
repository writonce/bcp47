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
        {"ok1", "fr-fr", 0, "fr-FR"},
        {"ok2", "fr-FR", 0, "fr-FR"},
        {"ok3", "FR-FR", 0, "fr-FR"},
        {"ko1", "", 1, "Missing language code\n"},
    }
    for _, tc := range cases {
        flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
        os.Args = []string{tc.Args}
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