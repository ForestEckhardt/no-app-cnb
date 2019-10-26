package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/no-app-cnb/noapp"
)

func main() {
	context, err := build.DefaultBuild()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create a default build context: %s", err)
		os.Exit(101)
	}

	code, err := runBuild(context)
	if err != nil {
		context.Logger.Info(err.Error())
	}

	os.Exit(code)

}

func runBuild(context build.Build) (int, error) {
	context.Logger.Title(context.Buildpack)

	noappContributor, willContribute, err := noapp.NewContributor(context)
	if err != nil {
		return context.Failure(102), err
	}

	if willContribute {
		if err := noappContributor.Contribute(); err != nil {
			return context.Failure(103), err
		}
	}

	return context.Success()
}
