package main_test

import (
	"code.cloudfoundry.org/cli/util/testhelpers/pluginbuilder"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCowsayPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	pluginbuilder.BuildTestBinary(".", "cowsay")
	RunSpecs(t, "Cowsay Suite")
}
