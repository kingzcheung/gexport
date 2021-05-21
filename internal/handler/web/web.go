package web

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kingzcheung/gexport/internal/static"
	"github.com/sirupsen/logrus"
	"net/http"
	"path"
	"strings"
)

var contentTypes = map[string]string{
	".html": "text/html; charset=UTF-8",
	".js":   "text/javascript; charset=UTF-8",
	".css":  "text/css; charset=utf-8",
	".png":  "image/png",
	".jpg":  "image/jpeg",
	".jpeg": "image/jpeg",
}

type Server struct {
	Session *scs.SessionManager
}

func (s *Server) Handle() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.NotFound(handlerIndex())
	r.Get(`/robots.txt`, robots())
	return r
}

func robots() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("User-agent: * \n"))
		_, _ = w.Write([]byte("Disallow: /"))
	}
}

func handlerIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 开发的反向代理
		//if false {
		//	frontEndUrl := "http://127.0.0.1:3000"
		//	remote, _ := url.Parse(frontEndUrl)
		//	proxy := httputil.NewSingleHostReverseProxy(remote)
		//	proxy.ServeHTTP(w, r)
		//	return
		//}

		//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		uri := r.URL.Path
		if uri == "/" {
			uri = "index.html"
		}

		//filePath := path.Join("", uri)
		filePath := path.Join("dist", uri)

		file, err := static.WebDict.ReadFile(filePath)
		if err != nil {
			logrus.WithField("filePath", filePath).Warn(err.Error())
			if strings.HasPrefix(uri, "/api") {
				http.Error(w, "404 not found", http.StatusNotFound)
				return
			}
			//前端的引入前面
			filePath = path.Join("dist", "index.html")
			file, _ = static.WebDict.ReadFile(filePath)
		}
		ext := path.Ext(filePath)
		var contentType, ok = contentTypes[ext]
		if !ok {
			contentType = "text/plain"
		}

		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(file)
	}
}

func (s *Server) Pattern() string {
	return "/"
}
