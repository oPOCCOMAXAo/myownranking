package models

import "github.com/samber/lo"

type ListElement struct {
	ID          int64  `gorm:"column:id;primaryKey;autoIncrement"`
	ListID      int64  `gorm:"column:list_id;not null;index"`
	Name        string `gorm:"column:name;not null;size:255"`
	ImageURL    string `gorm:"column:image_url;size:255"`
	VideoURL    string `gorm:"column:video_url;size:255"`
	Description string `gorm:"column:description;size:1024"`
}

func (ListElement) TableName() string {
	return "list_elements"
}

func (e *ListElement) PrepareForComparison(other *ListElement) {
	if e == nil || other == nil {
		return
	}

	e.ID = other.ID
}

func (e *ListElement) Equals(other *ListElement) bool {
	if e == nil || other == nil {
		return false
	}

	return *e == *other
}

type ListElementAPI struct {
	Name        string `json:"name"`
	ImageURL    string `json:"image_url"`
	VideoURL    string `json:"video_url"`
	Description string `json:"description"`
}

func (ListElementAPI) FromModel(e *ListElement, _ int) *ListElementAPI {
	return &ListElementAPI{
		Name:        e.Name,
		ImageURL:    e.ImageURL,
		VideoURL:    e.VideoURL,
		Description: e.Description,
	}
}

func (ListElementAPI) FromModels(values []*ListElement) []*ListElementAPI {
	return lo.Map(values, ListElementAPI{}.FromModel)
}

func (e *ListElementAPI) ToModel(_ int) *ListElement {
	return &ListElement{
		Name:        e.Name,
		ImageURL:    e.ImageURL,
		VideoURL:    e.VideoURL,
		Description: e.Description,
	}
}
