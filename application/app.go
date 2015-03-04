package application

type App struct {
	name     string
	services map[string]IService
}

func (a *App) SetName(name string) {
	a.name = name
}

func (a *App) Name() string {
	return a.name
}

var (
	apps map[string]*App
)

func NewApp(name string) *App {
	app := new(App)
	app.SetName(name)
	return app
}

func RegisterApp(app *App) {
	apps[app.Name()] = app
}

func GetApp(name string) *App {
	return apps[name]
}
