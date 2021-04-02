package service

import (
	"errors"
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type Reply struct {
	ReplyId   int64     `form:"replyid" json:"-"`
	CommentId int64     `form:"commentid" json:"commentid"`
	Content   string    `form:"content" json:"content"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Pid       int64     `form:"pid" json:"pid"`
	UserId    string    `form:"userid" json:"userid"`
}
type UpdateReply struct {
	ReplyId   int64     `form:"replyid" json:"replyid"`
	CommentId int64     `form:"commentid" json:"-"`
	Content   string    `form:"content" json:"content"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Pid       int64     `form:"pid" json:"-"`
	UserId    string    `form:"userid" json:"userid"`
}
type DeleteReply struct {
	ReplyId   int64     `form:"replyid" json:"replyid"`
	CommentId int64     `form:"commentid" json:"-"`
	Content   string    `form:"content" json:"-"`
	Createdat time.Time `form:"createdat" json:"-"`
	Updatedat time.Time `form:"updatedat" json:"-"`
	Ip        int64     `form:"ip" json:"-"`
	Pid       int64     `form:"pid" json:"-"`
	UserId    string    `form:"userid" json:"userid"`
}

//GetReply 根据页码获取评论
func GetReplyList(pageinfo *entity.Pageinfo, commentid int64) []entity.ReplyExtends {
	tr := make([]entity.ReplyExtends, 0)
	if repository.GetReplyList(pageinfo, commentid, &tr) == nil {
		return tr
	}
	return nil
}

//GetReplyById 根据Id获取回复
func GetReplyById(replyid int64) (*models.TReply, bool, error) {
	return repository.GetReplyById(replyid)
}

//AddReply 添加评论
func (t *Reply) AddReply() (bool, error) {
	tReply := &models.TReply{
		ReplyId:   t.ReplyId,
		CommentId: t.CommentId,
		Content:   t.Content,
		Createdat: time.Now(),
		Updatedat: time.Now(),
		Pid:       t.Pid,
		Ip:        t.Ip,
		UserId:    t.UserId,
	}
	return repository.AddReply(tReply)
}

func (t *UpdateReply) UpdateReply() (bool, error) {
	userid := repository.GetReplyUser(t.ReplyId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tReply := &models.TReply{
		ReplyId:   t.ReplyId,
		Content:   t.Content,
		Updatedat: time.Now(),
	}
	return repository.UpdateReply(tReply)
}

func (t *DeleteReply) DeleteReply() (bool, error) {
	userid := repository.GetReplyUser(t.ReplyId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tReply := &models.TReply{
		ReplyId: t.ReplyId,
	}
	return repository.DeleteReply(tReply)
}
