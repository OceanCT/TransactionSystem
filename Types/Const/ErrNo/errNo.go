package ErrNo

type ErrNo int

const (
	OK                 ErrNo = 0
	DBConnectionError  ErrNo = 1
	ParamInvalid       ErrNo = 2
	UserHasExisted     ErrNo = 3
	WrongPassword      ErrNo = 4
	LoginRequired      ErrNo = 5
	MoneyNotEnough     ErrNo = 6
	VerifyCodeNotValid ErrNo = 7
	PermDenied         ErrNo = 8
	UsernameNull       ErrNo = 9
	UserNotExisted     ErrNo = 10
	UnknownError       ErrNo = 255
)
