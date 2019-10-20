package data

import (
	"testing"
	"time"

	"github.com/shopspring/decimal"
	"gotest.tools/assert"
)

func TestAdd(t *testing.T) {

	ownerA := Owner{Shares: 10, Investor: "Billy Joel", CashPaid: createDecimal("1200.00")}
	ownerB := Owner{Shares: 30, Investor: "Billy Joel", CashPaid: createDecimal("200.00")}
	ownerC := Owner{Shares: 300, Investor: "Christie Brinkley", CashPaid: createDecimal("2500.00")}

	err := ownerC.add(&ownerA)
	if err == nil {
		t.Error("Should not be able to add owner object with different investors")
	}
	assert.ErrorContains(t, err, "Investors")

	err = ownerA.add(&ownerB)
	assert.NilError(t, err)

	assert.Equal(t, ownerA.Investor, "Billy Joel")
	assert.Equal(t, ownerA.Shares, 40)
	assert.Equal(t, ownerA.CashPaid.StringFixed(2), "1400.00")
}

func TestAddOwner(t *testing.T) {

	ownerA := Owner{Shares: 10, Investor: "Billy Joel", CashPaid: createDecimal("1200.00"), Date: createTime("2018-01-02")}
	ownerB := Owner{Shares: 30, Investor: "Billy Joel", CashPaid: createDecimal("200.00"), Date: createTime("2018-01-02")}
	ownerC := Owner{Shares: 300, Investor: "Christie Brinkley", CashPaid: createDecimal("2500.00"), Date: createTime("2018-01-02")}

	cap := CapTable{Date: createTime("2019-01-01"), Owners: make(OwnerList)}
	err := cap.AddInvestor(&ownerA)
	assert.NilError(t, err)
	err = cap.AddInvestor(&ownerB)
	assert.NilError(t, err)
	err = cap.AddInvestor(&ownerC)
	assert.NilError(t, err)
	assert.Equal(t, len(cap.Owners), 2)
	assert.Equal(t, cap.Owners["Billy Joel"].CashPaid.StringFixed(2), "1400.00")

	//Incorrect Date
	ownerD := Owner{Shares: 500, Investor: "Charles Barkley", CashPaid: createDecimal("2500.00"), Date: createTime("2019-01-02")}
	err = cap.AddInvestor(&ownerD)
	assert.NilError(t, err)
	assert.Equal(t, len(cap.Owners), 2)
}

// //////
// Helper functions
// //////

func createDecimal(dec string) decimal.Decimal {
	value, _ := decimal.NewFromString(dec)
	return value
}

func createTime(str string) time.Time {
	t, _ := time.Parse(DateFormat, str)
	return t
}
