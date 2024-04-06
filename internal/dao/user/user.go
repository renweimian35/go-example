package user

import (
	"go-example/internal/entity"
	"go-example/pkg/jerror"
)

func FindUserById(userID int64) (*entity.User, error) {
	u, err := findById(userID)
	if err != nil {
		return nil, jerror.Trace(err)
	}
	return u, nil
}

func findById(userID int64) (*entity.User, error) {
	if userID == 0 {
		return nil, jerror.New(2032, "userid empty")
	}
	u := entity.User{
		ID:   userID,
		Name: "zhangSan",
		Age:  30,
	}
	return &u, nil
}
