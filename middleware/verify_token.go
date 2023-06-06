package middleware

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"stock-service/conf"
	proto "stock-service/proto/authen"
	"strings"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")
		data := strings.Split(token, " ")[1]
		conn, err := grpc.Dial(conf.GetConfig().GRPCServer.AuthenServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			logrus.Error("did not connect: %v", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		defer conn.Close()
		client := proto.NewAuthenServiceClient(conn)
		_, err = client.VerifyToken(context.Background(), &proto.String{Value: data})
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}
		return next(c)
	}
}
