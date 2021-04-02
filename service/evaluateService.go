package service

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type Evaluate struct {
	Id        int64  `form:"id" json:"-"`
	UserId    string `form:"userid" json:"userid" example:"test"`
	SectionId int64  `form:"sectionid" json:"sectionid" example:"1"`
	TypeId    int64  `form:"typeid" json:"typeid"  example:"12345"`
	Type      string `form:"type" json:"type"  example:"1"`
	Status    int    `form:"status" json:"status" example:"0"`
}

func GetEvaluateByUseridAndTypeid(typ string, userid string, typeid string) entity.EvaluateMap {
	return repository.GetEvaluateByUseridAndTypeid(typ, userid, typeid)
}

func GetEvaluateIntByUseridAndTypeid(typ string, userid string, typeid int64) int {
	return repository.GetEvaluateIntByUseridAndTypeid(typ, userid, typeid)
}

func (evaluate *Evaluate) SetEvaluate() (bool, error) {
	tevaluate := &models.TUserEvaluate{
		Id:        evaluate.Id,
		UserId:    evaluate.UserId,
		SectionId: evaluate.SectionId,
		TypeId:    evaluate.TypeId,
		Type:      evaluate.Type,
		Status:    evaluate.Status,
	}
	return repository.SetEvaluate(tevaluate)
}
