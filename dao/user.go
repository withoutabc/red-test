package dao

import "proj/model"

func SearchUserByUsername(name string) (u model.User, err error) {
	row := DB.QueryRow("select * from user where name=?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.ID, &u.Username, &u.Password)
	return
}
func InsertRepository(m model.Management) (err error) {
	_, err = DB.Exec("insert into management(repository,name) values (?,?)", m.Repository, m.Name)
	return
}
