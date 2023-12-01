package filter

import (
	"fmt"
	"reflect"
	"time"
)

type Paging[T comparable] struct {
	Page     int    `json:"page" form:"page"`
	Take     int    `json:"take" form:"take"`
	OrderBy  string `json:"orderBy" form:"orderBy"`
	IsActive bool   `json:"-"`
	Filter   T
}

func (f *Paging[T]) SetDefault() {
	f.Page = 1
	f.Take = -1
}

func (f *Paging[T]) QueryBuilder() string {
	query := " WHERE 1=1"
	if f.IsActive {
		query += " AND status=1"
	}
	ref := reflect.ValueOf(f.Filter)
	tpe := ref.Type()

	// Adding where statement
	for i := 0; i < tpe.NumField(); i++ {
		if !isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			query += " AND " + tpe.Field(i).Tag.Get("db") + "='" + fmt.Sprint(ref.Field(i).Interface()) + "'"
		}
	}

	return query
}

func (f *Paging[T]) PaginationQuery() string {
	query := ""
	// Adding OrderBy Statement
	if f.OrderBy != "" {
		query += " ORDER BY " + f.OrderBy
	}

	// Adding Limit and Offset statement
	if f.Take > 0 {
		query += " LIMIT " + fmt.Sprint(f.Take)
		if f.Page > 0 {
			query += " OFFSET " + fmt.Sprint(f.Take*(f.Page-1))
		}
	}
	return query
}

func isEmpty(check string) bool {
	return check == "0" || check == "" || check == fmt.Sprint(time.Time{})
}
