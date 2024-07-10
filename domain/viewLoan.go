package domain

import "database/sql"

type ViewLoanReq struct {
	UserId string `json:"userId"`
	LoanId int    `json:"loanId"`
}
type ViewDetailedLoanRes struct {
	LoanInfo []ViewDetailedLoanEntity `json:"loanInfo"`
}

type ViewDetailedLoanEntity struct {
	LoanId             int64                 `json:"loanId"`
	Amount             int                   `json:"amount"`
	InitialTerm        int                   `json:"initialTerm"`
	Rate               int                   `json:"rate"`
	RepaymentAmount    int                   `json:"repaymentAmount"`
	Status             string                `json:"status"`
	AmountPaid         float64               `json:"amountPaid"`
	IsDelinquent       bool                  `json:"isDelinquent"`
	TotalTerms         int                   `json:"totalTerms"`
	OutstandingBalance float64               `json:"outstandingBalance"`
	SkippedTerms       int                   `json:"skippedTerms"`
	CreatedAt          string                `json:"createdAt"`
	UpdatedAt          string                `json:"updatedAt"`
	ApprovedAt         sql.NullString        `json:"approvedAt"`
	RepaymentsInfo     []ViewRepaymentEntity `json:"repaymentsInfo"`
}

type ViewLoanRes struct {
	LoanInfo []ViewLoanEntity `json:"loanInfo"`
}
type ViewLoanEntity struct {
	LoanId          int64          `json:"loanId"`
	Amount          int            `json:"amount"`
	InitialTerm     int            `json:"initialTerm"`
	Rate            int            `json:"rate"`
	RepaymentAmount int            `json:"repaymentAmount"`
	Status          string         `json:"status"`
	AmountPaid      float64        `json:"amountPaid"`
	IsDelinquent    bool           `json:"isDelinquent"`
	TotalTerms      int            `json:"totalTerms"`
	CreatedAt       string         `json:"createdAt"`
	UpdatedAt       string         `json:"updatedAt"`
	ApprovedAt      sql.NullString `json:"approvedAt"`
}
