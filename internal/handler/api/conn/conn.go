package conn

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kingzcheung/gexport/internal/core"
	"github.com/kingzcheung/gexport/sqlstruct"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
)

type DataConnRequest struct {
	Host     string `json:"host"`
	Port     int16  `json:"port"`
	Dbname   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *gorm.DB

func TableDetailCtx() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		if db == nil {
			http.Error(w, "数据库未连接", http.StatusBadRequest)
			return
		}
		data := core.ShowCreateTable(db.WithContext(r.Context()), name)

		ge := sqlstruct.New(sqlstruct.SQL)

		export, err := ge.Export(data.CreateTable)
		if err != nil {
			log.Println(err)
			http.Error(w, "解析失败", http.StatusBadRequest)
			return
		}

		//data.StructRes = string(export)

		//marshals, _ := json.Marshal(&data)

		w.Write(export)
	}
}

func ConnectCtx() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var form DataConnRequest
		var err error
		err = parseFromRequest(r, &form)
		if err != nil {
			http.Error(w, "参数错误", http.StatusBadRequest)
			return
		}
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			form.Username,
			form.Password,
			form.Host,
			form.Port,
			form.Dbname,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			http.Error(w, "数据库连接错误", http.StatusBadRequest)
			return
		}
		data := core.FetchTable(db.WithContext(r.Context()))

		marshal, _ := json.Marshal(&data)

		w.Write(marshal)
	}
}

func parseFromRequest(r *http.Request, v interface{}) error {
	all, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(all, v)
}
