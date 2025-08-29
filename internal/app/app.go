package app

import (
	"ai-hr-backend/internal/config"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (*App) Run(cfg config.Config) {

}
