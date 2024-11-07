package app

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/sarastee/application-design-test/internal/closer"
	"github.com/sarastee/application-design-test/internal/config"
)

// App ..
type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
	configPath      string
}

// NewApp ..
func NewApp(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

// Run ..
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failure while running HTTP server")
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initDepFunctions := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range initDepFunctions {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := a.initRoutes(ctx)

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider,
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
	}

	return nil
}
