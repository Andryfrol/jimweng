package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	_ "github.com/goPractice/crawler/crawler/plugins/inputs/all"

	// "github.com/goPractice/crawler/crawler/plugins/inputs/crawler"
	_ "github.com/goPractice/crawler/crawler/plugins/outputs/all"
	// "github.com/goPractice/crawler/crawler/plugins/outputs/mysql"
)

type CrawlerConfig struct {
	Interval int
	// Inputs   map[string]crawler.QueryUrl
	Inputs map[string]interface{}
	// Outputs  map[string]mysql.SQLConfig
	Outputs map[string]interface{}
}

var helpDoc = `
go run main.go -act [commadn]

commands:
  start : Start process.
  new   : New an configuration for process.
`

var envfile = `
# This is the default crawler configuration

[inputs]
  [inputs.crawler]
    url = "https://golang.google.cn/pkg/"

[outputs]
  [outputs.mysql]
    dbname = "pkg_lists"
    dbport = "3306"
    dbaddr = "172.18.0.3"
    user = "jim"
    password = "password"
    dbtype = "mysql"
    maxidelconns = 10
    maxopenconns = 0
	keepalive = -1
`

func NewConfig(conf string) *CrawlerConfig {
	switch conf {
	case "start":
		if _, err := os.Stat(".env.conf"); os.IsNotExist(err) {
			log.Fatal("Invalid configuration is used")
			return nil
		}
		c, err := ReadConfig(".env.conf")
		if err != nil {
			log.Fatal("Error happened while read configuration: %v", err)
			return nil
		}
		for i, j := range c.Inputs {
			fmt.Printf("inputs %v___%v\n", i, j)
		}
		for i, j := range c.Outputs {
			fmt.Printf("outputs %v___%v\n", i, j)
		}
		return c
	case "new":
		if err := genConfig(); err != nil {
			log.Fatal(err)
		}
		log.Fatal("New a configuration for crawler.")
		return nil
	default:
		log.Fatal(helpDoc)
	}
	return nil

}

func ReadConfig(filename string) (*CrawlerConfig, error) {
	var cfg CrawlerConfig
	if _, err := toml.DecodeFile(filename, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func genConfig() error {
	c := []byte(envfile)
	err := ioutil.WriteFile(".env.conf", c, 0644)
	if err != nil {
		return err
	}
	return nil
}
