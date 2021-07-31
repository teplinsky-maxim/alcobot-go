package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Value struct {
	gorm.Model
	Uid      uint
	Username string
	Value    string
	Time     uint64
}

var Db *gorm.DB

func InitDatabase() (err error) {
	Db, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	if err != nil {
		return err
	}

	err = Db.AutoMigrate(Value{})
	if err != nil {
		return err
	}

	return
}

func CheckIfResultAlreadyGeneratedForToday(uid uint) (bool, string) {
	var result Value

	todayTimestamp := getTodayTS()
	tomorrowTimestamp := todayTimestamp + 60*60*24

	affectedRows := Db.Where("uid = ? AND time BETWEEN ? AND ?", uid, todayTimestamp, tomorrowTimestamp).First(&result).RowsAffected
	if affectedRows == 0 {
		return false, ""
	}

	return true, result.Value
}

func InsertResult(username, value string, uid uint) {
	Db.Create(&Value{
		Uid:      uid,
		Username: username,
		Value:    value,
		Time:     uint64(time.Now().Unix()),
	})
}

func getTodayTS() uint64 {
	currentTime := time.Now()
	return uint64(time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		0,
		0,
		0,
		0,
		currentTime.Location(),
	).Unix())
}
