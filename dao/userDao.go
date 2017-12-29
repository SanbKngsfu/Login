package dao

import "login/model"

type UserDao struct{}

func (u *UserDao) QueryList(sql string) (objs []*model.User) {
	db := InitDB()
	defer db.Close()

	rows, res := CheckedResult(db.Query(sql))

	id := res.Map("id")
	name := res.Map("name")
	password := res.Map("password")

	length := len(rows)
	objs = make([]*model.User, length)

	for i, row := range rows {
		objs[i] = &model.User{
			row.Int(id),
			row.Str(name),
			row.Str(password),
		}
	}

	return objs
}

func (u *UserDao) Save(sql string) {
	db := InitDB()
	defer db.Close()

	CheckedResult(db.Query(sql))
}
