// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orm

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// oracle operators.
var oracleOperators = map[string]string{
	"exact":       "= ?",
	"gt":          "> ?",
	"gte":         ">= ?",
	"lt":          "< ?",
	"lte":         "<= ?",
	"//iendswith": "LIKE ?",
}

// oracle column field types.
var oracleTypes = map[string]string{
	"pk":              "NOT NULL PRIMARY KEY",
	"bool":            "bool",
	"string":          "VARCHAR2(%d)",
	"string-char":     "CHAR(%d)",
	"string-text":     "VARCHAR2(%d)",
	"time.Time-date":  "DATE",
	"time.Time":       "TIMESTAMP",
	"int8":            "INTEGER",
	"int16":           "INTEGER",
	"int32":           "INTEGER",
	"int64":           "INTEGER",
	"uint8":           "INTEGER",
	"uint16":          "INTEGER",
	"uint32":          "INTEGER",
	"uint64":          "INTEGER",
	"float64":         "NUMBER",
	"float64-decimal": "NUMBER(%d, %d)",
}

// oracle dbBaser
type dbBaseOracle struct {
	dbBase
}

var _ dbBaser = new(dbBaseOracle)

// create oracle dbBaser.
func newdbBaseOracle() dbBaser {
	b := new(dbBaseOracle)
	b.ins = b
	return b
}

// OperatorSQL get oracle operator.
func (d *dbBaseOracle) OperatorSQL(operator string) string {
	return oracleOperators[operator]
}

// DbTypes get oracle table field types.
func (d *dbBaseOracle) DbTypes() map[string]string {
	return oracleTypes
}

//ShowTablesQuery show all the tables in database
func (d *dbBaseOracle) ShowTablesQuery() string {
	return "SELECT TABLE_NAME FROM USER_TABLES"
}

// Oracle
func (d *dbBaseOracle) ShowColumnsQuery(table string) string {
	return fmt.Sprintf("SELECT COLUMN_NAME FROM ALL_TAB_COLUMNS "+
		"WHERE TABLE_NAME ='%s'", strings.ToUpper(table))
}

// check index is exist
func (d *dbBaseOracle) IndexExists(db dbQuerier, table string, name string) bool {
	row := db.QueryRow("SELECT COUNT(*) FROM USER_IND_COLUMNS, USER_INDEXES "+
		"WHERE USER_IND_COLUMNS.INDEX_NAME = USER_INDEXES.INDEX_NAME "+
		"AND  USER_IND_COLUMNS.TABLE_NAME = ? AND USER_IND_COLUMNS.INDEX_NAME = ?", strings.ToUpper(table), strings.ToUpper(name))

	var cnt int
	row.Scan(&cnt)
	return cnt > 0
}

// execute insert sql with given struct and given values.
// insert the given values, not the field values in struct.
func (d *dbBaseOracle) InsertValue(q dbQuerier, mi *modelInfo, isMulti bool, names []string, values []interface{}) (int64, error) {
	Q := d.ins.TableQuote()

	marks := make([]string, len(names))
	for i := range marks {
		marks[i] = ":" + names[i]
	}

	sep := fmt.Sprintf("%s, %s", Q, Q)
	qmarks := strings.Join(marks, ", ")
	columns := strings.Join(names, sep)

	multi := len(values) / len(names)

	if isMulti {
		qmarks = strings.Repeat(qmarks+"), (", multi-1) + qmarks
	}

	query := fmt.Sprintf("INSERT INTO %s%s%s (%s%s%s) VALUES (%s)", Q, mi.table, Q, Q, columns, Q, qmarks)

	d.ins.ReplaceMarks(&query)

	if isMulti || !d.ins.HasReturningID(mi, &query) {
		res, err := q.Exec(query, values...)
		if err == nil {
			if isMulti {
				return res.RowsAffected()
			}
			return res.LastInsertId()
		}
		return 0, err
	}
	row := q.QueryRow(query, values...)
	var id int64
	err := row.Scan(&id)
	return id, err
}

// return quote.
func (d *dbBaseOracle) TableQuote() string {
	return ""
}

// replace value placeholder in parametered sql string.
func (d *dbBaseOracle) ReplaceMarks(query *string) {
	num := strings.Count(*query, "?")
	var tmp string
	tmp = *query
	for i := 0; i < num; i++ {
		tmp = strings.Replace(tmp, "?", fmt.Sprintf(":P%d", i), 1)
	}
	tmp = strings.Replace(tmp, "LIMIT", "AND ROWNUM <=", 1)
	qptrV := reflect.ValueOf(query)
	qptrV.Elem().SetString(tmp)
}

// convert time from db.
func (d *dbBaseOracle) TimeFromDB(t *time.Time, tz *time.Location) {
	tadd, _ := time.ParseDuration("-8h")
	tz, _ = time.LoadLocation("Asia/Shanghai")
	*t = t.Add(tadd).In(tz)

}

// convert time to db.
func (d *dbBaseOracle) TimeToDB(t *time.Time, tz *time.Location) {
	tadd, _ := time.ParseDuration("8h")
	*t = t.Add(tadd).In(tz)

}
