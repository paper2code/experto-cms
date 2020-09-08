package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/admin"
)

type About struct {
	gorm.Model
	Name string
	Description
}

type Bookmark struct {
	gorm.Model
	URL         string
	Description string
}

type ReadingList struct {
	gorm.Model
	URL         string
	Name        string
	Description string
}

type Watcher struct {
	gorm.Model
	URL         string
	Name        string
	Description string
}

type Tag struct {
	gorm.Model
	Name string
}

func main() {

	DB, _ := gorm.Open("sqlite3", "experto-cms.db")
	DB.AutoMigrate(&About{}, &Bookmark{}, &ReadingList{}, &Watcher{}, &Tag{})

	// Initialize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Allow to use Admin to manage User, Product
	Admin.AddResource(&About{})
	Admin.AddResource(&Bookmark{})
	Admin.AddResource(&ReadingList{})
	Admin.AddResource(&Watcher{})
	Admin.AddResource(&Tag{})

	// initalize an HTTP request multiplexer
	mux := http.NewServeMux()

	// Mount admin interface to mux
	Admin.MountTo("/admin", mux)

	fmt.Println("Listening on: 9000")
	http.ListenAndServe(":9000", mux)

}
