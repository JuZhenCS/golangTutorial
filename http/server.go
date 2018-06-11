//name: server.go
//by: ZhenJu
//date: 2018-04-02
//这其实不是一个教程，只是自己写的工具，代码挺简单的，就放到教程里了。
//每次在windows和linux之间分享文件都特别麻烦，python的http.server是单线程，而且需要python环境。
//ftp太重，而且需要客户端。
//用go简单实现一个http的多线程文件共享工具。基于网页的，方便多了。
//127.0.0.1:8080 上传
//127.0.0.1：8080/files 下载

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/files/", http.StripPrefix("/files", fs))
	http.ListenAndServe(":8083", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
	fmt.Fprintln(w, "upload ok!")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(tpl))
}

const tpl = `
<html>
	<head>
		<title>上传文件</title>
	</head>
	<body>
		<form enctype="multipart/form-data" action="/upload" method="post">
			<input type="file" name="uploadfile" />
			<input type="hidden" name="token" value="{...{.}...}"/>
			<input type="submit" value="upload" />
		</form>
	</body>
</html>`
