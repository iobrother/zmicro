package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iobrother/zmicro"
	"github.com/iobrother/zmicro/core/log"
	zgin "github.com/iobrother/zmicro/core/transport/http"
	"github.com/iobrother/zmicro/examples/gin/api/api"
)

// curl http://127.0.0.1:5180/hello/zmicro
func main() {
	app := zmicro.New(zmicro.InitHttpServer(InitHttpServer))

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

func InitHttpServer(r *gin.Engine) error {
	gin.DisableBindValidation()

	g := r.Group("/")
	api.RegisterGreeterHTTPServer(g, &GreeterImpl{})

	return nil
}

type GreeterImpl struct {
	zgin.Implemented
}

func (s *GreeterImpl) SayHello(ctx context.Context, req *api.HelloRequest, rsp *api.HelloReply) error {
	*rsp = api.HelloReply{
		Message: fmt.Sprintf("hello %s!", req.Name),
	}

	return nil
}
