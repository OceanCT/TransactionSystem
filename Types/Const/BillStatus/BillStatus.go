package BillStatus

type BillStatus int

const (
	UnPaid   BillStatus = 0
	Paid     BillStatus = 1
	Canceled BillStatus = 2
)
