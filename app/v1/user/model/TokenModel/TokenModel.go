package TokenModel

import "main.go/tuuz"

const Table = "doudian_token"

func Api_insert(uid, token, device any) (err error) {
	db := tuuz.Db().Table(Table)
	db.Data(map[string]any{
		"uid":    uid,
		"token":  token,
		"device": device,
	})
	_, err = db.Insert()
	return
}

func Api_findByUidAndToken(uid, token any) (map[string]any, error) {
	db := tuuz.Db().Table(Table)
	db.Where("uid", "=", uid)
	db.Where("token", "=", token)
	return db.Find()
}

func Api_findByToken(token any) (map[string]any, error) {
	db := tuuz.Db().Table(Table)
	db.Where("token", "=", token)
	return db.Find()
}
