package domain

type CreateLoanReq struct {
	UserId string `json:"userId"`
	Amount int    `json:"amount"`
}

type CreateLoanRes struct {
	LoanId int64 `json:"loanId"`
}
