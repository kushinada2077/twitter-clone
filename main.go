package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"twitter-clone/handlers"
	"twitter-clone/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db := InitDB()
	fmt.Println("성공!")

	mux := http.NewServeMux()
	mux.HandleFunc("/signup", handlers.SignupHandler(db))
	mux.HandleFunc("/login", handlers.LoginHandler(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("서버가 %s 포트에서 시작되었습니다!\n", port)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}

func InitDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
			Colorful: true,
		},
	)

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("DB 연결 실패! dsn을 확인하세요")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("테이블 생성 실패")
	}

	return db
}
