# gorm-paginator

## Usage

```bash
go get github.com/hsamshod/gorm-paginator/pagination
```

```go
type User struct {
	ID       int
	UserName string `gorm:"not null;size:100;unique"`
}

var users []User
db = db.Where("id > ?", 0)

pagination.Pagging(&pagination.Param{
    DB:      db,
    Page:    1,
    Limit:   3,
    OrderBy: []string{"id desc"},
}, &users)
```