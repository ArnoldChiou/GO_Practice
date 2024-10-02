package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析 url 傳遞的參數，對於 POST 則解析 HTTP 回應內容的主體（request body）
	//注意 : 如果沒有呼叫 ParseForm 方法，下面無法取得表單的資料
	fmt.Println(r.Form) //這些資訊是輸出到伺服器端的列印資訊
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //這個寫入到 w 的是輸出到客戶端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		print(token)
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		//請求的是登入資料，那麼執行登入的邏輯判斷
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			print("yessss")
		} else {
			print("noppp")
		}
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //輸出到伺服器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //輸出到客戶端
		if len(r.Form["username"][0]) == 0 {
			fmt.Println("wrong username")
		}
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("username")); !m {
			fmt.Println("no")
		} else {
			fmt.Println("yes")
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //取得請求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此處假設當前目錄下已存在 test 目錄
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)       //設定存取的路由
	http.HandleFunc("/login", login)         //設定存取的路由
	http.HandleFunc("/upload", upload)       //設定上傳的路由
	err := http.ListenAndServe(":9090", nil) //設定監聽的埠
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
