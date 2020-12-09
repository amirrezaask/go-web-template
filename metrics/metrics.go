package metrics

import (
	"app/config"

	grpc_prom "github.com/grpc-ecosystem/go-grpc-prometheus"
	echo_prom "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

//RegisterEchoMetrics registers echo metrics
func RegisterEchoMetrics(e *echo.Echo) {
	p := echo_prom.NewPrometheus(config.AppName, nil)
	e.Use(p.HandlerFunc)
}

//RegisterGRPCMetrics registers GRPC metrics in prometheus
func RegisterGRPCMetrics(server *grpc.Server) {
	grpc_prom.Register(server)
}

func PrometheusHandlerFunc() echo.HandlerFunc {
	h := promhttp.Handler()
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
