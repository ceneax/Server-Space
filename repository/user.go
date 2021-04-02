package repository

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
)

func GetByUserId(UserID string) (*models.TUser, bool, error) {
	var user models.TUser
	has, err := models.X.ID(UserID).Get(&user)
	return &user, has, err
}

func AddUser(tuser *models.TUser, tuserInfo *models.TUserInfo) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(tuser)
	if err != nil {
		session.Rollback()
		return false, err
	}

	_, err = session.Insert(tuserInfo)
	if err != nil {
		session.Rollback()
		return false, err
	}

	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserInfo(userid string, userinfo *[]entity.UserInfoExtends) (bool, error) {
	return true, models.X.Join("INNER", "t_user", "t_user.user_id = t_user_info.user_id").Where("t_user.user_id =?", userid).Find(userinfo)
}

func UpdateUserInfo(userinfo *models.TUserInfo) (bool, error) {
	affected, err := models.X.Id(userinfo.UserId).Update(userinfo)
	if affected == 0 {
		return false, err
	}
	return true, err
}

func UpdateUser(user *models.TUser) (bool, error) {
	affected, err := models.X.Id(user.UserId).Update(user)
	if affected == 0 {
		return false, err
	}
	return true, err
}
