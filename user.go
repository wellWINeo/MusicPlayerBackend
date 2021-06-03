package MusicPlayerBackend

type User struct {
	Id         int    `json:"-" db:"id_user"`
	Username   string `json:"username" binding:"required" db:"username"`
	Email      string `json:"email" binding:"required" db:"email"`
	Password   string `json:"password" db:"passwd"`
	IsPremium  bool   `json:"is_premium" db:"is_premium"`
	IsVerified bool   `json:"is_verified,string" db:"is_verified"`
	Referal    int    `json:"referal,string"`
}
