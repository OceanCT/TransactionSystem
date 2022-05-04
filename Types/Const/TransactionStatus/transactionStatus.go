package TransactionStatus

type TransactionStatus int

const (
	WaittingForApproval TransactionStatus = 1
	Approvalled         TransactionStatus = 2
	DisApprovalled      TransactionStatus = 3
)
