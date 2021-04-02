package repository

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
)

func GetFavoriteByUseridAndTypeid(typ string, userid string, typeid string) entity.FavoriteMap {
	sql := " type_id in " + typeid + "and type = ? and user_id = ?"
	userFavorite := make([]models.TUserFavorite, 0)
	err := models.X.Where(sql, typ, userid).Find(&userFavorite)
	if err == nil {
		favoriteMap := make(entity.FavoriteMap)
		for _, value := range userFavorite {
			favoriteMap[value.TypeId] = value.Status
		}
		return favoriteMap
	}
	return nil
}

func GetFavoriteIntByUseridAndTypeid(typ string, userid string, typeid int64) int {
	sql := " type_id = ? and type = ? and user_id = ?"
	var userFavorite *models.TUserFavorite
	has, err := models.X.Where(sql, typ, userid).Get(userFavorite)
	if err == nil && has {
		return userFavorite.Status
	}
	return 0
}

func SetFavorite(favorite *models.TUserFavorite) (bool, error) {
	var count int64 = int64(favorite.Status)
	var newStatus int
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	newStatus = favorite.Status
	favorite.Status = 0
	has, err := session.Where("type_id = ? and user_id = ? and type = ?", favorite.TypeId, favorite.UserId, favorite.Type).Get(favorite)
	//查询无收藏
	if !has && err == nil {
		_, err = session.Insert(favorite)
		if err != nil {
			session.Rollback()
			return false, err
		}
	}
	//查询有评价
	if has && err == nil {
		favorite.Status = newStatus
		_, err = session.Id(favorite.Id).Cols("status").Update(favorite)
		if err != nil {
			session.Rollback()
			return false, err
		}
	}
	//查询出错
	if err != nil {
		session.Rollback()
		return false, err
	}
	//增加计数
	switch favorite.Type {
	case "1":
		ttopic := &models.TTopic{TopicId: favorite.TypeId}
		has, err := models.X.Get(ttopic)
		if !has || err != nil {
			session.Rollback()
			return false, err
		}
		ttopic.Likecount = ttopic.Favoritecount + count
		_, err = session.Id(ttopic.TopicId).Cols("favoriteCount").Update(ttopic)
		if err != nil {
			session.Rollback()
			return false, err
		}
		// case "2":
		// 	var tcomment *models.TComment
		// 	tcomment.CommentId = favorite.TypeId
		// 	has, err := models.X.Get(tcomment)
		// 	if !has || err != nil {
		// 		session.Rollback()
		// 		return false, err
		// 	}
		// 	tcomment.Likecount = tcomment.
		// 	_, err = session.Id(tcomment.CommentId).Cols("likeCount", "hateCount").Update(tcomment)
		// 	if err != nil {
		// 		session.Rollback()
		// 		return false, err
		// 	}
		// case "3":
		// 	var treply *models.TReply
		// 	treply.ReplyId = favorite.TypeId
		// 	has, err := models.X.Get(treply)
		// 	if !has || err != nil {
		// 		session.Rollback()
		// 		return false, err
		// 	}
		// 	treply.Likecount = treply.Likecount + like
		// 	_, err = session.Id(treply.ReplyId).Cols("likeCount", "hateCount").Update(treply)
		// 	if err != nil {
		// 		session.Rollback()
		// 		return false, err
		// 	}
	}
	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
