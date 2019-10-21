package cmd

import (
	"testing"
	"time"

	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
)

func TestProcessFile(t *testing.T) {

	fileLocation := "../test/example.csv"
	filterDate := time.Now()
	capTable, err := ProcessFile(fileLocation, filterDate)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if capTable == nil {
		t.Error("CapTable Object was nil")
	}

	assert.Equal(t, capTable.CashRaised.StringFixed(2), "165500.50")
	assert.Equal(t, capTable.TotalShares, 9500)
	assert.Check(t, cmp.Len(capTable.Owners, 4))
	assert.Check(t, cmp.Equal(capTable.Owners["Sandy Lerner"].OwnershipAmount.StringFixed(2), "31.58"))
}

func TestProcessLine(t *testing.T) {
	validLine := "2018-01-20,2000,40000.00, Don Valentine "
	inValidDate := "20-0-20,2000,40000.00,Don Valentine"
	inValidLength := "2018-01-20,2000,Don Valentine"
	inValidShares := "2018-01-20,abc,40000.00, Don Valentine "
	inValidCash := "2018-01-20,2000,abc, Don Valentine "

	owner, err := ProcessLine(validLine)
	assert.NilError(t, err)
	assert.Equal(t, owner.Shares, 2000)
	assert.Equal(t, owner.Investor, "Don Valentine")

	_, err = ProcessLine(inValidDate)
	assert.ErrorContains(t, err, "INVESTMENT DATE")
	_, err = ProcessLine(inValidLength)
	assert.ErrorContains(t, err, "correct amount of values")
	_, err = ProcessLine(inValidShares)
	assert.ErrorContains(t, err, "SHARES PURCHASED")
	_, err = ProcessLine(inValidCash)
	assert.ErrorContains(t, err, "CASH PAID")

}

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
