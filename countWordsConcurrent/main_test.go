package main

import (
 "fmt"
 "testing"
)

func Test(t *testing.T) {
 type testCase struct {
  input    []string
  expected int
 }

 runCases := []testCase{
  {
   []string{"need heals at the front gate", "incoming boss fight get ready"},
   11,
  },
  {
   []string{"", "single", "two words"},
   3,
  },
  {
   []string{"multiple   spaces between   words", "new\nline separated words"},
   8,
  },
 }

 submitCases := append(runCases, []testCase{
  {
   []string{},
   0,
  },
  {
   []string{"solo"},
   1,
  },
  {
   []string{"three whole words", "and four more here"},
   7,
  },
 }...)

 testCases := runCases
 if withSubmit {
  testCases = submitCases
 }

 skipped := len(submitCases) - len(testCases)

 passCount := 0
 failCount := 0

 for _, test := range testCases {
  result := countWordsConcurrent(test.input)
  if result != test.expected {
   failCount++
   t.Errorf(`---------------------------------
Input texts: %v

Expected total words: %d
Actual total words:   %d
Fail
`, test.input, test.expected, result)
  } else {
   passCount++
   fmt.Printf(`---------------------------------
Input texts: %v

Expected total words: %d
Actual total words:   %d
Pass
`, test.input, test.expected, result)
  }
 }

 fmt.Println("---------------------------------")
 if skipped > 0 {
  fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
 } else {
  fmt.Printf("%d passed, %d failed\n", passCount, failCount)
 }
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
