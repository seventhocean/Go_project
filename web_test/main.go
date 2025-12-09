package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var tmpl = template.Must(template.New("welcome").Parse(`
<!doctype html>
<html lang="zh-CN">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width,initial-scale=1">
  <title>{{.Title}}</title>
  <style>
	body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial; background:#f6f8fa; color:#111; margin:0; padding:40px; }
	.card { max-width:720px; margin:40px auto; background:white; border-radius:8px; box-shadow:0 6px 18px rgba(0,0,0,.08); padding:28px; }
	h1 { margin:0 0 6px 0; font-size:28px; }
	p.meta { color:#6b7280; margin:6px 0 18px 0; }
	.info { display:flex; gap:12px; flex-wrap:wrap; }
	.info div { background:#f3f4f6; padding:10px 12px; border-radius:6px; font-size:14px; }
	footer { margin-top:18px; color:#9ca3af; font-size:13px; }
  </style>
</head>
<body>
  <div class="card">
	<h1>{{.Title}}</h1>
	<p class="meta">一个简单的 Go 在线欢迎网站</p>
	<div class="info">
	  <div>主机: {{.Host}}</div>
	  <div>时间: {{.Now}}</div>
	  <div>Go 版本: {{.GoVersion}}</div>
	  <div>请求计数: {{.Count}}</div>
	</div>
	<footer>访问 <a href="/api" target="_blank">/api</a> 获取 JSON 信息。</footer>
  </div>
</body>
</html>
`))

type PageData struct {
	Title     string
	Host      string
	Now       string
	GoVersion string
	Count     int64
}

func main() {
	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	var count int64

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		host, _ := os.Hostname()
		data := PageData{
			Title:     "欢迎使用 Go 在线网站",
			Host:      host,
			Now:       time.Now().Format("2006-01-02 15:04:05"),
			GoVersion: runtime.Version(),
			Count:     count,
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "模板渲染错误", http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		count++
		host, _ := os.Hostname()
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json := []byte(`{"title":"欢迎使用 Go 在线网站","host":"` + host + `","time":"` + time.Now().Format(time.RFC3339) + `","go":"` + runtime.Version() + `","count":` + itoa(count) + `}`)
		w.Write(json)
	})

	log.Printf("服务器启动: http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

// minimal itoa to avoid extra imports
func itoa(i int64) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var b [32]byte
	pos := len(b)
	for i > 0 {
		pos--
		b[pos] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		pos--
		b[pos] = '-'
	}
	return string(b[pos:])
}
