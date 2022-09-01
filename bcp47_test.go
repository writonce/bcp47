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
        {"ko1", []string{""}, 0, "\n"},
        {"ko2", []string{"dfsdgdsgdf"}, 0, "\n"},
    }
    for _, tc := range cases {
        flag.CommandLine = flag.NewFlagSet(tc.Name, flag.ExitOnError)
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