package hexagonalgo

import "context"

// Here make your app starting point

type App struct {
	repo Repo
}

func NewApp(r Repo) *App {
	// This will probably need some lifecycle init
	return &App{repo: r}
}

// Run will run app
// Always pass context for more control from main
func (a *App) Run(ctx context.Context) error {
	return nil
}
