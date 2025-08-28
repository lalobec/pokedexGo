package main

import (
  "testing"
)

func TestCleanInput(t *testing.T) {
  // Create a slice of test case structs
  cases := []struct {
    input string
    expected []string
  }{
    {
      input: "   heLlo   world   ",
      expected: []string{"hello", "world"},
    },
    {
      input: "Charmander   Bulbasaur PIKaChu",
      expected: []string{"charmander", "bulbasaur", "pikachu"},
    },
    {
      input: "    ",
      expected: []string{},
    },
    {
      input: "  hellOo o  ",
      expected: []string{"helloo", "o"},
    },
    // We can add more cases here
  }

  // Then we loop over all cases and run the tests:
  for _, c := range cases {
    
    actual := cleanInput(c.input)

    if len(actual) != len(c.expected) {
      t.Errorf("Test failed ------ \n Expected: %v \n Actual: %v", c.expected, actual)
    }

    for i := range actual {
      word := actual[i]
      expectedWord := c.expected[i]
      if word != expectedWord {
        t.Errorf("Test failed ----- \n Expected: %v \n Actual: %v", expectedWord, word)
      }
    }
  }
}


// ----------------------------------------//
// Tests for the commandCatch function //
// ----------------------------------------//



