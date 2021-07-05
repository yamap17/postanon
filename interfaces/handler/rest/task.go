package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"kanban/application/usecase"
)

type TaskHandler interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type taskHandler struct {
	taskUseCase usecase.TaskUseCase
}

// NewTaskUseCase : Task データに関する Handler を生成
func NewTaskHandler(tu usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: tu,
	}
}

// TaskIndex : GET /tasks -> Task データ一覧を返す
func (th taskHandler) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	type request struct {
		Begin uint `query:"begin"`
		Limit uint `query:"limit"`
	}

	// taskField : response 内で使用する Book を表す構造体
	type taskField struct {
		Id        int64     `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
	}

	// response : タスク API のレスポンス
	type response struct {
		Tasks []taskField `json:"tasks"`
	}

	ctx := r.Context()

	// ユースケースの呼出
	tasks, err := th.taskUseCase.GetAll(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルを response に変換
	res := new(response)
	for _, task := range tasks {
		var tf taskField
		tf = taskField(*task)
		res.Tasks = append(res.Tasks, tf)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}