package main

import (
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"path"
	"strings"
	"time"
)

type Config struct {
	Quartie []Quartie   `json:"quartie"`
	Path    string      `json:"path"`
	Res     interface{} `json:"res"`
}
type Quartie struct {
	Percent int64 `json:"percent"`
	Sleep   int64 `json:"sleep"`
}

func (config Config) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//简单实现，排序依赖配置设置
	randomTotal, _ := rand.Int(rand.Reader, big.NewInt(100))
	var sleep int64
	for idx, item := range config.Quartie {
		if item.Percent > randomTotal.Int64() {
			start := int64(0)
			if idx > 0 {
				start = config.Quartie[idx-1].Sleep
			}
			diff := item.Sleep - start
			randomSleep, _ := rand.Int(rand.Reader, big.NewInt(diff))
			sleep = randomSleep.Int64() + start
			break
		}
	}
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	b, _ := json.Marshal(config.Res)
	w.Write(b)
}

var configList []Config
var port string

func init() {
	var p = flag.String("p", "8179", "string类型参数")
	var c = flag.String("c", "./config", "string类型参数")
	flag.Parse()
	port = *p
	configPath := *c
	files, _ := ioutil.ReadDir(configPath)
	for _, f := range files {
		if !strings.HasSuffix(strings.ToLower(f.Name()), ".json") {
			log.Println(f.Name() + " file be ignore，becase not has .json suffix")
		}
		if f.IsDir() {
			log.Println(f.Name() + " dir be ignore")
		}
		conf, err := ioutil.ReadFile(path.Join(configPath, f.Name()))
		if err != nil {
			panic("can not found " + path.Join(configPath, f.Name()))
		}
		var config Config
		err = json.Unmarshal(conf, &config)
		if err != nil {
			fmt.Println(err)
			panic("can not unmarshal " + path.Join(configPath, f.Name()))
		}
		configList = append(configList, config)
	}
	if len(configList) == 0 {
		panic("availability config list is empty")
	}
	for _, l := range configList {
		http.Handle(l.Path, l)
	}
}

func main() {
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
