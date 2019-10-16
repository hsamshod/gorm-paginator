package main

import (
	"fmt"
	"github.com/hsamshod/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"not null;size:100;unique"`
}

type Phone struct {
	ID     uint `gorm:"primary_key"`
	Number string
	UserId uint `gorm:"user_id"`
	User   User `gorm:"foreignkey:UserId"`
}

func main() {
	db, err := gorm.Open("sqlite3", "example.db")
	if err == nil {
		db.AutoMigrate(&User{}, &Phone{})

		count := 0
		db.Model(User{}).Count(&count)

		if count == 0 {
			user := User{UserName: "biezhi"}
			db.Create(&user)

			db.Create(&Phone{Number: "2251097", UserId: user.ID})

			user = User{UserName: "rose"}
			db.Create(&user)
			db.Create(&Phone{Number: "2251098", UserId: user.ID})

			user = User{UserName: "jack"}
			db.Create(&user)
			db.Create(&Phone{Number: "2251099", UserId: user.ID})

			user = User{UserName: "lili"}
			db.Create(&user)
			db.Create(&Phone{Number: "2251100", UserId: user.ID})

			user = User{UserName: "wat"}
			db.Create(&user)

			fmt.Println("Insert OK!")
		}
	} else {
		fmt.Println(err)
		return
	}

	var phones []Phone

	page := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    1,
		Limit:   10,
		OrderBy: []string{"id desc"},
		Preload: []string{"User"},
		ShowSQL: true,
	}, &phones)

	fmt.Println("phones:", phones)
	fmt.Println(fmt.Sprintln("phones: %v", page.Records))
}
