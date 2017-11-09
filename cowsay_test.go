package main_test

import (
	"strings"

	"code.cloudfoundry.org/cli/plugin/models"
	"code.cloudfoundry.org/cli/plugin/pluginfakes"
	io_helpers "code.cloudfoundry.org/cli/util/testhelpers/io"

	. "github.com/clijockey/cowsay"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		It("should work for simple", func() {
			c := Cowsay{}
			Ω(func() { c.Run(fakeCliConnection, []string{"cowsay"}) }).Should(Not(Panic()))

		})

		Context("when the user has an API set", func() {
			BeforeEach(func() {
				fakeCliConnection.HasAPIEndpointReturns(true, nil)
				fakeCliConnection.ApiEndpointReturns("http://run.pivotal.io", nil)
			})

			Context("and the user is logged in", func() {
				BeforeEach(func() {
					fakeCliConnection.IsLoggedInReturns(true, nil)
					fakeCliConnection.UsernameReturns("user@user.com", nil)
				})
				Context("and there is a username", func() {
					BeforeEach(func() {
						fakeCliConnection.UsernameReturns("user@user.com", nil)
					})
					// It("displays the username", func(done Done) {
					// 	defer close(done)
					// 	go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

					// 	var output []string
					// 	Eventually(outputChan, 2).Should(Receive(&output))
					// 	outputString := strings.Join(output, "")
					// 	Expect(outputString).To(ContainSubstring("You are logged in"))
					// 	Expect(outputString).To(ContainSubstring("user@user.com"))
					// 	Expect(outputString).To(ContainSubstring("http://run.pivotal.io"))
					// })

					Context("and the user has a space and org", func() {
						BeforeEach(func() {
							fakeCliConnection.UsernameReturns("user@user.com", nil)
							fakeCliConnection.HasSpaceReturns(true, nil)
							fakeCliConnection.HasOrganizationReturns(true, nil)
							fakeCliConnection.GetCurrentOrgReturns(plugin_models.Organization{OrganizationFields: plugin_models.OrganizationFields{Name: "testOrg"}}, nil)
							fakeCliConnection.GetCurrentSpaceReturns(plugin_models.Space{SpaceFields: plugin_models.SpaceFields{Name: "testSpace"}}, nil)
						})
						It("shows the org and space", func(done Done) {
							defer close(done)
							go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

							var output []string
							Eventually(outputChan, 2).Should(Receive(&output))
							outputString := strings.Join(output, "")
							Expect(outputString).To(ContainSubstring("You are targeting"))
							Expect(outputString).To(ContainSubstring("testOrg"))
							Expect(outputString).To(ContainSubstring("testSpace"))
						})
					})
					// Context("and the user has no space", func() {
					// 	BeforeEach(func() {
					// 		fakeCliConnection.HasSpaceReturns(false, nil)
					// 	})

					// 	It("shows no information about target", func(done Done) {
					// 		defer close(done)
					// 		go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

					// 		var output []string
					// 		Eventually(outputChan, 2).Should(Receive(&output))
					// 		outputString := strings.Join(output, "")
					// 		Expect(outputString).ToNot(ContainSubstring("You are targeting"))
					// 	})
					// })
					// Context("and the user has no org", func() {
					// 	BeforeEach(func() {
					// 		fakeCliConnection.HasOrganizationReturns(false, nil)
					// 	})

					// 	It("shows no information about target", func(done Done) {
					// 		defer close(done)
					// 		go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

					// 		var output []string
					// 		Eventually(outputChan, 2).Should(Receive(&output))
					// 		outputString := strings.Join(output, "")
					// 		Expect(outputString).ToNot(ContainSubstring("You are targeting"))
					// 	})
					// })
				})

				// Context("and there is no username", func() {
				// 	BeforeEach(func() {
				// 		fakeCliConnection.UsernameReturns("", nil)
				// 	})

				// 	It("prints out a helpful error message", func(done Done) {
				// 		defer GinkgoRecover()
				// 		defer close(done)
				// 		go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

				// 		var output []string
				// 		Eventually(outputChan, 2).Should(Receive(&output))
				// 		outputString := strings.Join(output, "")
				// 		Expect(outputString).To(ContainSubstring("you are logged in, but your username is empty"))
				// 	})
				// 	})

				// })

				// Context("and the user is not logged in", func() {
				// 	BeforeEach(func() {
				// 		fakeCliConnection.IsLoggedInReturns(false, nil)
				// 	})

				// 	It("prints out a helpful error message", func(done Done) {
				// 		defer GinkgoRecover()
				// 		defer close(done)
				// 		go invokeCmd(outputChan, whoamiCmd, fakeCliConnection)

				// 		var output []string
				// 		Eventually(outputChan, 2).Should(Receive(&output))
				// 		outputString := strings.Join(output, "")
				// 		Expect(outputString).To(ContainSubstring("Nobody is logged in"))
				// 	})
				// })

			})
		})

		// It("should really also work for spaces", func() {
		// 	c := Cowsay{}
		// 	Ω(func() { c.Run(fakeCliConnection, []string{"cowsay-space"}) }).Should(Not(Panic()))

		// })

	})

})

func invokeCmd(outputChan chan []string, cowsay *Cowsay, fakeCliConnection *pluginfakes.FakeCliConnection) {
	outputChan <- io_helpers.CaptureOutput(func() {
		cowsay.Run(fakeCliConnection, []string{"cowsay"})
	})
}
