package api

import (
	"encoding/json"
	"fmt"
	"github.com/Aquarian-Age/sjqm"
	"html/template"
	"liangzi.local/cal"
	"log"
	"net/http"
	"strconv"
	"time"
)

//
func Main(port string) {
	http.HandleFunc("/", Home)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("listenAndServer:", err)
	}
}
func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("home.html")
		err := t.Execute(w, nil)
		if err != nil {
			return
		}
	} else {
		handhome(w, r)
	}
}
func handhome(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		return
	}
	t := getTimes(r)
	body := NewBody(t)
	b := body.BodyBytes()

	err = json.NewEncoder(w).Encode(string(b))
	if err != nil {
		return
	}
}
func getTimes(r *http.Request) time.Time {
	//年
	ly, err := strconv.Atoi(r.Form["ly"][0])
	if err != nil {
		log.Fatalln("年异常:", err)
	}
	//月
	lm, err := strconv.Atoi(r.Form["lm"][0])
	if err != nil {
		log.Fatalln("月异常: ", err)
	}

	//日
	ld, err := strconv.Atoi(r.Form["ld"][0])
	if err != nil {
		log.Fatalln("日异常:", err)
	}
	//时辰 子时1 丑时2 寅时3...
	lh, err := strconv.Atoi(r.Form["lh"][0])
	if err != nil {
		log.Fatalln("时辰异常:", err)
	}
	t := time.Date(ly, time.Month(lm), ld, lh, 0, 0, 0, time.Local)
	return t
}

//############################################################
type Body struct {
	*sjqm.G
	Gzs string `json:"gzs"` //干支
}

func NewBody(t time.Time) *Body {
	gz := cal.NewGanZhiInfo(t.Year(), int(t.Month()), t.Day(), t.Hour())
	ygz := gz.YearGZ
	mgz := gz.MonthGZ
	dgz := gz.DayGZ
	hgz := gz.HourGZ
	cust := t
	g, _ := sjqm.Result(t.Year(), dgz, hgz, cust)
	gzs := fmt.Sprintf("%s-%s-%s-%s", ygz, mgz, dgz, hgz)
	return &Body{
		g,
		gzs,
	}
}

func (body *Body) BodyBytes() []byte {
	b, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
