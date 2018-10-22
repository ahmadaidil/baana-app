package main

import (
	"fmt"
	"github.com/ahmadaidil/baana-app/server"
	"github.com/ahmadaidil/baana-app/migrations"
	"github.com/ahmadaidil/baana-app/route"
	bsvc "gitlab.com/ajithnn/baana/service"
	"os"
	"io/ioutil"
)

func main() {

	if len(os.Args) < 2 {
	fmt.Println("Usage: baana-app <Mode:('migrate'|'server')>; baana-app migrate <mode>:<version>; baana-app server")
	fmt.Println("E.g.: baana-app migrate up [to migrate all remaining]")
	fmt.Println("E.g.: baana-app migrate up:20180909180909 [to migrate one version]")
	fmt.Println("E.g.: baana-app server")
		os.Exit(1)
	}

	route.Init()


  	dbConf, err := ioutil.ReadFile("config/db.json")
  	if err != nil {
  		panic("Invalid db config / no db config found.")
    		os.Exit(1)
	}

  	svc, err := bsvc.New(dbConf)
  	if err != nil {
    		panic(err)
    		os.Exit(1)
  	}


	switch os.Args[1] {
	case "migrate":
		fmt.Println("Migrating...")
		migrations.Migrate(svc.DB, os.Args[2])
	case "server":
		server.Run(svc)
	default:
		fmt.Println("Unknown Mode , exiting")
		os.Exit(1)
	}
}
