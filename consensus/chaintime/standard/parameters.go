package standard

import (
	eth2client "github.com/attestantio/go-eth2-client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type parameters struct {
	logLevel            zerolog.Level
	genesisTimeProvider eth2client.GenesisTimeProvider
	specProvider        eth2client.SpecProvider
}

// Parameter is the interface for service parameters.
type Parameter interface {
	apply(*parameters)
}

type parameterFunc func(*parameters)

func (f parameterFunc) apply(p *parameters) {
	f(p)
}

// WithLogLevel sets the log level for the module.
func WithLogLevel(logLevel zerolog.Level) Parameter {
	return parameterFunc(func(p *parameters) {
		p.logLevel = logLevel
	})
}

// WithGenesisTimeProvider sets the genesis time provider.
func WithGenesisTimeProvider(provider eth2client.GenesisTimeProvider) Parameter {
	return parameterFunc(func(p *parameters) {
		p.genesisTimeProvider = provider
	})
}

// WithSpecProvider sets the spec provider.
func WithSpecProvider(provider eth2client.SpecProvider) Parameter {
	return parameterFunc(func(p *parameters) {
		p.specProvider = provider
	})
}

// parseAndCheckParameters parses and checks parameters to ensure that mandatory parameters are present and correct.
func parseAndCheckParameters(params ...Parameter) (*parameters, error) {
	parameters := parameters{
		logLevel: zerolog.GlobalLevel(),
	}
	for _, p := range params {
		if params != nil {
			p.apply(&parameters)
		}
	}

	if parameters.specProvider == nil {
		return nil, errors.New("no spec provider specified")
	}
	if parameters.genesisTimeProvider == nil {
		return nil, errors.New("no genesis time provider specified")
	}

	return &parameters, nil
}
