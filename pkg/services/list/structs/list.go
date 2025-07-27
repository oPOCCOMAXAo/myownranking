package structs

import "github.com/opoccomaxao/myownranking/pkg/models"

type ListParams struct {
	UserID     int64
	Limit      int
	Offset     int
	OnlyPublic bool
	WithTotal  bool
}

type ListResult struct {
	Total int64
	Lists []*models.List
}
