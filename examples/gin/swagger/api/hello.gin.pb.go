// Code generated by protoc-gen-zmicro-gin. DO NOT EDIT.
// versions:
// - protoc-gen-zmicro-gin v0.1.0
// - protoc                v3.19.0
// source: api/hello.proto

package proto

import (
	context "context"
	errors "errors"
	gin "github.com/gin-gonic/gin"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = errors.New
var _ = context.TODO
var _ = gin.New

type GreeterHTTPServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Validate(context.Context, any) error
	ErrorEncoder(c *gin.Context, err error, isBadRequest bool)
}

type UnimplementedGreeterHTTPServer struct{}

func (*UnimplementedGreeterHTTPServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, errors.New("method SayHello not implemented")
}
func (*UnimplementedGreeterHTTPServer) Validate(context.Context, any) error { return nil }
func (*UnimplementedGreeterHTTPServer) ErrorEncoder(c *gin.Context, err error, isBadRequest bool) {
	var code = 500
	if isBadRequest {
		code = 400
	}
	c.String(code, err.Error())
}

func RegisterGreeterHTTPServer(g *gin.RouterGroup, srv GreeterHTTPServer) {
	r := g.Group("")
	r.GET("/hello/:name", _Greeter_SayHello0_HTTP_Handler(srv))
}

func _Greeter_SayHello0_HTTP_Handler(srv GreeterHTTPServer) gin.HandlerFunc {
	return func(c *gin.Context) {
		shouldBind := func(req any) error {
			if err := c.ShouldBindQuery(req); err != nil {
				return err
			}
			if err := c.ShouldBindUri(req); err != nil {
				return err
			}
			return srv.Validate(c.Request.Context(), req)
		}

		var req HelloRequest
		if err := shouldBind(&req); err != nil {
			srv.ErrorEncoder(c, err, true)
			return
		}
		result, err := srv.SayHello(c.Request.Context(), &req)
		if err != nil {
			srv.ErrorEncoder(c, err, false)
			return
		}
		c.JSON(200, result)
	}
}
