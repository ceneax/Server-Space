package repository

import "xcdh.space/space/models"

func GetSetionList(ts *[]models.TSection) error {
	return models.X.Find(ts)
}

func AddSetion(ts *models.TSection) (bool, error) {
	_, err := models.X.Insert(ts)
	if err != nil {
		return false, err
	}
	return true, nil
}
