package repository

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xorm.io/core"
)

func GetTopicList(pageinfo *entity.Pageinfo, sectionid int64, tc *[]entity.TopicExtends) error {
	topic := new(models.TTopic)
	total, _ := models.X.Where("section_id =?", sectionid).Count(topic)
	pagesize, pagebegin := pageinfo.GetLimit(total)
	pageinfo.Total = total
	return models.X.Join("INNER", "t_user", "t_user.user_id = t_topic.user_id").Where("section_id =?", sectionid).Limit(pagesize, pagebegin).Desc("createdAt").Find(tc)
	//	return models.X.Sql("select t_topic.*,t_user.username from t_topic,t_user where t_user.user_id = t_topic.user_id and section_id = ?", sectionid).Limit(pagesize, pagebegin).Find(tc)
}

func GetTopicById(topicid int64) (*models.TTopic, bool, error) {
	var topic models.TTopic
	has, err := models.X.ID(topicid).Get(&topic)
	return &topic, has, err
}

func GetTopciContent(tc *models.TTopicContent) (bool, error) {
	has, err := models.X.Get(tc)
	return has, err
}

func AddTopic(to *models.TTopic, tc *models.TTopicContent) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Insert(to)
	if err != nil {
		session.Rollback()
		return false, err
	}
	tc.TopicId = to.TopicId
	_, err = session.Insert(tc)
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

func UpdateTopic(to *models.TTopic, tc *models.TTopicContent) (bool, error) {
	session := models.X.NewSession()
	defer session.Close()
	err := session.Begin()
	_, err = session.Id(to.TopicId).Update(to)
	if err != nil {
		session.Rollback()
		return false, err
	}
	_, err = session.Id(core.PK{tc.TopicId, tc.SectionId}).Update(tc)
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

func DeleteTopic(to *models.TTopic) (bool, error) {
	affect, err := models.X.Id(to.TopicId).Delete(to)
	if affect == 1 {
		return true, nil
	} else {
		return false, err
	}
}

func GetTopicUser(topicid int64) string {
	tc := models.TTopic{
		TopicId: topicid,
	}
	has, err := models.X.Get(&tc)
	if has && err == nil {
		return tc.UserId
	}
	return ""
}
