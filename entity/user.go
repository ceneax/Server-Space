package entity

type ChangePassword struct {
	UserId      string `form:"userid"`
	Username    string `form:"username"`
	OldPassword string `form:"oldPassword"`
	Password    string `form:"password"`
}

type Login struct {
	UserId   string `form:"userid" json:"userid"`
	Password string `form:"password" json:"password"`
}

type EvaluateMap map[int64]int

type FavoriteMap map[int64]int
