package main

import (
	"fmt"
	"github.com/hsamshod/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

type Phone struct {
	Number string
	UserId int
	User   User
}

func main() {
	db, err := gorm.Open("sqlite3", "example.db")
	if err == nil {
		db.AutoMigrate(&User{}, &Phone{})

		count := 0
		db.Model(User{}).Count(&count)

		if count == 0 {
			db.Create(User{ID: 1, UserName: "biezhi"})
			db.Create(Phone{Number: "2251097", UserId: 1})
			db.Create(User{ID: 2, UserName: "rose"})
			db.Create(Phone{Number: "2251098", UserId: 2})
			db.Create(User{ID: 3, UserName: "jack"})
			db.Create(Phone{Number: "2251099", UserId: 3})
			db.Create(User{ID: 4, UserName: "lili"})
			db.Create(Phone{Number: "2251100", UserId: 4})
			db.Create(User{ID: 5, UserName: "bob"})
			db.Create(User{ID: 6, UserName: "tom"})
			db.Create(User{ID: 7, UserName: "anny"})
			db.Create(User{ID: 8, UserName: "wat"})
			fmt.Println("Insert OK!")
		}
	} else {
		fmt.Println(err)
		return
	}

	var phones []Phone

	pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    1,
		Limit:   3,
		OrderBy: []string{"id desc"},
		Preload: []string{"User"},
		ShowSQL: true,
	}, &phones)

	fmt.Println("phones:", phones)
}
