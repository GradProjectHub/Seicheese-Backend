package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"seicheese/internal/handler"
	firebase "seicheese/internal/infrastructure"
	"seicheese/internal/infrastructure/database"
	router "seicheese/internal/middleware/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Firebaseの初期化
	firebaseApp, err := firebase.InitializeFirebaseApp()
	if err != nil {
		log.Fatalf("Firebase initialization error: %v", err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("Auth client initialization error: %v", err)
	}

	// データベース接続
	dbConfig := database.NewDBConfig()
	db, err := database.InitializeDB(dbConfig)
	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}
	defer db.Close()

	// ハンドラーの初期化
	authHandler := &handler.AuthHandler{
		DB:         db,
		AuthClient: authClient,
	}

	// ルーターの登録
	router.RegisterAuthRoutes(e, authClient, authHandler)

	// サーバー起動
	port := os.Getenv("PORT")
	if port == "" {
		port = "1300"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
