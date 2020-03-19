package diplomas

import (
	"github.com/seuc-frp-utn/api/database"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	database.SetupDatabaseTests()
	os.Exit(m.Run())
}
