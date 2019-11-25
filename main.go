package main

import (
	"fmt"
	"force-to-rex/conf"
	"force-to-rex/rex"
	"io/ioutil"
	"os"
	"time"

	"github.com/dutchcoders/goftp"
)

func main() {
	if !conf.Load() {
		return
	}
	currentTime := time.Now()
	fileDateId := fmt.Sprintf("%s_%s_%s", conf.Config.Str("integratorID"),
		fmt.Sprint(currentTime.Format("02012006")), fmt.Sprint(currentTime.Format("150405")))
	rex.GenerateFiles(fileDateId)
	//UploadOnFTP(fileDateId)
	// TODO remove "temp" folder
}

func UploadOnFTP(fileDateId string) {
	var ftp *goftp.FTP
	var err error
	if ftp, err = goftp.Connect(conf.Config.Str("fptAddr")); err != nil {
		fmt.Println(err)
		return
	}
	defer ftp.Close()

	if err = ftp.Login(conf.Config.Str("userName"), conf.Config.Str("password")); err != nil {
		fmt.Println(err)
		return
	}

	if err = ftp.Cwd("/XMLData/"); err != nil {
		fmt.Println(err)
		return
	}

	var file *os.File
	if file, err = os.Open("temp/P_" + fileDateId + ".xml"); err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = ftp.Stor("P_"+fileDateId+".xml", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()

	if err = ftp.Cwd("/XMLImage/"); err != nil {
		fmt.Println(err)
		return
	}

	var files []os.FileInfo
	files, err = ioutil.ReadDir("temp/images")
	for _, fileInfo := range files {
		if file, err = os.Open("temp/images/" + fileInfo.Name()); err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		if err = ftp.Stor(fileInfo.Name(), file); err != nil {
			fmt.Println(err)
			return
		}
		file.Close()
	}

	if err = ftp.Cwd("/XMLData/"); err != nil {
		fmt.Println(err)
		return
	}

	if file, err = os.Open("temp/JOB_" + fileDateId + ".xml"); err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	if err = ftp.Stor("JOB_"+fileDateId+".xml", file); err != nil {
		fmt.Println(err)
	}
}
