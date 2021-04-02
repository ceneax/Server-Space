package service

import (
	"errors"
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type Comment struct {
	CommentId int64     `form:"commentid" json:"-"`
	Content   string    `form:"content" json:"content"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Replyedat time.Time `form:"replyedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Havereply int       `form:"havereply" json:"-"`
	UserId    string    `form:"userid" json:"userid"`
	TopicId   int64     `form:"topicid" json:"topicid"`
}
type UpdateComment struct {
	CommentId int64     `form:"commentid" json:"commentid"`
	Content   string    `form:"content" json:"content"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Replyedat time.Time `form:"replyedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Havereply int       `form:"havereply" json:"-"`
	UserId    string    `form:"userid" json:"-"`
	TopicId   int64     `form:"topicid" json:"-"`
}
type DeleteComment struct {
	CommentId int64     `form:"commentid" json:"commentid"`
	Content   string    `form:"content" json:"-"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Replyedat time.Time `form:"replyedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Havereply int       `form:"havereply" json:"-"`
	UserId    string    `form:"userid" json:"-"`
	TopicId   int64     `form:"topicid" json:"topicid"`
}

//GetComment 根据页码获取评论
func GetCommentList(pageinfo *entity.Pageinfo, topicid int64) []entity.CommentExtends {
	tc := make([]entity.CommentExtends, 0)
	if repository.GetCommentList(pageinfo, topicid, &tc) == nil {
		return tc
	}
	return nil
}

//GetCommentById 根据ID获取回复
func GetCommentById(commentid int64) (*models.TComment, bool, error) {
	return repository.GetCommentById(commentid)
}

//AddComment 添加评论
func (t *Comment) AddComment() (bool, error) {
	tComment := &models.TComment{
		CommentId: t.CommentId,
		Content:   t.Content,
		Createdat: time.Now(),
		Updatedat: time.Now(),
		Replyedat: time.Now(),
		Ip:        t.Ip,
		Havereply: t.Havereply,
		UserId:    t.UserId,
	}
	tTopicComment := &models.TTopicComment{
		CommentId: t.CommentId,
		TopicId:   t.TopicId,
	}
	return repository.AddComment(tComment, tTopicComment)
}

func (t *UpdateComment) UpdateComment() (bool, error) {
	userid := repository.GetCommentUser(t.CommentId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tComment := &models.TComment{
		CommentId: t.CommentId,
		Content:   t.Content,
		Ip:        t.Ip,
		Updatedat: time.Now(),
	}
	return repository.UpdateComment(tComment)
}

func (t *DeleteComment) DeleteComment() (bool, error) {
	userid := repository.GetCommentUser(t.CommentId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tTopicComment := &models.TTopicComment{
		CommentId: t.CommentId,
		TopicId:   t.TopicId,
	}
	return repository.DeleteComment(tTopicComment)
}
