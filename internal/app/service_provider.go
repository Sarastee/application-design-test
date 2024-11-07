package app

import (
	"log"

	"github.com/sarastee/application-design-test/internal/config"
	"github.com/sarastee/application-design-test/internal/config/env"
)

type serviceProvider struct {
	httpConfig *config.HTTPConfig

	// orderRepo

	// orderService

	// orderImpl
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// HTTPConfig ...
func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}
