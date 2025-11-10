package models

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
