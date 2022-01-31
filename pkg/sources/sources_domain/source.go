package sources_domain

import (
	"github.com/JonathanNasc/toroburro/pkg/assets"
)

type Source interface {
	GetAsset(code string) assets.AssetResult
	GetName() string
}
