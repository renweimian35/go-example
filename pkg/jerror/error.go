package jerror

import (
	"errors"
	jujuerrors "github.com/juju/errors"
)

type Err struct {
	*jujuerrors.Err
	Code int
	Msg  string
}

// New 根据错误码生成一个错误，并加入自定义信息
func New(code int, message string) error {
	juErr := jujuerrors.NewErr(message)
	juErr.SetLocation(1)
	return &Err{
		Code: code,
		Msg:  message,
		Err:  &juErr,
	}
}

// Trace 从里向外原样返回错误时，必须调用这个方法，以记录里层错误的栈信息
func Trace(other error) error {
	if other == nil {
		return nil
	}
	newErr := new(Err)
	var err *Err
	if errors.As(other, &err) {
		newErr.Code = err.Code
		newErr.Msg = err.Msg
	}
	errors.As(jujuerrors.Trace(other), &newErr.Err)
	newErr.SetLocation(1)
	return newErr

}
