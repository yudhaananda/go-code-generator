package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yudhaananda/go-code-generator/models"
	"github.com/yudhaananda/go-code-generator/services"
)

type handler struct {
	Service *services.Services
}

func Init(service *services.Services) *handler {
	return &handler{
		Service: service,
	}
}

func (h *handler) Run() {
	if err := h.register().Run(); err != nil {
		panic(err)
	}
}

func (h *handler) register() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
	}))

	// Main
	router.GET("/", h.Index)
	router.GET("/favicon.ico", h.Icon)
	router.POST("/project.json", h.Template)

	api := router.Group("/api/v1")

	// API Route
	api.POST("/generate", h.GenerateAPI)

	return router
}

func (h *handler) GenerateAPI(ctx *gin.Context) {
	content, err := ctx.FormFile("file")
	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}

	split := strings.Split(content.Filename, ".")
	extention := split[len(split)-1]

	if extention != "json" {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte("Only .json file"))
		return
	}

	file, err := content.Open()

	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}

	var jsonContent models.Model
	err = json.Unmarshal(data, &jsonContent)

	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}

	result, err := h.Service.Generate.Generate(jsonContent)
	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	ctx.Data(http.StatusOK, "Application/zip", result)
}

func (h *handler) Index(ctx *gin.Context) {
	html, err := os.ReadFile("index.html")
	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", html)
}

func (h *handler) Icon(ctx *gin.Context) {
	png, err := os.ReadFile("favicon.ico")
	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
	}
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", png)
}

func (h *handler) Template(ctx *gin.Context) {
	template, err := os.ReadFile("project.json")
	if err != nil {
		ctx.Data(http.StatusBadGateway, "text/html; charset=utf-8", []byte(err.Error()))
		return
	}
	ctx.Data(http.StatusOK, "Application/file", template)

}
