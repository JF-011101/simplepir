package errno

var (
	Success    = NewErrNo(0, "OK")
	ErrUnknown = NewErrNo(100001, "Internal server error")
	ErrBind    = NewErrNo(10002, "Error occurred while binding the request body to the struct")

	ErrPiNotFound = NewErrNo(110001, "User not find")
)
