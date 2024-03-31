package repository

import (
	"gorm.io/gorm"
)

func appendConditionsToQuery(query *gorm.DB, conditions map[string]any) *gorm.DB {
	for k, v := range conditions {
		query = query.Where(k, v)
	}

	return query
}
