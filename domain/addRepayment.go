package domain

type AddRepaymentReq struct {
	UserId     string  `json:"userId"`
	LoanId     int     `json:"loanId"`
	TermAmount float64 `json:"termAmount"`
	TermNo     int     `json:"termNo"`
}

type AddRepaymentRes struct {
	LoanId  int     `json:"loanId"`
	Balance float64 `json:"balance"`
}
