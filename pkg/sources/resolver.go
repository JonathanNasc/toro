package sources

import (
	"github.com/JonathanNasc/toroburro/pkg/sources/sources_domain"
	"github.com/JonathanNasc/toroburro/pkg/sources/sources_services"
)

func Resolve(source_name string) sources_domain.Source {
	switch source_name {
	case alpha_vantage.SourceName:
		return &alpha_vantage.AlphaVantage{}
	default:
		panic("This source name is not allowed: " + source_name)
	}
}
