// File:		main.go
// Created by:	Hoven
// Created on:	2024-08-07
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package main

import (
	"github.com/go-puzzles/cores"
	"github.com/go-puzzles/pflags"
	"github.com/go-puzzles/pgorm"
	"github.com/go-puzzles/plog"
	"github.com/superwhys/go-ddd-demo/application/account"
	"github.com/superwhys/go-ddd-demo/domain/account/service"
	"github.com/superwhys/go-ddd-demo/server"

	httppuzzle "github.com/go-puzzles/cores/puzzles/http-puzzle"
	accountRepo "github.com/superwhys/go-ddd-demo/component/repo/account"
)

var (
	port           = pflags.Int("port", 28880, "server run port")
	mysqlConfFlags = pflags.Struct("mysql-conf", (*pgorm.MysqlConfig)(nil), "mysql config")
)

func main() {
	pflags.Parse()

	mysqlConf := new(pgorm.MysqlConfig)
	plog.PanicError(mysqlConfFlags(mysqlConf))

	db, err := mysqlConf.DialGorm()
	plog.PanicError(err)
	plog.PanicError(db.AutoMigrate(&accountRepo.User{}, &accountRepo.Account{}))

	accRepo := accountRepo.NewAccountRepo(db)
	userRepo := accountRepo.NewUserRepo(db)

	accountService := service.NewAccountService(accRepo, userRepo)
	accountApp := account.NewAccountApp(accountService)

	httpServer := server.NewHttpServer(accountApp)
	httpServer.InitRouter()

	core := cores.NewPuzzleCore(
		httppuzzle.WithCoreHttpCORS(),
		httppuzzle.WithCoreHttpPuzzle("/", httpServer),
	)

	cores.Start(core, port())
}
