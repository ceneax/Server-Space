package service

import (
	"xcdh.space/space/entity"
	"xcdh.space/space/models"
	"xcdh.space/space/repository"
)

type Favorite struct {
	Id        int64  `form:"id" json:"-"`
	UserId    string `form:"userid" json:"userid" example:"test"`
	SectionId int64  `form:"sectionid" json:"sectionid" example:"1"`
	TypeId    int64  `form:"typeid" json:"typeid" example:"1"`
	Type      string `form:"type" json:"type" example:"234"`
	Status    int    `form:"status" json:"status" example:"1"`
}

func GetFavoriteByUseridAndTypeid(typ string, userid string, typeid string) entity.FavoriteMap {
	return repository.GetFavoriteByUseridAndTypeid(typ, userid, typeid)
}

func GetFavoriteIntByUseridAndTypeid(typ string, userid string, typeid int64) int {
	return repository.GetFavoriteIntByUseridAndTypeid(typ, userid, typeid)
}

func (favorite *Favorite) SetFavorite() (bool, error) {
	tfavorite := &models.TUserFavorite{
		Id:        favorite.Id,
		UserId:    favorite.UserId,
		SectionId: favorite.SectionId,
		TypeId:    favorite.TypeId,
		Type:      favorite.Type,
		Status:    favorite.Status,
	}
	return repository.SetFavorite(tfavorite)
}
