package api

import (
	"Solvee-User-Management/api/controllers"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var server = controllers.Server{}

// koneksi ke database, dan menjalankan aplikasi di lokal
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	}else{
		fmt.Println("We are getting the env values")
	}

	//nilai DB_DRIVER dll, didapat dari file .env
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	//seed.Load(server.DB)

	//port yang akan anda gunakan
	server.Run(":8080")
}