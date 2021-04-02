package service

import (
	"errors"
	"time"

	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
	"xcdh.space/space/utils"
)

type Topic struct {
	TopicId        int64     `form:"topicid" json:"-"`
	SectionId      int64     `form:"sectionid" json:"sectionid"`
	Title          string    `form:"title" json:"title"`
	Shortcontent   string    `form:"shortcontent" json:"-"`
	UserId         string    `form:"userid" json:"userid"`
	Istop          int       `form:"istop" json:"istop"`
	Commentcount   int64     `form:"commentcount" json:"-"`
	Createdat      time.Time `form:"createdat" json:"-"`
	Updatedat      time.Time `form:"updatedat" json:"-"`
	Commnetat      time.Time `form:"commentat" json:"-"`
	Content        string    `form:"content" json:"content"`
	EvaluateStatus int       `json:"evaluateStatus"`
	FavoriteStatus int       `json:"favoriteStatus"`
}
type UpdateTopic struct {
	TopicId      int64     `form:"topicid" json:"topic"`
	SectionId    int64     `form:"sectionid" json:"sectionid"`
	Title        string    `form:"title" json:"title"`
	Shortcontent string    `form:"shortcontent" json:"-"`
	UserId       string    `form:"userid" json:"userid"`
	Istop        int       `form:"istop" json:"istop"`
	Commentcount int64     `form:"commentcount" json:"-"`
	Createdat    time.Time `form:"createdat" json:"-"`
	Updatedat    time.Time `form:"updatedat" json:"-"`
	Commnetat    time.Time `form:"commentat" json:"-"`
	Content      string    `form:"content" json:"content"`
}

type DeleteTopic struct {
	TopicId      int64     `form:"topicid" json:"topic"`
	SectionId    int64     `form:"sectionid" json:"sectionid"`
	Title        string    `form:"title" json:"-"`
	Shortcontent string    `form:"shortcontent" json:"-"`
	UserId       string    `form:"userid" json:"userid"`
	Istop        int       `form:"istop" json:"-"`
	Commentcount int64     `form:"commentcount" json:"-"`
	Createdat    time.Time `form:"createdat" json:"-"`
	Updatedat    time.Time `form:"updatedat" json:"-"`
	Commnetat    time.Time `form:"commentat" json:"-"`
	Content      string    `form:"contnet" json:"-"`
}

//GetTopic 根据页码获取主题
func GetTopic(pageinfo *entity.Pageinfo, sectionid int64) []entity.TopicExtends {
	tc := make([]entity.TopicExtends, 0)
	if repository.GetTopicList(pageinfo, sectionid, &tc) == nil {
		return tc
	}
	return nil
}

//GetTopicById 根据Id获取主题
func GetTopicById(topicid int64) (*models.TTopic, bool, error) {
	return repository.GetTopicById(topicid)
}

//GetTopicContentById 根据主题ID获取主题内容
func GetTopicContentById(topicid int64) (*models.TTopicContent, bool, error) {
	tTopicContent := &models.TTopicContent{
		TopicId: topicid,
	}
	b, err := repository.GetTopciContent(tTopicContent)
	return tTopicContent, b, err
}

//AddTopic 添
func (t *Topic) AddTopic() (bool, error) {
	tTopic := &models.TTopic{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
		Title:     t.Title,
		UserId:    t.UserId,
		Createdat: time.Now(),
		Updatedat: time.Now(),
	}
	tTopic.Shortcontent = utils.Substring(t.Content, 100)
	tTopicContent := &models.TTopicContent{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
		Content:   t.Content,
	}
	return repository.AddTopic(tTopic, tTopicContent)
}

func (t *Topic) GetTopicContent() (bool, error) {
	tTopicContent := &models.TTopicContent{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
	}
	b, err := repository.GetTopciContent(tTopicContent)
	t.Content = tTopicContent.Content
	return b, err
}

func (t *UpdateTopic) UpdateTopic() (bool, error) {
	userid := repository.GetTopicUser(t.TopicId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tTopic := &models.TTopic{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
		Title:     t.Title,
		Updatedat: time.Now(),
	}
	tTopic.Shortcontent = utils.Substring(t.Content, 100)
	tTopicContent := &models.TTopicContent{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
		Content:   t.Content,
	}
	return repository.UpdateTopic(tTopic, tTopicContent)
}

func (t *DeleteTopic) DeleteTopic() (bool, error) {
	userid := repository.GetTopicUser(t.TopicId)
	if userid != t.UserId {
		return false, errors.New("用户名不一致，无法删除")
	}
	tTopic := &models.TTopic{
		TopicId:   t.TopicId,
		SectionId: t.SectionId,
	}
	return repository.DeleteTopic(tTopic)
}
