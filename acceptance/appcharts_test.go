package acceptance_test

import (
	. "github.com/epinio/epinio/acceptance/helpers/matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("apps chart", func() {

	Describe("app chart list", func() {
		It("lists the known app charts", func() {
			out := env.Epinio("", "apps", "chart", "list")
			Expect(out).To(ContainSubstring("Show Application Charts"))

			Expect(out).To(
				HaveATable(
					WithHeaders("DEFAULT", "NAME", "CREATED", "DESCRIPTION"),
					WithRow("standard", WithDate(), "Epinio standard deployment"),
				),
			)
		})
	})

	Describe("app chart show", func() {
		It("shows the details of the standard app chart", func() {
			out := env.Epinio("", "apps", "chart", "show", "standard")
			Expect(out).To(ContainSubstring("Show application chart details"))

			Expect(out).To(
				HaveATable(
					WithHeaders("KEY", "VALUE"),
					WithRow("Name", "standard"),
					WithRow("Created", WithDate()),
					WithRow("Short", "Epinio standard deployment"),
					WithRow("Description", "Epinio standard support chart"),
					WithRow("", "for application deployment"),
					WithRow("Helm Repository", ""),
					WithRow("Helm Chart", "https.*epinio-application.*tgz"),
				),
			)
		})

		It("fails to show the details of a bogus app chart", func() {
			out := env.Epinio("", "apps", "chart", "show", "bogus")
			Expect(out).To(ContainSubstring("Show application chart details"))
			Expect(out).To(ContainSubstring("Not Found: application chart 'bogus' does not exist"))
		})
	})

	Describe("app chart default", func() {
		AfterEach(func() {
			// Reset to empty default as the state to be seen at the
			// beginning of each test, regardless of ordering.
			_ = env.Epinio("", "apps", "chart", "default", "")
		})

		It("shows nothing by default", func() {
			out := env.Epinio("", "apps", "chart", "default")
			Expect(out).To(ContainSubstring("Name: not set, system default applies"))
		})

		It("sets a default", func() {
			out := env.Epinio("", "apps", "chart", "default", "standard")
			Expect(out).To(ContainSubstring("New Default Application Chart"))
			Expect(out).To(ContainSubstring("Name: standard"))

			out = env.Epinio("", "apps", "chart", "default")
			Expect(out).To(ContainSubstring("Name: standard"))
		})

		It("fails to sets a bogus default", func() {
			out := env.Epinio("", "apps", "chart", "default", "bogus")
			Expect(out).To(ContainSubstring("Not Found: application chart 'bogus' does not exist"))
		})

		It("unsets a default", func() {
			By("setting default")
			_ = env.Epinio("", "apps", "chart", "default", "standard")

			By("unsetting default")
			out := env.Epinio("", "apps", "chart", "default", "")
			Expect(out).To(ContainSubstring("Unset Default Application Chart"))

			out = env.Epinio("", "apps", "chart", "default")
			Expect(out).To(ContainSubstring("Name: not set, system default applies"))
		})
	})
})
