package service

import (
	"proj/dao"
	"proj/model"
)

func SearchUserByUsername(name string) (u model.User, err error) {
	u, err = dao.SearchUserByUsername(name)
	return
}
func AddRepository(m model.Management) error {
	err := dao.InsertRepository(m)
	return err
}
