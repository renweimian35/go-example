package user

import (
	"go-example/internal/dao/user"
	"go-example/internal/entity"
	"go-example/pkg/jerror"
)

func ById(id int64) (*entity.User, error) {
	u, err := user.FindUserById(id)
	if err != nil {
		return nil, jerror.Trace(err)
	}
	return u, nil
}
