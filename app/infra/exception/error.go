package exception

type Error struct {
	Code int
	Err  error
}

const (
	IntenalError    = 500
	BadRequestError = 400
)
