package customError

type CustomError struct {
	errorMsg  error
	typeError string
}

func (err *CustomError) Error() string {
	return err.errorMsg.Error()
}

func (err *CustomError) ErrorType() string {
	return err.typeError
}

func ThrowError(errorType string, err error) *CustomError {
	return &CustomError{
		typeError: errorType,
		errorMsg:  err,
	}
}
