package config_test

import(
	"testing"

	"git.expense-app.com/ExpenseApp/config"
	"github.com/stretchr/testify/assert"
)

func TestReadConfigSuccess(t *testing.T){
	appConfig,err := config.ReadConfig("./../testfiles/test_config.json")

	assert.Equal(t, "9009", appConfig.Port)
	assert.NoError(t, err)
}

func TestReadConfigFailsForFileReadError(t *testing.T){
	appConfig,err := config.ReadConfig("noSuchFile.txt")

	assert.Error(t, err)
	assert.Nil(t, appConfig)
}