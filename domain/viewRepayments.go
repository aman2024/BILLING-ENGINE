package domain

type ViewRepaymentEntity struct {
	Id              int     `json:"id"`
	LoanId          int     `json:"loanId"`
	RepaymentAmount float64 `json:"repaymentAmount"`
	TermNo          int     `json:"termNo"`
	RepaymentDate   string  `json:"repaymentDate"`
	Status          string  `json:"status"`
}
