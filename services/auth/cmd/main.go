package main

import (
	"github.com/imdinnesh/prepaid-card/services/auth/config"
	"github.com/imdinnesh/prepaid-card/services/auth/db"
	"github.com/imdinnesh/prepaid-card/services/auth/router"
)

func main(){
	cfg:=config.Load();
	db:=db.InitDB(cfg)
	Router:=router.New(cfg,db)
	Router.Run(":"+cfg.ServerPort)
}