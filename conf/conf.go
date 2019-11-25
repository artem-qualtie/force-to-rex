package conf

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	FileConfigPath = "./app.conf"
	Config         = make(Conf)
)

func Load() bool {
	_, err := ioutil.ReadFile(FileConfigPath)
	if err != nil {
		err = ioutil.WriteFile(FileConfigPath, []byte(DefaultConfig), 0666)
		if err != nil {
			panic("Error save Default Config")
		} else {
			fmt.Println("Default Config saved. Edit it and start the server again!")
		}

		return false
	}

	if err = loadConfig(); err != nil {
		fmt.Println("Error parsing config: " + err.Error())
		return false
	}

	return true
}

func loadConfig() (err error) {
	file, err := os.Open(FileConfigPath)
	if err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				Config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}

	return
}
