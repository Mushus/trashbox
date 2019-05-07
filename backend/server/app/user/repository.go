package user

// Repository ユーザーを保存するリポジトリ
type Repository interface {
	// FindByLogin ログイン名が login なユーザーを探す。
	// 見つかった場合, そのユーザーを返す
	// 見つからなかった場合、 nil と ErrUserNotFound を返す
	FindByLogin(login string) (*User, error)
	// Add ユーザーを追加する
	Add(user *User) error
}
