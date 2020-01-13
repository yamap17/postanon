package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	handler "rest_api/interfaces/handler/rest"
	"rest_api/infrastructure/persistence"
	"rest_api/application/usecase"
)

func main() {
	// 依存関係を注入
	// TODO DI ライブラリを使う
	taskPersistence := persistence.NewTaskPersistence()
	taskUseCase := usecase.NewTaskUseCase(taskPersistence)
	taskHandler := handler.NewTaskHandler(taskUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/tasks", taskHandler.Index)

	// サーバ起動
	fmt.Println("Server Start >> http://localhost:1323")
	log.Fatal(http.ListenAndServe(":1323", router))
}
