package structs

import "github.com/opoccomaxao/myownranking/pkg/models"

type ListParams struct {
	UserID      int64
	ListIDs     []int64
	Limit       int
	Offset      int
	OnlyPublic  bool
	OnlyDeleted bool
	WithTotal   bool
}

type ListResult struct {
	Total int64
	Lists []*models.List
}

type SingleListParams struct {
	ID     int64
	UserID int64
}

type UpdateListParams struct {
	ID        int64
	UserID    int64
	Name      *string
	IsPublic  *bool
	DeletedAt *int64
}
