package repository

import (
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
)

func GetReplyList(pageinfo *entity.Pageinfo, commentid int64, tr *[]entity.ReplyExtends) error {
	reply := new(models.TReply)
	total, _ := models.X.Where("comment_id =?", commentid).Count(reply)
	pagesize, pagebegin := pageinfo.GetLimit(total)
	pageinfo.Total = total
	return models.X.Join("INNER", "t_user", "t_user.user_id = t_reply.user_id").Where("comment_id =?", commentid).Limit(pagesize, pagebegin).Find(tr)
}

func GetReplyById(replyid int64) (*models.TReply, bool, error) {
	var reply models.TReply
	has, err := models.X.ID(replyid).Get(&reply)
	return &reply, has, err
}

func AddReply(tr *models.TReply) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(tr)
	if err != nil {
		session.Rollback()
		return false, err
	}
	var comment models.TComment
	comment.CommentId = tr.CommentId
	comment.Havereply = 1
	comment.Replyedat = time.Now()
	_, err = models.X.ID(comment.CommentId).Update(&comment)
	if err != nil {
		session.Rollback()
		return false, err
	}
	AddReplyNotice(session, tr)
	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateReply(tr *models.TReply) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Id(tr.ReplyId).Update(tr)
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

func DeleteReply(to *models.TReply) (bool, error) {
	affect, err := models.X.Id(to.ReplyId).Delete(to)
	if affect == 1 {
		return true, nil
	} else {
		return false, err
	}
}

func GetReplyUser(topicid int64) string {
	tc := models.TReply{
		ReplyId: topicid,
	}
	has, err := models.X.Get(&tc)
	if has && err == nil {
		return tc.UserId
	}
	return ""
}
