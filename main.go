// main.go
package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// ToDo構造体の定義
type ToDo struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

// ToDoリストの初期化
var todoList []ToDo

func main() {
	// ルーターの初期化
	r := mux.NewRouter()

	// ルートハンドラの設定
	r.HandleFunc("/api/hello", helloHandler).Methods("GET")
	r.HandleFunc("/api/todos", createToDoHandler).Methods("POST")

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

func createToDoHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストボディからToDoを読み取る
	var todo ToDo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ToDoにIDを追加してリストに追加
	todo.ID = generateID()
	todoList = append(todoList, todo)

	// レスポンスの書き込み
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// generateIDは簡易的なID生成関数
func generateID() string {
	return time.Now().Format("20060102150405")
}
