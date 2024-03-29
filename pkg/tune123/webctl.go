package tune123

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"path"
	"strconv"
	"time"
)

type WebCtl struct {
	host     net.IP
	port     uint16
	islisten bool
}

type SrvConfig struct {
	managerSrv managerSrv
}

//Представляет адрес сервера управления программой и порт
type managerSrv struct {
	Addr string
	Port uint16
}

type Page struct {
	Title   string
	Body    template.HTML
	LnkHome string
	DateNow template.HTML
	SeqCnt  int
}

var (
	HTTPServerConfig SrvConfig = SrvConfig{} //Глобальная переменная для хранения настроек
	home_template              = template.Must(template.ParseFiles(path.Join("static", "tpl", "main.gtpl"), path.Join("static", "tpl", "index.html")))
	WaitExit         bool
	Quit             = make(chan int, 1) //канал для завершения сервера HTTP

)

/*Сервер*/
//Запускает goroutine ListenAndServe
func (w *WebCtl) StartServe() (err error) {
	//signal.Notify(Quit, os.Interrupt)
	srv := &http.Server{Addr: w.ConnString(), Handler: http.DefaultServeMux}

	// для отдачи сервером статичных файлов из папки public/static
	fs := http.FileServer(http.Dir("./static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	cssFileServer := http.StripPrefix("/static/", fs)
	http.Handle("/static/", cssFileServer)
	http.HandleFunc("/", urlhome) //Страница управления

	go func() {
		// log.Println("Starting HTTP-server ...")
		log.Fatalln("WebCtl:", srv.ListenAndServe())
	}()

	go func() {
		<-Quit
		fmt.Println("Shutting down HTTP-server...")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln("HTTP Shutdown error:", err)
		}
	}()
	w.islisten = true
	return err
}

//Обработчик запросов для home
func urlhome(w http.ResponseWriter, r *http.Request) {
	title := "Welcome to tune123"
	body := ""
	lnkhome := "http://" + HTTPServerConfig.managerSrv.Addr + ":" + strconv.Itoa(int(HTTPServerConfig.managerSrv.Port))
	//page := Page{title, template.HTML(body), lnkhome, "" }
	now := time.Now()
	datenow := now.Format("02/01/2006")
	page := Page{title, template.HTML(body), lnkhome, template.HTML(datenow), 0}

	if r.Method == "GET" {
		if err := home_template.ExecuteTemplate(w, "main", page); err != nil {
			log.Println(err.Error())
			http.Error(w, http.StatusText(500), 500)
		}
	} else {
		fmt.Println("POST")
	}

}

//Функции установки значений
func (w *WebCtl) SetHost(host net.IP) {
	w.host = host
}

func (w *WebCtl) SetPort(port uint16) {
	w.port = port
}

/**/
func (w WebCtl) ConnString() string {
	return fmt.Sprintf("%s:%d", w.host.String(), w.port)
}

func (s *SrvConfig) SetManagerSrv(addr string, port uint16) {
	s.managerSrv = managerSrv{
		Addr: addr,
		Port: port,
	}
}

func (s *SrvConfig) ManagerSrvAddr() string {
	return s.managerSrv.Addr
}

func (s *SrvConfig) ManagerSrvPort() uint16 {
	return s.managerSrv.Port
}

func (w WebCtl) Port() uint16 {
	return w.port
}
