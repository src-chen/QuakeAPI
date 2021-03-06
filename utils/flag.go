package utils

import (
	"QuakeAPI/log"
	"QuakeAPI/model"
	"flag"
	"os"
	"strings"
)

func ParseInput() model.Input {
	var userInfo bool
	var key string
	var search string
	var help bool
	var output string
	var total int
	var fofa bool
	var quake bool
	var email string
	var config bool
	result := model.Input{}
	flag.BoolVar(&config, "config", false, "Use Config Yaml File.")
	flag.StringVar(&key, "key", "", "Input Your API Key.")
	flag.StringVar(&email, "email", "", "If You Use Fofa,You Should Enter This.")
	flag.IntVar(&total, "total", 100, "Number Of Queries You Want.")
	flag.StringVar(&search, "search", "", "Input Search String.")
	flag.StringVar(&output, "output", "result.txt", "Output File.")
	flag.BoolVar(&userInfo, "userinfo", false, "Show Your User Information.")
	flag.BoolVar(&help, "help", false, "Show Help Information.")
	flag.BoolVar(&fofa, "fofa", false, "Use Fofa.")
	flag.BoolVar(&quake, "quake", false, "Use Quake.")
	flag.Parse()
	if help == true {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if config == true {
		result.Config = config
		return result
	} else {
		if strings.TrimSpace(key) == "" {
			log.Log("Need Key", log.ERROR)
			os.Exit(0)
		}
		if strings.TrimSpace(search) == "" && userInfo == false {
			log.Log("Need Search String", log.ERROR)
			os.Exit(0)
		}
		if strings.TrimSpace(email) == "" && fofa == true {
			log.Log("Fofa Need Email", log.ERROR)
			os.Exit(0)
		}
		if quake == true && fofa == true {
			log.Log("What Do You Want To Do", log.ERROR)
			os.Exit(0)
		}
	}
	result.UserInfo = userInfo
	result.Key = key
	result.Search = search
	result.Output = output
	result.Total = total
	result.Fofa = fofa
	result.Quake = quake
	result.Email = email
	return result
}
