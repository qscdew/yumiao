package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"yumiao/conf"
)


func main() {

	e, err := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv")
	fmt.Println(e,err)
	conf.RunServer()


}

