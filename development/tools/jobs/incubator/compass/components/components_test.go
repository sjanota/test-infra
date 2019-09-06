package components

import (
	"github.com/kyma-project/test-infra/development/tools/jobs/tester"
	"github.com/kyma-project/test-infra/development/tools/jobs/tester/buildjob"
	"testing"
)

var components = []struct {
	path              string
	image             string
	additionalOptions []buildjob.Option
}{
	{path: "connector", image: tester.ImageGolangBuildpack1_11},
	{path: "director", image: tester.ImageGolangBuildpack1_11},
	{path: "gateway", image: tester.ImageGolangBuildpack1_11},
	{path: "healthchecker", image: tester.ImageGolangBuildpack1_11},
	{path: "provisioner", image: tester.ImageGolangBuildpack1_11},
	{path: "schema-migrator", image: tester.ImageGolangBuildpack1_11},
}

func TestComponentJobs(t *testing.T) {
	for _, component := range components {
		t.Run(component.path, func(t *testing.T) {
			opts := []buildjob.Option{
				buildjob.Component(component.path, component.image),
				buildjob.CompassRepo(),
				buildjob.NoReleases(),
				buildjob.RepositoryRoot("./../../../../../../"),
			}
			opts = append(opts, component.additionalOptions...)
			buildjob.NewSuite(opts...).Run(t)
		})
	}
}
