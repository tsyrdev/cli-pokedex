package repl

import (
    "testing"
    "fmt"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input       string
        expected    []string
    }{
        {
        input:      "   hello      world ",
        expected:   []string{"hello", "world"},
        },
        {
            input:      "my name is @manel",
            expected:   []string{"my", "name", "is", "@manel"},
        },
    }

    passCount := 0
    failCount := 0 

    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            failCount++
            t.Errorf(`------------------
Test Failed for length of returned input slice
Expecting:  %v 
Actual:     %v
Fail
`, len(c.expected), len(actual))
            continue 
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]

            if word != expectedWord {
                failCount++
                t.Errorf(`------------------
Test Failed for equal words: %v 
Expecting:  %v 
Actual:     %v
Fail 
`, c.input, expectedWord, word)
            } else {
                passCount++ 
                fmt.Printf(`----------------
Test Passed for equal words: %v 
Expecting:  %v
Actual:     %v
Pass 
`, c.input, expectedWord, word)
            }
        }

        fmt.Println("--------------")
        fmt.Printf("%d passed, %d failed\n", passCount, failCount)
    }
}
