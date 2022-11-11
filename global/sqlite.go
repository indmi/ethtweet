package global

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

const SqliteDatabaseName = "EthTweet.0.2.db"
const SqliteDatabaseDefaultDir = "databases"

var sqliteDb *SqliteDB

type SqliteDB struct {
	dbName string
	dir    string
	*gorm.DB
}

func (sdb *SqliteDB) GetDsn() string {
	return sdb.dir + "/" + sdb.dbName
}

func GetSqliteDB() *SqliteDB {
	return sqliteDb
}

func InitSqliteDatabase(dir, name string) error {
	var err error
	if name == "" {
		name = SqliteDatabaseName
	}
	sqliteDb = &SqliteDB{
		dbName: name,
		dir:    dir,
	}
	if !IsDir(sqliteDb.dir) {
		err = os.Mkdir(sqliteDb.dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	sqliteDb.DB, err = gorm.Open(sqlite.Open(sqliteDb.GetDsn()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return err
	}
	return nil
}
