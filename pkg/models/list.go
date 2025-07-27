package models

import "github.com/samber/lo"

type List struct {
	ID        int64  `gorm:"column:id;primaryKey;autoIncrement"`
	UserID    int64  `gorm:"column:user_id;not null;index:idx_search"`
	Name      string `gorm:"column:name;not null;size:255"`
	IsPublic  bool   `gorm:"column:is_public;not null;default:false;index:idx_search"`
	DeletedAt int64  `gorm:"column:deleted_at;default:0;index:idx_search"`
}

func (List) TableName() string {
	return "lists"
}

type ListAPI struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	IsPublic bool   `json:"is_public"`
}

func (ListAPI) FromModel(list *List, _ int) *ListAPI {
	return &ListAPI{
		ID:       list.ID,
		Name:     list.Name,
		IsPublic: list.IsPublic,
	}
}

func (ListAPI) FromModels(values []*List) []*ListAPI {
	return lo.Map(values, func(item *List, _ int) *ListAPI {
		return &ListAPI{
			ID:       item.ID,
			Name:     item.Name,
			IsPublic: item.IsPublic,
		}
	})
}
