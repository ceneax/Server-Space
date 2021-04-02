package repository

import (
	"time"

	"github.com/go-xorm/xorm"
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
)

func GetNoticeList(userid string, no *[]entity.NoticeExtends) error {
	return models.X.Join("INNER", "t_user", "t_user.user_id = t_notice.causeuser_id").Where("t_notice.user_id =?", userid).Find(no)
}

func SetNoticeRead(userid string) (bool, error) {
	sql := "update t_notice set isRead = ? where user_id = ?"
	_, err := models.X.Exec(sql, 1, userid)
	if err == nil {
		return true, nil
	}
	return false, err
}

func AddReplyNotice(session *xorm.Session, tr *models.TReply) (bool, error) {
	var notice models.TNotice
	if tr.Pid != 0 {
		var reply models.TReply
		reply.ReplyId = tr.Pid
		has, err := models.X.Get(&reply)
		if err != nil || has == false {
			session.Rollback()
			return false, err
		}
		notice.ResultId = reply.ReplyId
		notice.Type = "R"
		notice.UserId = reply.UserId
		notice.CauseId = tr.ReplyId
		notice.CauseuserId = tr.UserId
		notice.Createdat = time.Now()
		_, err = session.Insert(&notice)
		if err != nil {
			session.Rollback()
			return false, err
		}
	} else {
		var comment models.TComment
		comment.CommentId = tr.CommentId
		has, err := models.X.Get(&comment)
		if err != nil || has == false {
			session.Rollback()
			return false, err
		}
		notice.ResultId = tr.CommentId
		notice.Type = "C"
		notice.UserId = comment.UserId
		notice.CauseId = tr.ReplyId
		notice.CauseuserId = tr.UserId
		notice.Createdat = time.Now()
		_, err = session.Insert(&notice)
		if err != nil {
			session.Rollback()
			return false, err
		}
	}
	return true, nil
}

func AddCommentNotice(session *xorm.Session, ttc *models.TTopicComment, commentuserid string) (bool, error) {
	var notice models.TNotice
	var topic models.TTopic
	topic.TopicId = ttc.TopicId
	has, err := models.X.Get(&topic)
	if err != nil || has == false {
		session.Rollback()
		return false, err
	}
	notice.ResultId = ttc.TopicId
	notice.Type = "T"
	notice.UserId = topic.UserId
	notice.CauseId = ttc.CommentId
	notice.CauseuserId = commentuserid
	notice.Createdat = time.Now()
	_, err = session.Insert(&notice)
	if err != nil {
		session.Rollback()
		return false, err
	}

	return true, nil
}
