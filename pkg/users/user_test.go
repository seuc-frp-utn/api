package users

import (
	"github.com/seuc-frp-utn/api/pkg/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database.SetupDatabaseTests()
	os.Exit(m.Run())
}
