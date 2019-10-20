package data

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

const (
	DateFormat = "2006-01-02"
)

//Owner is the individual owner struct
type Owner struct {
	Shares          int             `json:"shares"`
	Investor        string          `json:"investor"`
	CashPaid        decimal.Decimal `json:"cash_paid"`
	OwnershipAmount decimal.Decimal `json:"ownership"`
	Date            time.Time       `json:"-"`
}

func (o *Owner) add(a *Owner) error {
	if !strings.EqualFold(o.Investor, a.Investor) {
		return errors.New("Investors must match to add")
	}
	o.Shares = o.Shares + a.Shares
	d := o.CashPaid.Add(a.CashPaid)
	log.Printf("\nValue: %v", d.String())
	o.CashPaid = d
	return nil
}

// OwnerList is a custom map type to allow custom unmarshalling, but easy look ups by Investor Name
type OwnerList map[string]Owner

// UnmarshalJSON uses custom code for unmarshalling the types. Maps normally unmarshall
// into an associative array like structure in JSON. We want it to look like an array of objects.
func (ol OwnerList) UnmarshalJSON(b []byte) error {
	return nil
}

// CapTable is the top level data struct
type CapTable struct {
	Date        time.Time       `json:"-"`
	DateStr     string          `json:"date"`
	CashRaised  decimal.Decimal `json:"cash_raised"`
	TotalShares int             `json:"total_number_of_shares"`
	Owners      OwnerList       `json:"ownership"`
}

// CalculateTotals will calculate all the ownership amounts and CapTable values. This is
// intended to be called after loading all values.
// It could be called after adding a new investor each time also, but thats less efficient.
func (c *CapTable) CalculateTotals() {
	// todo: recalculate all values here or as part of Add...

}

// addInvestor adds an investor. If the Investor already exists their values are only adjusted.
// If the investment date is after the capTable date the investor is ignored
func (c *CapTable) AddInvestor(newInvestor *Owner) error {
	if newInvestor.Date.After(c.Date) {
		return nil
	}

	investorName := newInvestor.Investor
	if foundOwner, exists := c.Owners[investorName]; exists {
		err := foundOwner.add(newInvestor)
		if err != nil {
			return err
		}
		c.Owners[investorName] = foundOwner
	} else {
		c.Owners[investorName] = *newInvestor
	}
	return nil
}
