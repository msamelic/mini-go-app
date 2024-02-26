package ginrouter

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mini-go-app/internal/util/envconf"
	"net/http"
)

func New(env *envconf.Spec, log *zap.Logger) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	NewHandler(r, env, log)

	return r
}

type Handler struct {
	Router *gin.Engine
	Env    *envconf.Spec
	Log    *zap.Logger
}

func NewHandler(r *gin.Engine, env *envconf.Spec, log *zap.Logger) {
	h := &Handler{
		Router: r,
		Env:    env,
		Log:    log,
	}
	r.GET("/health", h.HealthCheck())
}

func (h *Handler) HealthCheck() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}
