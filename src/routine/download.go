package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func download(downloadUrl string) (string, error) {
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(downloadUrl)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wait sync.WaitGroup
	var urls = []string{
		"https://upload.wikimedia.org/wikipedia/commons/d/d0/161204_%ED%8A%B8%EC%99%80%EC%9D%B4%EC%8A%A4_%EC%82%AC%EB%82%98_%EB%8B%A4%ED%98%84_ICN_MAMA_%EC%B6%9C%EA%B5%AD_by._%EC%A0%95%EC%97%B0%EC%9D%80_%EC%82%AC%EB%9E%91_%282%29.jpg",
		"https://photo.jtbc.joins.com/news/2018/04/28/20180428164615418.jpg",
		"https://img.hankyung.com/photo/201904/03.19476949.1.jpg",
	}
	
	for _, url := range urls {
		wait.Add(1)
		go func(url string){
			defer wait.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wait.Wait()
	
	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("Dahyun.zip", filenames)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Finish!")
}