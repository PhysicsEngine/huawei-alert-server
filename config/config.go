package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

const (
	envLocal      = "local"
	envHerokuProd = "heroku_prod"
	envHerokuDev  = "heroku_dev"
)

// Env stores configuration settings extract from environmental variables
// by using https://github.com/kelseyhightower/envconfig
//
// The practice getting from environmental variables comes from https://12factor.net.
type Env struct {
	// Env is environment where application is running. This value is used to
	// annotate datadog metrics or sentry error reporting. The value must be
	// "development" or "production".
	Env  string `envconfig:"ENV" default:"local"`
	Port string `envconfig:"PORT" default:"8888"`
}

// IsProduction returns true if it is production environment
func (e *Env) IsProduction() bool {
	return e.Env == envHerokuProd
}

// validate validates
func (e *Env) validate() error {
	checks := []struct {
		bad    bool
		errMsg string
	}{
		{
			e.Env != envLocal && e.Env != envHerokuProd,
			fmt.Sprintf("invalid env is specifed: %q", e.Env),
		},

		// Add your own validation here
	}

	for _, check := range checks {
		if check.bad {
			return errors.Errorf(check.errMsg)
		}
	}

	return nil
}

// ReadFromEnv reads configuration from environmental variables
// defined by Env struct.
func ReadFromEnv() (*Env, error) {
	var env Env
	if err := envconfig.Process("", &env); err != nil {
		return nil, errors.Wrap(err, "failed to process envconfig")
	}

	if err := env.validate(); err != nil {
		return nil, errors.Wrap(err, "validation failed")
	}

	return &env, nil
}
