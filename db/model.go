package db

import (
	"encoding/json"
	"time"

	sql "github.com/FloatTech/sqlite"

	"github.com/fumiama/langwiki-dict-local/query"
	"github.com/fumiama/langwiki-dict-local/tools/helper"
)

// Lang 某种方言
type Lang struct {
	Char rune   // Char 字
	Info string // Info 读音和解释
}

// Similar 相似字
type Similar struct {
	Char rune   // Char 字
	Simi string // Simi json: ["字", "字"]
}

type DB sql.Sqlite

func init() {
	sql.DriverName = "sqlite"
}

func NewDB(file string, cachettl time.Duration) (*DB, error) {
	db := sql.Sqlite{DBPath: file}
	err := db.Open(cachettl)
	if err != nil {
		return nil, err
	}
	nillang := &Lang{}
	for i := query.LangType(0); i < query.LangType(query.TYPE_END); i++ {
		err = db.Create(i.String(), nillang)
		if err != nil {
			return nil, err
		}
	}
	return (*DB)(&db), db.Create("相似", &Similar{})
}

func (db *DB) Close() error {
	return (*sql.Sqlite)(db).Close()
}

func (db *DB) AddChar(lang query.LangType, char *Lang) error {
	return (*sql.Sqlite)(db).Insert(lang.String(), char)
}

func (db *DB) AddSimilar(char rune, simis ...string) error {
	data, err := json.Marshal(&simis)
	if err != nil {
		return err
	}
	return (*sql.Sqlite)(db).Insert("相似", &Similar{Char: char, Simi: helper.BytesToString(data)})
}
