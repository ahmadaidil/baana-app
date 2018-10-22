package migrations

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/ahmadaidil/baana-app/models"
	"github.com/jinzhu/gorm"
	"gitlab.com/ajithnn/baana/migrator"
)

type migrate struct{}

// Migrate .
func Migrate(db *gorm.DB, mode string) {

	m := migrate{}

	splitMode := strings.Split(mode, ":")
	version := ""
	mode = splitMode[0]

	if len(splitMode) == 2 {
		version = splitMode[1]
	}

	db.AutoMigrate(&models.Migration{})
	err := migrator.Migrate(db, version, mode, &m, &models.Migration{})
	if err != nil {
		fmt.Println("Error in migration: " + err.Error())
	}
}

func updateMigrations(dbConn *gorm.DB, mode string) {
	fpcs := make([]uintptr, 1)
	_ = runtime.Callers(2, fpcs)
	fun := runtime.FuncForPC(fpcs[0] - 1)
	ver := strings.Split(fun.Name(), "_")
	switch mode {
	case "up":
		dbConn.Create(&models.Migration{Version: ver[len(ver)-1]})
	case "down":
		dbConn.Delete(&models.Migration{Version: ver[len(ver)-1]})
	}
}
