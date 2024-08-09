// File:		http.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package server

import (
	"context"
	"net/http"

	"github.com/go-puzzles/prouter"
	"github.com/superwhys/go-ddd-demo/application/account"
	"github.com/superwhys/go-ddd-demo/application/account/dto"
)

type Server struct {
	mux        *prouter.Prouter
	accountApp *account.AccountApp
}

func NewHttpServer(accountApp *account.AccountApp) *Server {
	return &Server{
		mux:        prouter.NewProuter(),
		accountApp: accountApp,
	}
}

func (s *Server) InitRouter() {
	accountGroup := s.mux.Group("/account")
	{
		accountGroup.POST("/login", prouter.BodyParserHandleFunc(s.loginHandler))
		accountGroup.POST("/register", prouter.BodyParserHandleFunc(s.registerHandler))

	}
}

func (s *Server) loginHandler(_ context.Context, req *dto.LoginRequest) (*dto.Account, error) {
	return s.accountApp.Login(req)
}

func (s *Server) registerHandler(_ context.Context, req *dto.RegisterRequest) (*dto.Account, error) {
	return s.accountApp.Register(req)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
