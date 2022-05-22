package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	conn     *sql.DB
	host     = os.Getenv("APP_DB_HOST")
	userName = os.Getenv("APP_DB_USER")
	password = os.Getenv("APP_DB_PASSWORD")
	port     = os.Getenv("APP_DB_PORT")
)

func init() {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/telegram_bot", userName,
			password, host, port))
	if err != nil {
		panic(err)
	}
	conn = db
}

func TestBotRepositoryImp_GetSubCategories(t *testing.T) {
	repository := NewBotRepository(conn)
	subCategory, err := repository.GetSubCategories("category_id = ?", 1)
	assert.Nil(t, err)
	assert.NotNil(t, subCategory)
	fmt.Println(subCategory)
}

func TestBotRepositoryImp_GetSubCategory(t *testing.T) {
	repository := NewBotRepository(conn)
	subCategory, err := repository.GetSubCategory("name = ?", "create_ssh")
	assert.Nil(t, err)
	assert.NotNil(t, subCategory)
	fmt.Println(subCategory)
}

func TestBotRepositoryImp_GetAllCategories(t *testing.T) {
	repository := NewBotRepository(conn)
	categories, err := repository.GetAllCategories()
	assert.Nil(t, err)
	assert.NotNil(t, categories)
	fmt.Println(categories)
}
