package service

import (
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type Section struct {
	SectionId     int64  `form:"sectionid" json:"-"`
	Name          string `form:"name" json:"name"`
	Clickrate     int64  `form:"clickrate" json:"-"`
	Posts         int64  `form:"posts" json:"-"`
	Classfication string `form:"classfication" json:"-"`
}

//GetAllSection 获取所有板块
func GetAllSection() []models.TSection {
	ts := make([]models.TSection, 0)
	if repository.GetSetionList(&ts) == nil {
		return ts
	}
	return nil
}

//AddSetion 添加板块
func (s *Section) AddSetion() (bool, error) {
	tSection := &models.TSection{Name: s.Name, Clickrate: 0, Posts: 0}
	return repository.AddSetion(tSection)
}
