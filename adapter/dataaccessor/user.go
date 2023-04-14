package dataaccessor

import "github.com/pepnova-9/cqrs-sample/domain"

type User struct {
}

// Create はユーザーを作成する
func (u *User) Create(entity domain.User) error {
	_ = `
	INSERT xxx INTO users;
`
	return nil
}

// Find はユーザーを取得する
func (u *User) Find(id domain.UserID) (domain.User, error) {
	_ = `
	SELECT * FROM users WHERE id = ${id};
`
	// DataModelからEntityに変換する

	return domain.User{}, nil
}
