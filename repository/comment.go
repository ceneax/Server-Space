package repository

import (
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xorm.io/core"
)

func GetCommentList(pageinfo *entity.Pageinfo, topicid int64, tc *[]entity.CommentExtends) error {
	comment := new(models.TComment)
	total, _ := models.X.Join("INNER", "t_topic_comment", "t_topic_comment.comment_id = t_comment.comment_id").Where("topic_id =?", topicid).Count(comment)
	pagesize, pagebegin := pageinfo.GetLimit(total)
	pageinfo.Total = total
	return models.X.Join("INNER", "t_user", "t_user.user_id = t_comment.user_id").Join("INNER", "t_topic_comment", "t_topic_comment.comment_id = t_comment.comment_id").Where("topic_id =?", topicid).Limit(pagesize, pagebegin).Find(tc)
}

func GetCommentById(commentid int64) (*models.TComment, bool, error) {
	var comment models.TComment
	has, err := models.X.ID(commentid).Get(&comment)
	return &comment, has, err
}

func AddComment(tc *models.TComment, ttc *models.TTopicComment) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(tc)
	if err != nil {
		session.Rollback()
		return false, err
	}
	ttc.CommentId = tc.CommentId
	_, err = session.Insert(ttc)
	if err != nil {
		session.Rollback()
		return false, err
	}
	var topic models.TTopic
	topic.TopicId = ttc.TopicId
	has, err := models.X.Get(&topic)
	if err != nil || has == false {
		session.Rollback()
		return false, err
	}
	topic.Commentcount = topic.Commentcount + 1
	topic.Commnetat = time.Now()
	_, err = models.X.ID(topic.TopicId).Update(&topic)
	if err != nil {
		session.Rollback()
		return false, err
	}
	AddCommentNotice(session, ttc, tc.UserId)
	err = session.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}
func UpdateComment(tr *models.TComment) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Id(tr.CommentId).Update(tr)
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

func DeleteComment(ttc *models.TTopicComment) (bool, error) {
	affect, err := models.X.Id(core.PK{ttc.TopicId, ttc.CommentId}).Delete(ttc)
	if affect == 1 {
		return true, nil
	} else {
		return false, err
	}
}

func GetCommentUser(topicid int64) string {
	tc := models.TComment{
		CommentId: topicid,
	}
	has, err := models.X.Get(&tc)
	if has && err == nil {
		return tc.UserId
	}
	return ""
}
