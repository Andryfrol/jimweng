package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/goPractice/crawler/crawler/plugins/inputs/crawler"
	"github.com/goPractice/crawler/crawler/plugins/outputs/mysql"
	"github.com/goPractice/crawler/crawler/utils"
)

type CrawlerConfig struct {
	Interval      int
	InputPlugins  map[string]utils.Input
	OutputPlugins map[string]utils.Output
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

# modify mysql dbaddr while start with docker-compose,
# since there is a networkspace in proxy-next
[outputs]
  [outputs.mysql]
    dbname = "mysql"
    dbport = "3306"
    dbaddr = "127.0.0.1"
    user = "root"
    password = "secret"
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
	cfg := CrawlerConfig{}
	type confcfg struct {
		Interval int
		Inputs   map[string]interface{}
		Outputs  map[string]interface{}
	}

	var ccfg = &confcfg{}

	if _, err := toml.DecodeFile(filename, &ccfg); err != nil {
		return nil, err
	}

	// TOD: make a loop for auto collect implement plugins
	cfgInputMap := buildInputMap("crawler", newInputPlugin(&ccfg.Inputs))
	// assign implement input plugins to cfg input plugin
	cfg.InputPlugins = *cfgInputMap

	cfgOutputMap := buildOutputMap("mysql", newOutPlugin(&ccfg.Outputs))
	// assign implement output plugins to cfg input plugin
	cfg.OutputPlugins = *cfgOutputMap

	return &cfg, nil
}

func newOutPlugin(m *map[string]interface{}) utils.Output {
	var outputMysql = mysql.SQLConfig{}
	for i, j := range *m {
		// fmt.Printf("implement output plugin '%v' with value %v\n", i, j)
		if i == "mysql" {
			for key, value := range j.(map[string]interface{}) {
				switch key {
				case "dbname":
					outputMysql.DBName = value.(string)
				case "dbaddr":
					outputMysql.DBAddr = value.(string)
				case "password":
					outputMysql.Password = value.(string)
				case "dbtype":
					outputMysql.DBType = value.(string)
				case "maxidelconns":
					outputMysql.MaxIdleConns = int(value.(int64))
				case "maxopenconns":
					outputMysql.MaxOpenConns = int(value.(int64))
				case "dbport":
					outputMysql.DBPort = value.(string)
				case "user":
					outputMysql.User = value.(string)
				case "keepalive":
					outputMysql.KeepAlive = int(value.(int64))
				}
			}
		}
	}
	return &outputMysql
}

func newInputPlugin(m *map[string]interface{}) utils.Input {
	var inputCrawler = crawler.QueryUrl{}
	for i, j := range *m {
		// fmt.Printf("implement input plugin '%v' with value %v\n", i, j)
		if i == "crawler" {
			for key, value := range j.(map[string]interface{}) {
				if key == "url" {
					inputCrawler.Url = value.(string)
				}
			}
		}
	}
	return &inputCrawler
}

func buildInputMap(name string, utilInput utils.Input) *map[string]utils.Input {
	m := make(map[string]utils.Input)
	m[name] = utilInput
	return &m
}

func buildOutputMap(name string, utilInput utils.Output) *map[string]utils.Output {
	m := make(map[string]utils.Output)
	m[name] = utilInput
	return &m
}

func genConfig() error {
	c := []byte(envfile)
	err := ioutil.WriteFile(".env.conf", c, 0644)
	if err != nil {
		return err
	}
	return nil
}
