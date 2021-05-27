package MusicPlayerBackend

type User struct {
	Id        int    `json:"-" db:"id_user"`
	Username  string `json:"username" binding:"required" db:"username"`
	Email     string `json:"email" binding:"required" db:"email"`
	Password  string `json:"password" binding:"required" db:"passwd"`
	IsPremium bool   `json:"is_premium" db:"is_premium"`
}
