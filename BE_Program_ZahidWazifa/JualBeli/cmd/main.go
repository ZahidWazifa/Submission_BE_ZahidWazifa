package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "buy-sell-backend/internal/api/routes"
)

func main() {
    // Membuat router baru
    r := mux.NewRouter()

    // Mengatur rute untuk API
    routes.SetupRoutes(r)

    // Menjalankan server pada port 8080
    log.Println("Server berjalan di http://localhost:8080")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatal(err)
    }
}