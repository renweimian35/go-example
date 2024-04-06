package jerror

import (
	"errors"
	"fmt"
	jujuerrors "github.com/juju/errors"
	"net/http"
)

type ErrPayload struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewErrPayload(err error) *ErrPayload {
	var code int
	var message string
	switch err.(type) { //nolint
	case nil:
		code = http.StatusOK
		message = "success"
	case *Err:
		var e *Err
		errors.As(err, &e) //nolint
		if e == nil {
			code = http.StatusOK
			message = "success"
		} else {
			code = e.Code
			message = e.Msg
			stack := jujuerrors.ErrorStack(err)
			//jujuerrors.ErrorStack(err)
			//stack := jujuerrors.Details(err)
			fmt.Println(stack)
		}
	default:
		code = http.StatusInternalServerError
		message = "server internal error"
		s := jujuerrors.ErrorStack(err)
		fmt.Println(s)
	}

	return &ErrPayload{
		Code:    code,
		Message: message,
	}
}
