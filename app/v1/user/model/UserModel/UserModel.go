package UserModel

import "main.go/tuuz"

const Table = "doudian_user"

func Api_insert(username, mail, password any) (id int64, err error) {
	db := tuuz.Db().Table(Table)
	db.Data(map[string]any{
		"username": username,
		"mail":     mail,
		"password": password,
	})
	return db.InsertGetId()
}

func Api_findByEmail(mail any) (data map[string]any, err error) {
	db := tuuz.Db().Table(Table)
	db.Where("mail", "=", mail)
	return db.Find()
}
