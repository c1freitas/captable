package cmd

import (
	"bufio"
	"c1freitas/captable/data"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	jsonPrefixPadding = ""
	jsonIndentPadding = "  "
)

// ProcessFile is the main function which processes the supplied file.
func ProcessFile(path string, filterDate time.Time) (*data.CapTable, error) {
	fileHandle, err := os.Open(path)
	defer fileHandle.Close()
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(fileHandle)

	capTable := data.CapTable{Owners: make(data.OwnerList), Date: filterDate, DateStr: filterDate.Format(data.DateFormat)}
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if ValidLine(txt) {
			owner, err := ProcessLine(txt)
			if err != nil {
				return nil, err
			}
			err = capTable.AddInvestor(owner)
			if err != nil {
				return nil, err
			}
		}
	}
	capTable.CalculateTotals()
	return &capTable, nil
}

// ValidLine checks to make sure the supplied line is valid for parsing.
// Currently only checks to see if the line starts with a `#` mark,  indicating this is the initial line.
// This can be expanded to do any validation checks before attempting to parse the line.
func ValidLine(txt string) bool {
	if strings.Index(txt, "#") == 0 {
		return false
	}
	return true
}

// ProcessLine take a string and converts the data to a Owner struct. Minor validation checks,
// returns an error if the string can't be parsed.
func ProcessLine(line string) (*data.Owner, error) {
	values := strings.Split(line, ",")
	if len(values) != 4 {
		return nil, fmt.Errorf("ProcessLine: Line did not contain the correct amount of values %d", len(values))
	}
	investedDate, err := time.Parse(data.DateFormat, strings.TrimSpace(values[0]))
	if err != nil {
		return nil, fmt.Errorf("ProcessLine: Could not convert format of INVESTMENT DATE, %v", values[0])
	}
	shares, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf("ProcessLine: Could not convert format of SHARES PURCHASED %v", values[1])
	}
	cashPaid, err := decimal.NewFromString(values[2])
	if err != nil {
		return nil, fmt.Errorf("ProcessLine: Could not convert format of CASH PAID %v", values[2])
	}
	investor := strings.TrimSpace(values[3])
	owner := data.Owner{Shares: shares, CashPaid: cashPaid, Investor: investor, Date: investedDate}
	return &owner, nil
}

// RenderData Marshals the supplied capTable into a json byte array
func RenderData(capTable *data.CapTable) ([]byte, error) {
	if capTable == nil {
		return nil, errors.New("Captable was nil")
	}
	return json.MarshalIndent(capTable, jsonPrefixPadding, jsonIndentPadding)
}
