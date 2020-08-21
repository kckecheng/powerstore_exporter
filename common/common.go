package common

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"
)

const (
	// MAXREQ Max num. of concurrent requests
	MAXREQ = 200
	// CFG default config file
	CFG = "config.yml"
	// ResUpdateInterval resource refresh interval in seconds
	ResUpdateInterval = 300
)

var (
	// DEBUG log level
	DEBUG = false
	// Logger global log object
	Logger *log.Logger
	// ReqCounter control the total num. of concurrent requests
	ReqCounter chan int
	// Config config options
	Config cfgOptions
)

func init() {
	ReqCounter = make(chan int, MAXREQ)

	debug, ok := os.LookupEnv("DEBUG")
	if ok {
		switch debug {
		case "true", "True", "TRUE", "yes", "Yes", "YES", "y", "Y":
			DEBUG = true
		default:
			DEBUG = false
		}
	}
	logInit()
}

func logInit() {
	Logger = log.New()
	Logger.SetOutput(os.Stdout)

	formater := log.TextFormatter{
		FullTimestamp: true,
	}
	Logger.SetFormatter(&formater)

	if DEBUG {
		Logger.SetReportCaller(true)
		Logger.SetLevel(log.TraceLevel)
	} else {
		// Logger.SetLevel(log.WarnLevel)
		Logger.SetLevel(log.InfoLevel)
	}
}

type cfgOptions struct {
	PowerStore struct {
		Address  string `yaml:"address"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"powerstore"`
	Exporter struct {
		Resources []string `yaml:"resources"`
		Update    bool     `yaml:"update"`
		// VolFilter struct {
		//   Match   string `yaml:"match"`
		//   Exclude string `yaml:"exclude"`
		// } `yaml:"volume_filter"`
		// FSFilter struct {
		//   Match   string `yaml:"match"`
		//   Exclude string `yaml:"exclude"`
		// } `yaml:"file_system_filter"`
		Rollup bool  `yaml:"rollup"`
		Port   int64 `yaml:"port"`
	} `yaml:"exporter"`
}

// CfgInit init config options
func CfgInit() {
	cfg := flag.String("config", CFG, fmt.Sprintf("configuration file, %s as default", CFG))
	flag.Parse()

	Logger.Infof("Use %s as config file", *cfg)

	contents, err := ioutil.ReadFile(*cfg)
	if err != nil {
		Logger.Fatalf("Fail to read config file %s: %s", *cfg, err.Error())
	}

	err = yaml.Unmarshal(contents, &Config)
	if err != nil {
		Logger.Fatalf("Fail to decode config file %s: %s", *cfg, err.Error())
	}
}
