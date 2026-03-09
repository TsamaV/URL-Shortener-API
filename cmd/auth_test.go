package main

import (
	"Documents/Web_GO/internal/auth"
	"Documents/Web_GO/internal/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func InitData(db *gorm.DB) {
	db.Create(&user.User{
		Email: "a2@a.ru",
		Password: "$2a$10$qvexQwyk0x1eWeppXT2lYeeYJMk4Kz58c3DX9xI6VD72x.Ps.Ohcy",
		Name: "Lexa",
	})
}

func RemoveData(db *gorm.DB) {
	db.Unscoped().Where("email = ?", "a2@a.ru").Delete(&user.User{})
}
func TestLoginSuccess(t *testing.T) {
	// Prepare
	db := InitDb()
	InitData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "12233",
	})

	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Logf("Response body: %s", body)
	if resp.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, resp.StatusCode)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token empty")
	}
	RemoveData(db)
}

func TestLohinFail(t *testing.T) {
	db := InitDb()
	InitData(db)

	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "a2@a.ru",
		Password: "122323",
	})
	resp, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	t.Logf("Response body: %s", body)
	if resp.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, resp.StatusCode)
	}
	RemoveData(db)
}
