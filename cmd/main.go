package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./web/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	r.Get("/", handleIndex)
	r.Post("/api/system-info", handleSystemInfo)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, nil)
}

type SystemInfoData struct {
	Browser  string
	OS       string
	IP       string
	Language string
}

func handleSystemInfo(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("User-Agent")

	browser := detectBrowser(userAgent)
	os := detectOS(userAgent)

	data := SystemInfoData{
		Browser:  browser,
		OS:       os,
		IP:       getIP(r),
		Language: r.Header.Get("Accept-Language"),
	}

	tmpl, err := template.ParseFiles("web/template/system-info.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, data)
}

func detectBrowser(ua string) string {
	ua = strings.ToLower(ua)

	if strings.Contains(ua, "edg") {
		return "Microsoft Edge"
	} else if strings.Contains(ua, "chrome") {
		return "Google Chrome"
	} else if strings.Contains(ua, "firefox") {
		return "Mozilla Firefox"
	} else if strings.Contains(ua, "safari") && !strings.Contains(ua, "chrome") {
		return "Safari"
	} else if strings.Contains(ua, "opera") || strings.Contains(ua, "opr") {
		return "Opera"
	}

	return "Unknown Browser"
}

func detectOS(ua string) string {
	ua = strings.ToLower(ua)

	if strings.Contains(ua, "windows nt 10.0") {
		return "Windows 10/11"
	} else if strings.Contains(ua, "windows nt") {
		return "Windows"
	} else if strings.Contains(ua, "mac os x") {
		return "macOS"
	} else if strings.Contains(ua, "linux") {
		return "Linux"
	} else if strings.Contains(ua, "android") {
		return "Android"
	} else if strings.Contains(ua, "iphone") || strings.Contains(ua, "ipad") {
		return "iOS"
	}

	return "Unknown OS"
}

func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	}

	realIP := r.Header.Get("X-Real-IP")
	if realIP != "" {
		return realIP
	}

	return strings.Split(r.RemoteAddr, ":")[0]
}
