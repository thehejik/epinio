package machine

import (
	"fmt"

	"github.com/epinio/epinio/acceptance/helpers/proc"
	. "github.com/onsi/gomega"
)

func (m *Machine) MakeConfiguration(configurationName string) {
	_ = m.Epinio("", "configuration", "create", configurationName, "username", "epinio-user")

	// check presence
	out := m.Epinio("", "configuration", "list")
	ExpectWithOffset(1, out).To(MatchRegexp(configurationName))
}

func (m *Machine) BindAppConfiguration(appName, configurationName, namespace string) {
	_ = m.Epinio("", "configuration", "bind", configurationName, appName)

	// check deep into the kube structures
	m.VerifyAppConfigurationBound(appName, configurationName, namespace, 2)
}

func (m *Machine) VerifyAppConfigurationBound(appName, configurationName, namespace string, offset int) {
	out, err := proc.Kubectl("get", "deployments",
		"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s", appName, namespace),
		"--namespace", namespace,
		"-o", "jsonpath={.items[].spec.template.spec.volumes}")
	ExpectWithOffset(offset, err).ToNot(HaveOccurred(), out)
	ExpectWithOffset(offset, out).To(MatchRegexp(configurationName))

	out, err = proc.Kubectl("get", "deployments",
		"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s", appName, namespace),
		"--namespace", namespace,
		"-o", "jsonpath={.items[].spec.template.spec.containers[0].volumeMounts}")
	ExpectWithOffset(1, err).ToNot(HaveOccurred(), out)
	ExpectWithOffset(1, out).To(MatchRegexp("/configurations/" + configurationName))
}

func (m *Machine) DeleteConfiguration(configurationName string) {
	m.DeleteConfigurationWithUnbind(configurationName, false)
}

func (m *Machine) DeleteConfigurationUnbind(configurationName string) {
	m.DeleteConfigurationWithUnbind(configurationName, true)
}

func (m *Machine) DeleteConfigurationWithUnbind(configurationName string, unbind bool) {
	if unbind {
		_ = m.Epinio("", "configuration", "delete", configurationName, "--unbind")
	} else {
		_ = m.Epinio("", "configuration", "delete", configurationName)
	}

	// And check non-presence
	EventuallyWithOffset(1,
		m.Epinio("", "configuration", "list"),
		"10m",
	).ShouldNot(MatchRegexp(configurationName))
}

func (m *Machine) CleanupConfiguration(configurationName string) {
	_ = m.Epinio("", "configuration", "delete", configurationName)
}

func (m *Machine) UnbindAppConfiguration(appName, configurationName, namespace string) {
	_ = m.Epinio("", "configuration", "unbind", configurationName, appName)

	// deep check in kube structures for non-presence
	m.VerifyAppConfigurationNotbound(appName, configurationName, namespace, 2)
}

func (m *Machine) VerifyAppConfigurationNotbound(appName, configurationName, namespace string, offset int) {
	out, err := proc.Kubectl("get", "deployments",
		"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s", appName, namespace),
		"--namespace", namespace,
		"-o", "jsonpath={.items[].spec.template.spec.volumes}")
	ExpectWithOffset(offset, err).ToNot(HaveOccurred(), out)
	ExpectWithOffset(offset, out).ToNot(MatchRegexp(configurationName))

	out, err = proc.Kubectl("get", "deployments",
		"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s", appName, namespace),
		"--namespace", namespace,
		"-o", "jsonpath={.items[].spec.template.spec.containers[0].volumeMounts}")
	ExpectWithOffset(offset, err).ToNot(HaveOccurred(), out)
	ExpectWithOffset(offset, out).ToNot(MatchRegexp("/configurations/" + configurationName))
}
