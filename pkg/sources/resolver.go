package sources

import (
	"os"

	"github.com/JonathanNasc/toroburro/pkg/sources/sources_domain"
	"github.com/JonathanNasc/toroburro/pkg/sources/sources_services"
)

type SourcesByCode map[string]sources_domain.Source

var inMemorySources SourcesByCode

func Resolve(sourceName string) sources_domain.Source {
	setInMemorySourcesIfItIsNotAlreadSet()
	if source, ok := inMemorySources[sourceName]; ok {
		return source
	}

	panic("Source name not allowed: " + sourceName)
}

func SubscribeSourceForTests(source sources_domain.Source) {
	if os.Getenv("GO_ENV") == "test" {
		inMemorySources = make(SourcesByCode)
		inMemorySources[source.GetName()] = source
	}
}

func setInMemorySourcesIfItIsNotAlreadSet() {
	if (inMemorySources) == nil {
		inMemorySources = make(SourcesByCode)
		inMemorySources[alpha_vantage.SourceName] = alpha_vantage.New()
	}
}
