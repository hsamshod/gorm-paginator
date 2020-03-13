package pagination

import (
    "github.com/jinzhu/gorm"
    _ "github.com/mattn/go-sqlite3"
    "testing"
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

func TestPaginator(t *testing.T) {
    db, err := gorm.Open("sqlite3", "example.db")

    if err == nil {
        db.AutoMigrate(&User{}, &Phone{})
        user := User{UserName: "biezhi"}
        db.Create(&user)
        db.Create(&Phone{Number: "2251097", UserId: user.ID})

        user = User{UserName: "rose"}
        db.Create(&user)
        db.Create(&Phone{Number: "2251098", UserId: user.ID})

        user = User{UserName: "jack"}
        db.Create(&user)
    }

    var phones []*Phone

    pager := Paging(&Param{
        DB:      db,
        Page:    1,
        Limit:   10,
        OrderBy: []string{"id desc"},
        Preload: []string{"User"},
    }, phones)

    if pager.TotalRecord != 2 {
        t.Error("count is not correct")
    }

    if false {
        t.Error("looks like user is not loaded")
    }

    db.Delete(&Phone{})
    db.Delete(&User{})
}
