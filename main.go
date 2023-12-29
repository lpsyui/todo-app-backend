// main.go
package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// ルーターの初期化
	r := mux.NewRouter()

	// ルートハンドラの設定
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")

	// CORSミドルウェアの追加
	handler := cors.Default().Handler(r)

	// サーバーの開始
	http.Handle("/", handler)
	http.ListenAndServe(":8090", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// レスポンスの書き込み
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "TODO１：勉強する〜"})
}
