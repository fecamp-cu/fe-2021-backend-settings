package main

import (
	"github.com/fecamp-cu/fe-2021-backend-settings/src/databases"
	"github.com/fecamp-cu/fe-2021-backend-settings/src/mock"
)

func main() {
	db := databases.GetDB()
	db.CreateSetting(&mock.MockSettings)
}
