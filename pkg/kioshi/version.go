package kioshi

import (
	"github.com/deislabs/porter/pkg/mixin"
	"github.com/deislabs/porter/pkg/porter/version"
	"github.com/ffoysal/try-mixings/pkg"
)

func (m *Mixin) PrintVersion(opts version.Options) error {
	metadata := mixin.Metadata{
		Name: "kioshi",
		VersionInfo: mixin.VersionInfo{
			Version: pkg.Version,
			Commit:  pkg.Commit,
			Author:  "Foysal",
		},
	}
	return version.PrintVersion(m.Context, opts, metadata)
}
