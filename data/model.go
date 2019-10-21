package data

import (
	"encoding/json"
	"errors"
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
	o.CashPaid = o.CashPaid.Add(a.CashPaid)
	return nil
}

// OwnerList is a custom map type to allow custom unmarshalling, but easy look ups by Investor Name
type OwnerList map[string]Owner

// MarshalJSON uses custom code for unmarshalling the types. Maps normally unmarshall
// into an associative array like structure in JSON. We want it to look like an array of objects.
func (ol OwnerList) MarshalJSON() ([]byte, error) {
	// if you want to optimize you can use a bytes.Buffer and write the strings out yourself.
	ownerArray := make([]Owner, len(ol))
	i := 0
	for _, v := range ol {
		ownerArray[i] = v
		i++
	}
	return json.Marshal(ownerArray)
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
// NOTE: ownership amounts are rounded and may not add up to 100%.
func (c *CapTable) CalculateTotals() {
	// Generate the shares and cash raised
	for _, owner := range c.Owners {
		c.CashRaised = c.CashRaised.Add(owner.CashPaid)
		c.TotalShares = c.TotalShares + owner.Shares
	}

	// assign ownership amounts
	totalSharesDecimal := decimal.NewFromFloat(float64(c.TotalShares))
	for k, owner := range c.Owners {
		ownerShare := decimal.NewFromFloat(float64(owner.Shares))
		p := ownerShare.Div(totalSharesDecimal)
		// for more exact ownership amounts, do not round values...
		owner.OwnershipAmount = p.Shift(2).Round(2)
		c.Owners[k] = owner
	}
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
