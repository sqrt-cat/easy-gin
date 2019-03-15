package models

import (
	"database/sql"
	"easy-gin/drivers"
	"errors"
	"log"
	"reflect"
	"strings"
)

// models package's db obj
// all db operation should be done in models pkg
// so db is a pkg inner var
var db *sql.DB = drivers.MysqlDb

// 基础包
type Model struct {
	tableName  string
	primaryKey string
	columns    map[string]string
}

func (model *Model) reflectDbMapping() {
	modelKeys := reflect.TypeOf(model).Elem()
	// 解析模型字段
	for i := 0; i < modelKeys.NumField(); i++ {
		keyName := modelKeys.Field(i).Name
		// 没有定义/或定义为空 则columnName则为空字符串
		columnName, _ := modelKeys.Field(i).Tag.Lookup("column")

		// 如果没有定义 column 标签 则默认结构体字段名为数据库字段名
		if "" == columnName {
			columnName = keyName
		}

		model.columns[keyName] = columnName

		// 是否为主键
		_, isPK := modelKeys.Field(i).Tag.Lookup("primaryKey")

		if isPK {
			// 主键字段
			model.primaryKey = columnName
		}
	}
}

func (model *Model) getValueByKeyName(keyName string) (value interface{}) {
	modelValues := reflect.ValueOf(model)
	return modelValues.FieldByName(keyName).Interface()
}

// 获取模型主键
func (model Model) getPrimaryKeyName() string {
	return model.primaryKey
}

func (model *Model) save() (insertId int64, err error) {
	primaryKeyName := model.getPrimaryKeyName()

	if "" == primaryKeyName {
		err = errors.New("model should define primary key")
		return
	}

	primaryKeyVal := model.getValueByKeyName(primaryKeyName)

	// 将modelObj的 colName & colValue 位序映射好
	valPlaceholderSlice := make([]string, 0)
	sqlValMapping := make([]interface{}, 0)

	sqlValMapping = append(sqlValMapping, model.tableName)

	for colKeyName, colName := range model.columns {
		if colName != model.primaryKey {
			valPlaceholderSlice = append(valPlaceholderSlice, "?=?")
			sqlValMapping = append(sqlValMapping, colName, model.getValueByKeyName(colKeyName))
		}

	}

	sqlValMapping = append(sqlValMapping, model.primaryKey, primaryKeyVal)

	statPlaceholderStr := strings.Join(valPlaceholderSlice, ",")
	result, err := db.Exec("UPDATE ? SET "+statPlaceholderStr+" WHERE ?=?", sqlValMapping...)

	if nil != err {
		log.Println(err.Error())
		return
	}

	insertId, err = result.LastInsertId()
	return
}
