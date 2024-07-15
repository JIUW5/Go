package main

import (
	"Basic/Utils/seq"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

/*
详细解释：以下代码
1. 通过http.HandleFunc注册了四个处理函数，分别是Index、HelloHandler、UploadHandler、DownloadHandler
2. Index函数是一个处理函数，用于处理根路径的请求，返回一个index.html文件
3. HelloHandler函数是一个处理函数，用于处理/hello路径的请求，返回一个hello页面
4. UploadHandler函数是一个处理函数，用于处理/upload路径的请求，上传文件
5. DownloadHandler函数是一个处理函数，用于处理/download路径的请求，下载文件
6. FileServer函数是一个处理函数，用于处理文件服务器的请求
7. main函数中通过http.ListenAndServe启动了一个http服务，监听8000端口
*/

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/hello", HelloHandler)
	http.HandleFunc("/upload", UploadHandler)
	http.HandleFunc("/download", DownloadHandler)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Printf("启动服务失败: err%s \n", err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// FormValue：获取请求参数
	name := r.FormValue("name")
	/*
		fmt.Sprintf("<h1>hello,%s</h1>", name): 这是 Go 语言中的一个函数，用于格式化字符串。<h1>hello,%s</h1> 是一个字符串模板，其中的 %s 是一个占位符，会被后面的 name 参数的值替换。
												所以，如果 name 的值是 "John"，那么这个函数的返回值就是 <h1>hello,John</h1>。
		[]byte(fmt.Sprintf("<h1>hello,%s</h1>", name)): 这是将上面的字符串转换为字节切片。在 Go 语言中，网络数据通常以字节切片的形式进行传输。
		w.Write([]byte(fmt.Sprintf("<h1>hello,%s</h1>", name))): 这是将上面的字节切片写入到 HTTP 响应中。
																w 是一个 http.ResponseWriter 对象，它有一个 Write 方法，可以将字节切片写入到 HTTP 响应中
	*/
	w.Write([]byte(fmt.Sprintf("<h1>hello,%s</h1>", name)))
}

/* Index
ioutil.ReadFile("index.html"): 这是 Go 语言中的一个函数，用于读取文件的内容。这个函数接收一个文件路径作为参数，在这里是 "index.html"。
								这个函数会返回两个值，一个是文件的内容（以字节切片的形式），另一个是可能出现的错误。
bytes, err := ioutil.ReadFile("index.html"): 这是使用 Go 语言的 := 操作符将 ioutil.ReadFile("index.html") 的返回值赋值给 bytes 和 err。
bytes 是一个字节切片，包含了文件的内容。err 是一个错误对象，如果在读取文件时出现了错误，那么 err 就会包含错误的详细信息，否则 err 的值就会是 nil。
go中一个语句或者函数执行成功，会返回一个内容为nil的error，如果执行失败，会返回一个非nil的error
if err != nil: 这是一个 if 语句，用于检查 err 是否不等于 nil。如果 err 不等于 nil，那么就说明在读取文件时出现了错误。
				w.WriteHeader(http.StatusNotFound): 这是将 HTTP 响应的状态码设置为 404。w 是一个 http.ResponseWriter 对象，它有一个 WriteHeader 方法，可以用来设置 HTTP 响应的状态码。
log.Printf("没有找到文件"): 这是使用 Go 语言的 log 包打印一条日志，内容是 "没有找到文件"。
return: 这是一个 return 语句，用于结束当前的函数。因为在读取文件时出现了错误，所以没有必要继续执行后面的代码。
io.WriteString(w, string(bytes)): 这是将文件的内容写入到 HTTP 响应中。io.WriteString 是一个函数，接收两个参数，一个是 io.Writer 对象，另一个是字符串。
	在这里，w 是一个 http.ResponseWriter 对象，它实现了 io.Writer 接口，所以可以作为 io.WriteString 的第一个参数。string(bytes) 是将字节切片 bytes 转换为字符串
*/

func Index(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadFile("index.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Printf("没有找到文件")
		return
	}
	io.WriteString(w, string(bytes))
}

func FileServer() {
	err := http.ListenAndServe(":8089", http.FileServer(http.Dir(".")))
	if err != nil {
		log.Printf("启动服务失败: err%s \n", err.Error())
	}
}

const (
	maxUploadfile = 2 * 1024 * 1024
	uploadpath    = "upload"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//判断文件大小
	//r.ParseMultipartForm(maxUploadfile): 这是一个方法，用于解析请求中的表单数据。
	//maxUploadfile 是一个常量，表示上传文件的最大大小。
	if err := r.ParseMultipartForm(maxUploadfile); err != nil {
		//http.StatusBadRequest: 这是一个常量，表示 HTTP 的状态码 400。
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("文件超过最大限制")
		return
	}
	//r.FormFile("file"): 这是一个方法，用于获取表单中的文件。"file" 是表单中文件的字段名。
	file, _, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("无效的file")
		return
	}
	//defer: 这是 Go 语言的一个关键字，用于注册一个函数，当当前函数执行结束时，这个函数就会被调用。
	defer file.Close()

	//ioutil.ReadAll(file): 这是一个方法，用于读取文件的内容。
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("读取文件失败")
		return
	}

	//http.DetectContentType(fileBytes): 这是一个方法，用于检测文件的类型。它接收一个字节切片作为参数，返回一个字符串，表示文件的类型。
	fileType := http.DetectContentType(fileBytes)
	if fileType != "image/jpg" && fileType != "image/png" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("此文件类型是：%s，目前仅支持jpg和png\n", fileType)
		return
	}

	fileName := seq.UUID()
	//mime.ExtensionsByType(fileType): 这是一个方法，用于获取文件类型对应的文件后缀。
	fileEndings, err := mime.ExtensionsByType(fileType)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("无法获取文件后缀")
		return
	}
	//filepath.Join(uploadpath, fileName+fileEndings[0]): 这是一个方法，用于拼接文件路径。它接收多个字符串作为参数，返回一个字符串，表示拼接后的文件路径。
	newPath := filepath.Join(uploadpath, fileName+fileEndings[0])
	log.Printf("获取到的文件类型：%s，文件路径：%s\n", fileType, newPath)
	//os.Stat(uploadpath): 这是一个方法，用于获取文件的信息。它接收一个文件路径作为参数，返回一个 os.FileInfo 对象，表示文件的信息。
	//_：这是一个占位符，用于接收 os.Stat(uploadpath) 的返回值。因为我们只关心是否存在这个文件，所以不需要获取 os.FileInfo 对象。
	_, err = os.Stat(uploadpath)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(uploadpath, 0666)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("文件夹不存在，创建文件夹失败")
				return
			}
		}
	}
	//os.Create(newPath): 这是一个方法，用于创建文件。它接收一个文件路径作为参数，返回一个 os.File 对象，表示创建的文件。
	newFile, err := os.Create(newPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("创建文件失败")
		return
	}
	defer newFile.Close()

	//newFile.Write(fileBytes): 这是一个方法，用于将字节切片写入到文件中。从旧文件中读取的内容，通过这个方法写入到新文件中。
	if _, err := newFile.Write(fileBytes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("写入文件失败")
		return
	}
	w.Write([]byte("上传成功"))
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filePath := r.FormValue("filePath")
	//os.Open(filePath): 这是一个方法，用于打开文件。它接收一个文件路径作为参数，返回一个 os.File 对象，表示打开的文件。绝对路径
	file, err := os.Open(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("打开文件失败,err:%s\n", err.Error())
		return
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("打开文件失败,err:%s\n", err.Error())
		return
	}
	w.Header().Add("Content-Disposition", "attachment;filename=\""+filePath+"\"")
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Write(bytes)
}
