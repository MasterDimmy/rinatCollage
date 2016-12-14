package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	fmt.Println(r.URL.String())
	switch r.URL.String() {
	case "/":
		http.Redirect(w, r, "/static/index.html", http.StatusMovedPermanently)
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache")
	fmt.Println(r.URL.String())
	w.Write([]byte("your data"))
}

func (a *TApp) createWebServer() error {
	if len(a.Cfg.WebServer.Static) == 0 {
		return fmt.Errorf("Ошибка! Пусть путь к папке Веб-сервера static (задается в настройках)")
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/get_data", getData)
	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir(app.Cfg.WebServer.Static)))
	http.Handle("/static/", fileServer)

	fmt.Println("\nЗапуск Веб-сервера по адресу http://" + a.Cfg.WebServer.IpPort + "\nКаталог Веб-сервера: " + app.Cfg.WebServer.Static)

	err := http.ListenAndServe(a.Cfg.WebServer.IpPort, nil)
	return err
}
