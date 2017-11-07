package main_test

import (
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	io_helpers "code.cloudfoundry.org/cli/util/testhelpers/io"

	. "github.com/clijockey/cowsay"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Cowsay", func() {
	Describe(".Run", func() {
		var fakeCliConnection *pluginfakes.FakeCliConnection
		var cowsay *Cowsay
		var outputChan chan []string

		BeforeEach(func() {
			outputChan = make(chan []string)

			fakeCliConnection = &pluginfakes.FakeCliConnection{}
			fakeCliConnection.UsernameReturns("user@user.com", nil)
			cowsay = &Cowsay{}
		})

	})

	It("should work for simple", func() {
		c := Cowsay{}
		c.Run(nil, []string{"cowsay"})

	})
})

func invokeCmd(outputChan chan []string, cowsay *Cowsay, fakeCliConnection *pluginfakes.FakeCliConnection) {
	outputChan <- io_helpers.CaptureOutput(func() {
		cowsay.Run(fakeCliConnection, []string{"cowsay"})
	})
}
