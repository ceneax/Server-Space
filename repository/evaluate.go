package repository

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
)

func GetEvaluateByUseridAndTypeid(typ string, userid string, typeid string) entity.EvaluateMap {
	sql := " type_id in " + typeid + "and type = ? and user_id = ?"
	userEvaluate := make([]models.TUserEvaluate, 0)
	err := models.X.Where(sql, typ, userid).Find(&userEvaluate)
	if err == nil {
		evaluateMap := make(entity.EvaluateMap)
		for _, value := range userEvaluate {
			evaluateMap[value.TypeId] = value.Status
		}
		return evaluateMap
	}
	return nil
}

func GetEvaluateIntByUseridAndTypeid(typ string, userid string, typeid int64) int {
	sql := " type_id = ? and type = ? and user_id = ?"
	var userEvaluate *models.TUserEvaluate
	has, err := models.X.Where(sql, typ, userid).Get(userEvaluate)
	if err == nil && has {
		return userEvaluate.Status
	}
	return 0
}

func SetEvaluate(userEvaluate *models.TUserEvaluate) (bool, error) {
	var like, hate int64
	var newStatus int
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	newStatus = userEvaluate.Status
	userEvaluate.Status = 0
	has, err := session.Where("type_id = ? and user_id = ? and type = ?", userEvaluate.TypeId, userEvaluate.UserId, userEvaluate.Type).Get(userEvaluate)
	//查询无评价
	if !has && err == nil {
		if newStatus == 1 {
			like = 1
		}
		if newStatus == 2 {
			hate = 1
		}
		_, err = session.Insert(userEvaluate)
		if err != nil {
			session.Rollback()
			return false, err
		}
	}
	//查询有评价
	if has && err == nil {
		switch newStatus {
		case -1:
			if userEvaluate.Status == 1 {
				like = -1
			}
			if userEvaluate.Status == 2 {
				hate = -1
			}
		case 1:
			if userEvaluate.Status == -1 {
				like = 1
			}
			if userEvaluate.Status == 2 {
				like = 1
				hate = -1
			}
		case 2:
			if userEvaluate.Status == -1 {
				hate = 1
			}
			if userEvaluate.Status == 1 {
				hate = 1
				like = -1
			}
		}
		userEvaluate.Status = newStatus
		_, err = session.Id(userEvaluate.Id).Cols("status").Update(userEvaluate)
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
	switch userEvaluate.Type {
	case "1":
		ttopic := &models.TTopic{TopicId: userEvaluate.TypeId}
		has, err := models.X.Get(ttopic)
		if !has || err != nil {
			session.Rollback()
			return false, err
		}
		ttopic.Likecount = ttopic.Likecount + like
		ttopic.Hatecount = ttopic.Hatecount + hate
		_, err = session.Id(ttopic.TopicId).Cols("likeCount", "hateCount").Update(ttopic)
		if err != nil {
			session.Rollback()
			return false, err
		}
	case "2":
		tcomment := &models.TComment{CommentId: userEvaluate.TypeId}
		has, err := models.X.Get(tcomment)
		if !has || err != nil {
			session.Rollback()
			return false, err
		}
		tcomment.Likecount = tcomment.Likecount + like
		tcomment.Hatecount = tcomment.Hatecount + hate
		_, err = session.Id(tcomment.CommentId).Cols("likeCount", "hateCount").Update(tcomment)
		if err != nil {
			session.Rollback()
			return false, err
		}
	case "3":
		treply := &models.TReply{ReplyId: userEvaluate.TypeId}
		has, err := models.X.Get(treply)
		if !has || err != nil {
			session.Rollback()
			return false, err
		}
		treply.Likecount = treply.Likecount + like
		treply.Hatecount = treply.Hatecount + hate
		_, err = session.Id(treply.ReplyId).Cols("likeCount", "hateCount").Update(treply)
		if err != nil {
			session.Rollback()
			return false, err
		}
	}
	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
