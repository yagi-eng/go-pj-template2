package myerror

type errorCode uint

const (
	BadRequest    = errorCode(100)
	NotFound      = errorCode(101)
	InternalError = errorCode(200)
	GormError     = errorCode(201)
)
