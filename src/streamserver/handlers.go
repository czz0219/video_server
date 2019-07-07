package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	complete_fn := GetCompletePath() + "upload.html"
	t, _ := template.ParseFiles(complete_fn)
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	//	vl:=VIDEO_DIR+vid
	vl := GetCompletePath() + vid
	log.Printf("show videos:%s\n",vl)
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("open file %v:%v", vl, err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error:open video error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil { //当提交视频大于MaxBytesReader规定值,一定出错
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	//此处的FormFile对应的前端的 var formData = new FormData()
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	fn := p.ByName("vid-id")
	vl := GetCompletePath() + fn
	log.Printf("upload:%s\n",vl)
	err = ioutil.WriteFile(vl, data, 0666)
	if err != nil {
		log.Printf("wRITE FILE Error:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
