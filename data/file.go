// migrations.go
package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type File struct {
	gorm.Model
	Filename   string
	UUID       int
	Downloaded bool
}
