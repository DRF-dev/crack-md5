package internal

type CustomError struct {
	Message string
	Code    int
}

func (c *CustomError) Error() string {
	return c.Message
}
