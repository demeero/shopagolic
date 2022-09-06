package currency

import "context"

type LoaderRepository interface {
	LoadCurrencyCodes(ctx context.Context) ([]string, error)
}

type Loader struct {
	repo LoaderRepository
}

func NewLoader(repo LoaderRepository) *Loader {
	return &Loader{repo: repo}
}

func (l *Loader) LoadCurrencyCodes(ctx context.Context) ([]string, error) {
	return l.repo.LoadCurrencyCodes(ctx)
}
