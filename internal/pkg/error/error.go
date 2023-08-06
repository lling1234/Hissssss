package error

// api
const (
	Success       = 200
	BadRequest    = 400
	Unauthorized  = 401
	Forbidden     = 403
	InternalError = 500
)

// biz
const (
	SignUpSuccess  = 2001
	SignInSuccess  = 2002
	SendMsgSuccess = 2003

	SignUpFailure  = 4001
	SignInFailure  = 4002
	SendMsgFailure = 4003
)
