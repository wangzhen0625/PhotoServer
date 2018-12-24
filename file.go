package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// curl -F file=@1.jpeg "http://10.58.122.238:8888/wz/?collection=staff" //单个文件上传
// sample usage
// func main() {
// 	target_url := "http://10.58.122.238:8888/wz/4.jpg?collection=staff"
// 	filename := "http://10.58.122.238:8888/wz/1.jpg"
// 	postFile2(filename, target_url)
// }

// 读取远程文件
func getFile(targetUrl string) {
	res, err := http.Get(targetUrl)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("wz.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, res.Body)
}

// 上传图片
func postFile(targetUrl string, base64Content string) error {
	// targetUrl := Config.Services.Web.Imgserveraddr + filepath
	decodeBytes, err := base64.StdEncoding.DecodeString(base64Content)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bytes.NewReader(decodeBytes)

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	_, filename := filepath.Split(targetUrl)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		SysLogger.Info("error writing to buffer:" + err.Error())
		return err
	}
	//iocopy
	_, err = io.Copy(fileWriter, reader)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)
	// resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	SysLogger.Info(resp.Status)
	// SysLogger.Info(string(resp_body))
	return nil
}

func postFile2(filename string, targetUrl string) error {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
