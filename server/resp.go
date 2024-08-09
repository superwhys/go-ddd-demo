// File:		resp.go
// Created by:	Hoven
// Created on:	2024-08-08
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package server

import (
	"net/http"
)

type HttpResp struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
	Err  error  `json:"-"`
}

func (r *HttpResp) GetCode() int {
	return r.Code
}

func (r *HttpResp) GetMessage() string {
	return r.Msg
}

func (r *HttpResp) GetData() any {
	return r.Data
}

func (r *HttpResp) GetError() error {
	return r.Err
}

func SuccessResponse(data any) *HttpResp {
	return &HttpResp{
		Code: http.StatusOK,
		Data: data,
	}
}

func ErrorResponse(code int, message string, err error) *HttpResp {
	return &HttpResp{
		Code: code,
		Msg:  message,
		Err:  err,
	}
}
