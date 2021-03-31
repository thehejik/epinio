package acceptance_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bounds between Apps & Services", func() {
	var org = "apps-org"
	BeforeEach(func() {
		setupAndTargetOrg(org)
	})
	Describe("Display", func() {
		var appName string
		var serviceName string
		BeforeEach(func() {
			appName = newAppName()
			serviceName = newServiceName()

			makeApp(appName)
			makeCustomService(serviceName)
			bindAppService(appName, serviceName, org)
		})
		It("shows the bound app for services list, and vice versa", func() {
			out, err := Carrier("service list", "")
			Expect(err).ToNot(HaveOccurred(), out)
			Expect(out).To(MatchRegexp(serviceName + `.*` + appName))

			Eventually(func() string {
				out, err = Carrier("app list", "")
				Expect(err).ToNot(HaveOccurred(), out)
				return out
			}, "1m").Should(MatchRegexp(appName + `.*\|.*1\/1.*\|.*` + serviceName))
		})
		AfterEach(func() {
			// Delete app first, as this also unbinds the service
			cleanupApp(appName)
			cleanupService(serviceName)
		})
	})
})