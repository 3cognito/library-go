package initializers

import (
	"fmt"

	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/modules/users"
	"github.com/3cognito/library/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbConfigs := config.Configs.DB

	sslmode := "required"
	if config.Configs.ENV == config.Dev {
		sslmode = "disable"
	}

	dbconnStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		dbConfigs.Host,
		dbConfigs.Port,
		dbConfigs.User,
		dbConfigs.Password,
		dbConfigs.Name,
		sslmode,
	)

	db, err := gorm.Open(postgres.Open(dbconnStr), &gorm.Config{
		NowFunc: utils.TimeNow,
	})
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		panic(err)
	}

	err = db.AutoMigrate(
		&users.User{}, &books.Book{},
	)

	if err != nil {
		fmt.Println("Error migrating database: ", err)
		panic(err)
	}
	fmt.Println("Database connected")

	DB = db
}
