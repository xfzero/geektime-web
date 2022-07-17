package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	//设定一个路由，命中该路由的会执行handlerFunc的代码
	Route(pattern string, handleFunc http.HandlerFunc)

	//启动我们的服务器
	Start(address string) error
}

// sdkHttpServer 基于 http 库实现
type sdkHttpServer struct {
	Name string
}

// Route 注册路由
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	//panic("implemnent me")
	http.HandlerFunc(pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	//panic("implemnent me")
	return http.ListenAndServer(address, nil)
}

func NewHttpServer(name string) Server {
	//不是指针会报错(返回值是接口实现)
	return &sdkHttpServer{
		Name: name,
	}
}

//有多种实现时
// func NewServerBaseOnXXX() Server {
// }

//多种实现2
// type Factory func() Server
// var factory Factory
// func RegisterFactory(f Factory) {
// 	factory = f
// }
// func NewServer() Server {
// 	return factory
// }

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &SignUpReq{}

	ctx := &Context{
		W: w,
		R: r,
	}

	err := ctx.ReadJson(req)
	if err != nil {
		// fmt.Fprintf(w, "err:%v", err)
		// return
		ctx.BadRequestJson(err)
	}

	resp := &commonResponse{
		Data: 123,
	}

	//StatusOK这种返回会很常见，这里可以继续分装(jsonOk方法)
	err := ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

type SignUpReq struct {
	// Tag
	Email              string `json:"email"`
	Password           string `json:"password"`
	ConfiremedPassword string `json:"confiremed_passeord"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
