package main

// 每一次请求都要处理body内容，会有一大段代码，这里定义一个content来统一处理

import (
	"encoding/json"
	"io"
	"net/http"
)

//为了简单，这里没有定义接口，直接用结构体，不会太扩展和用多种实现
type Context struct {
	W http.ResponseWriter //接口
	R *http.Request       //结构体(底层实现的时候应该也认为Reuqest比较固定)
}

// 处理body(帮我读取body,帮我反序列化)
func (c *Context) ReadJson(req interface{}) error {
	r := c.R
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err // 框架里不用打日志，将错误返回，让调用者决定怎么处理
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}

	return nil
	// 注意：最后处理完并没有返回处理的数据，只告诉有没有处理成功
	// go语言一般这样处理：接收一个结构(req)，数据处理后直接填充到这个结构
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson) //这里不考虑多次写入的情况，匿名
	return err
}

//基于核心方法(WriteJson)扩展出来的方便用户使用的方法
func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) SystemErrorJson(resp interface{}) error {
	return c.WriteJson(http.StatusInternalServerError, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}
