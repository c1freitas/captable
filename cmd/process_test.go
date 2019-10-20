package cmd

import (
	"testing"
	"time"
)

func TestSingleCommand(t *testing.T) {

	fileLocation := "../test/example.csv"
	filterDate := time.Now()
	capTable, err := ProcessFile(fileLocation, filterDate)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if capTable == nil {
		t.Error("CapTable Object was nil")
	}

	// if output != "" {
	// 	t.Errorf("Unexpected output: %v", output)
	// }

	// got := strings.Join(, " ")
	// expected := "one two"
	// if got != expected {
	// 	t.Errorf("rootCmdArgs expected: %q, got: %q", expected, got)
	// }
}

// func TestAddInvestorFiltered(t *testing.T) {

// 	addInvestorFiltered()
// }

func TestValidLine(t *testing.T) {
	validCases := []string{
		"2019-01-02,1500,13500.00,Fred Wilson",
		",1000,10000.00,Billy Joel",
		"2#019-01-02,1500,13500.00,Fred Wilson",
	}

	invalidCases := []string{
		"#INVESTMENT DATE, SHARES PURCHASED, CASH PAID, INVESTOR",
	}

	for _, s := range validCases {
		if !ValidLine(s) {
			t.Errorf("%v should be valid", s)
		}
	}

	for _, s := range invalidCases {
		if ValidLine(s) {
			t.Errorf("%v should NOT be valid", s)
		}
	}
}
