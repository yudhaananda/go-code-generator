package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Query[T comparable] struct {
	Model T
}

func (q *Query[T]) BuildTableMember() string {
	member := ""

	ref := reflect.ValueOf(q.Model)
	tpe := ref.Type()

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if isQueryNeedComa {
			member += ", "
		}
		member += tpe.Field(i).Tag.Get("db")
		isQueryNeedComa = true
	}
	return member
}

func (q *Query[T]) BuildCreateQuery() string {
	query := " "

	ref := reflect.ValueOf(q.Model)
	tpe := ref.Type()

	table := "("
	values := "("

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if !isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			if isQueryNeedComa {
				table += ", "
				values += ", "
			}
			table += tpe.Field(i).Tag.Get("db")

			tempValue := fmt.Sprint(ref.Field(i).Interface())
			if fmt.Sprint(ref.Field(i).Type()) == "time.Time" {
				val := fmt.Sprint(ref.Field(i).Interface())
				vals := strings.Split(val, " ")
				tempValue = vals[0] + " " + vals[1]
			}
			isQueryNeedComa = true
			values += "'" + tempValue + "'"
		}
		if i == tpe.NumField()-1 {
			table += ")"
			values += ")"
		}
	}
	query += table + " VALUES " + values

	return query
}

func (q *Query[T]) BuildUpdateQuery(id int) string {
	query := " SET "

	ref := reflect.ValueOf(q.Model)
	tpe := ref.Type()

	isQueryNeedComa := false
	for i := 0; i < tpe.NumField(); i++ {
		if !isEmpty(fmt.Sprint(ref.Field(i).Interface())) {
			if isQueryNeedComa {
				query += ", "
			}

			tempValue := fmt.Sprint(ref.Field(i).Interface())
			if fmt.Sprint(ref.Field(i).Type()) == "time.Time" {
				val := fmt.Sprint(ref.Field(i).Interface())
				vals := strings.Split(val, " ")
				tempValue = vals[0] + " " + vals[1]
			}
			isQueryNeedComa = true

			query += tpe.Field(i).Tag.Get("db") + "='" + tempValue + "'"
		}
	}
	query += " WHERE id=" + fmt.Sprint(id)
	return query
}

func isEmpty(check string) bool {
	return check == "0" || check == "" || check == fmt.Sprint(time.Time{})
}

type Response struct {
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}, errors interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	jsonResponse := Response{
		Meta:   meta,
		Data:   data,
		Errors: errors,
	}
	return jsonResponse
}
