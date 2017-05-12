package Init

import (
	"log"
	"github.com/BoredTape/httppool"
	. "btspider/Global"
	"strconv"
	"fmt"
	"time"
)

func Init(name string) {
	conf, err := InitReader(`D:\\workspace-go\\src\\btspider\\config.ini`)
	if err != nil {
		log.Fatal(err)
	}
	size, err := strconv.Atoi(conf.Value(name, `poolsize`))
	if err != nil {
		fmt.Println(err)
		size = 0
	}
	var timeout time.Duration
	second, err := strconv.Atoi(conf.Value(name, `timeout`))
	if err != nil {
		second = 0
	}
	timeout = time.Duration(second) * time.Second
	mysql_db := conf.Value(name, `mysql_db`)
	mysql_ip := conf.Value(name, `mysql_url`)
	mysql_port := conf.Value(name, `mysql_port`)
	mysql_user := conf.Value(name, `mysql_user`)
	mysql_passwd := conf.Value(name, `mysql_passwd`)
	Spider.Pool = httppool.NewPools(&httppool.Options{Size: size, Timeout: timeout})
	if mysql_db != "" && mysql_ip != "" && mysql_port != "" && mysql_user != "" && mysql_passwd != "" {
		Spider.Err = Spider.DB.InitDB(mysql_ip, mysql_port, mysql_user, mysql_passwd, mysql_db, "utf8")
		if Spider.Err != nil {
			fmt.Println(Spider.Err)
		}
	}
}
