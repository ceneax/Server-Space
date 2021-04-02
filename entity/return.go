package entity

import "time"

type TopicBackList struct {
	TopicId        int64     `json:"topic"`
	SectionId      int64     `json:"sectionid"`
	Title          string    `json:"title"`
	Shortcontent   string    `json:"shortContent"`
	UserId         string    `json:"userid"`
	Istop          int       `json:"istop"`
	Commentcount   int64     `json:"commentcount"`
	Likecount      int64     `json:"likecount"`
	Hatecount      int64     `json:"hatecount"`
	Sharecount     int64     `json:"sharecount"`
	Favoritecount  int64     `json:"favoritecount"`
	Createdat      time.Time `json:"createdat"`
	Updatedat      time.Time `json:"updatedat"`
	Commnetat      time.Time `json:"commentat"`
	Username       string    `json:"username"`
	Avatar         string    `json:"avatar"`
	EvaluateStatus int       `json:"evaluateStatus"`
	FavoriteStatus int       `json:"favoriteStatus"`
}
type TopicBack struct {
	TopicBackList []TopicBackList `json:"topiclist"`
	Pageinfo      Pageinfo        `json:"pageinfo"`
}

type CommentBackList struct {
	CommentId      int64     `form:"commentid" json:"commentid"`
	Content        string    `form:"content" json:"content"`
	Createdat      time.Time `form:"createdat" json:"createdat"`
	Updatedat      time.Time `form:"updatedat" json:"updatedat"`
	Replyedat      time.Time `form:"replyedat" json:"replyedat"`
	Ip             int64     `form:"ip" json:"ip"`
	Havereply      int       `form:"havereply" json:"havereply"`
	UserId         string    `form:"userid" json:"userid"`
	TopicId        int64     `form:"topicid" json:"topicid"`
	Likecount      int64     `form:"likecount" json:"likecount"`
	Hatecount      int64     `form:"hatecount" json:"hatecount"`
	Username       string    `json:"username"`
	Avatar         string    `form:"avatar" json:"avatar"`
	EvaluateStatus int       `form:"evaluateStatus" json:"evaluateStatus"`
	FavoriteStatus int       `form:"favoriteStatus" json:"favoriteStatus"`
}

type CommentBack struct {
	CommentBackList []CommentBackList `json:"commentlist"`
	Pageinfo        Pageinfo          `json:"pageinfo"`
}

type ReplyBackList struct {
	ReplyId        int64     `form:"replyid" json:"replyid"`
	CommentId      int64     `form:"commentid" json:"commentid"`
	Content        string    `form:"content" json:"content"`
	Createdat      time.Time `form:"createdat" json:"createdat"`
	Updatedat      time.Time `form:"updatedat" json:"updatedat"`
	Ip             int64     `form:"ip" json:"ip"`
	Pid            int64     `form:"pid" json:"pid"`
	UserId         string    `form:"userid" json:"userid"`
	Likecount      int64     `form:"likecount" json:"likecount"`
	Hatecount      int64     `form:"hatecount" json:"hatecount"`
	Username       string    `json:"username"`
	Avatar         string    `form:"avatar" json:"avatar"`
	EvaluateStatus int       `form:"evaluateStatus" json:"evaluateStatus"`
	FavoriteStatus int       `form:"favoriteStatus" json:"favoriteStatus"`
}

type ReplyBack struct {
	ReplyBackList []ReplyBackList `json:"replylist"`
	Pageinfo      Pageinfo        `json:"pageinfo"`
}

type NoticeBackList struct {
	NoticeId      int64     `form:"noticeid" json:"noticeid"`
	ResultId      int64     `form:"resultid" json:"resultid"`
	Type          string    `form:"type" json:"type"`
	Isread        int       `form:"isread" json:"isread"`
	UserId        string    `form:"userid" json:"userid"`
	Username      string    `json:"username"`
	Avatar        string    `form:"avatar" json:"avatar"`
	Content       string    `form:"content" json:"content"`
	Createdat     time.Time `form:"createdat" json:"createdat"`
	CauseId       int64     `form:"causeid" json:"causeid"`
	CauseuserId   string    `form:"causeuserid" json:"causeuserid"`
	CauseUsername string    `json:"causeusername"`
	CauseAvatar   string    `form:"causeavatar" json:"causeavatar"`
	CauseContent  string    `form:"causecontent" json:"causecontent"`
}
