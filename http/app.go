package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"golang-softplayer-apply/controller"
	"golang-softplayer-apply/repository"
	service2 "golang-softplayer-apply/service"
	"os"
)

type App struct {
	engine     *gin.Engine
	controller *controller.Controller
}

func NewApp() *App {
	app := App{}
	app.engine = gin.New()
	db, err := NewDB()
	if err != nil {
		os.Exit(1)
	}
	repository := repository.NewRepository(db)
	service := service2.NewService(*repository)
	app.controller = controller.NewController(*service)

	return &app
}

func (a *App) UrlMapping() {
	a.engine.GET("/ping", a.controller.Get)
	a.engine.GET("/person", a.controller.GetAllPeople)
}

func (a *App) Run(port string) error {
	return a.engine.Run(port)
}

func NewDB() (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "whatever", "whatever", "whatever", "whatever")
	connect, err := sqlx.Connect("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	return connect, nil
}
