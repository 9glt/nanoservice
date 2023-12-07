package records

import (
	"flag"
	"io"
	"strings"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func SetGORMDB(g *gorm.DB) {
	db = g
}

func DBCSqlite(dsn string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Migrator().AutoMigrate(&Record{})
}

func DB() *gorm.DB {
	return db
}

func Create(flags *Flags) (record Record, err error) {
	record.UUID = uuid.NewString()
	record.Category = flags.Category
	record.Body = strings.Join(flags.Args, " ")

	if flag.Arg(0) == "-" && flags.Rdr != nil {
		var b []byte
		b, err = io.ReadAll(flags.Rdr)
		if err != nil {
			return record, err
		}
		record.Body = string(b)
	}

	record.Domain = flags.Domain
	err = db.Create(&record).Error
	if CheckError(err) != nil {
		return record, err
	}
	return
}

func List(flags *Flags) (recordz []Record, err error) {
	err = db.Where("category = ? and domain = ?", flags.Category, flags.Domain).Find(&recordz).Error
	if CheckError(err) != nil {
		return recordz, err
	}
	return
}

func Read(flags *Flags) (record Record, err error) {
	err = db.Where("category = ? and uuid = ? and domain = ?", flags.Category, flags.Arg(0), flags.Domain).First(&record).Error
	if CheckError(err) != nil {
		return record, err
	}
	return
}

func Update(flags *Flags) (record Record, err error) {
	err = db.Where("category = ? and uuid = ? and domain = ?", flags.Category, flags.Arg(0), flags.Domain).First(&record).Error
	if CheckError(err) != nil {
		return record, err
	}

	record.Body = strings.Join(flags.Args[2:], " ")
	if flag.Arg(1) == "-" && flags.Rdr != nil {
		var b []byte
		b, err = io.ReadAll(flags.Rdr)
		if err != nil {
			return record, err
		}
		record.Body = string(b)
	}
	err = db.Save(&record).Error
	if CheckError(err) != nil {
		return record, err
	}
	return
}

func Delete(flags *Flags) (record Record, err error) {
	err = db.Where("category = ? and uuid = ? and domain = ?", flags.Category, flag.Args()[0], flags.Domain).First(&record).Error
	if CheckError(err) != nil {
		return
	}

	err = db.Delete(&record).Error
	if CheckError(err) != nil {
		return
	}
	return
}
