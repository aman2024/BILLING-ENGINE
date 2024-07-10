package domain

type ApproveLoanReq struct {
	AdminId string `json:"adminId"`
	LoanId  int    `json:"loanId"`
}

type ApproveLoanRes struct {
	LoanId int64  `json:"loanId"`
	Status string `json:"status"`
}
