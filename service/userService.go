package service

import (
	"errors"
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type User struct {
	UserId        string    `form:"userid" json:"userid"`
	Username      string    `form:"username" json:"username"`
	Password      string    `form:"password" json:"password"`
	Sex           int       `form:"sex" json:"sex"`
	Email         string    `form:"email" json:"email"`
	Avatar        string    `form:"avatar" json:"avatar"`
	Emailverified int       `form:"emailverified" json:"emailverified"`
	Createdat     time.Time `form:"createdat" json:"-"`
	Updatedat     time.Time `form:"updatedat" json:"-"`
}

//GetByUserId 根据userid获取用户基本信息
func GetByUserId(userid string) (models.TUser, bool, error) {
	tuser, b, err := repository.GetByUserId(userid)
	return *tuser, b, err
}

//CheckLogin 检查用户登陆信息
func (u *User) CheckLogin() bool {
	user, b, err := repository.GetByUserId(u.UserId)
	if b == true && u.Password == user.Password && err == nil {
		u.Username = user.Username
		return b
	} else {
		return false
	}
}

//Register 用户注册
func (u *User) Register() (bool, error) {
	_, b, err := repository.GetByUserId(u.UserId)
	if b == false && err == nil {
		tuser := &models.TUser{UserId: u.UserId, Username: u.Username, Password: u.Password, Avatar: u.Avatar}
		tuserInfo := &models.TUserInfo{
			UserId: u.UserId, Sex: u.Sex,
			Emailverified: 0,
			Email:         u.Email,
			Createdat:     time.Now(),
			Updatedat:     time.Now(),
		}
		return repository.AddUser(tuser, tuserInfo)
	} else {
		return false, errors.New("用户名已注册")
	}
}

//UpdateUserInfo 修改用户信息
func (u *User) UpdateUserInfo() (bool, error) {
	tuserInfo := &models.TUserInfo{
		UserId:        u.UserId,
		Sex:           u.Sex,
		Emailverified: u.Emailverified,
		Email:         u.Email,
		//Avatar:        u.Avatar,
		Updatedat: time.Now(),
	}
	tuser, b, err := repository.GetByUserId(u.UserId)
	if b == true && err == nil {
		if tuser.Username != u.Username || tuser.Avatar != u.Avatar {
			tuser.Username = u.Username
			tuser.Avatar = u.Avatar
			b, err = repository.UpdateUser(tuser)
			if b == false && err != nil {
				return b, err
			}
		}
	}
	return repository.UpdateUserInfo(tuserInfo)
}

func UpdateUserPassword(user entity.ChangePassword) (bool, error) {
	tuser, b, err := repository.GetByUserId(user.UserId)
	if user.OldPassword == tuser.Password && b == true && err == nil {
		tuser.Password = user.Password
		return repository.UpdateUser(tuser)
	} else {
		return false, errors.New("旧密码不正确")
	}

}

func GetUserInfo(UserId string) []entity.UserInfoExtends {
	userinfo := make([]entity.UserInfoExtends, 0)
	if _, err := repository.GetUserInfo(UserId, &userinfo); err == nil {
		return userinfo
	}
	return nil
}
