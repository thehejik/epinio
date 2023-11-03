// Copyright © 2021 - 2023 SUSE LLC
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package install_test

import (
	//"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"time"

	"github.com/Netflix/go-expect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	//"github.com/epinio/epinio/acceptance/helpers/catalog" //needed for NewAppName
	"github.com/epinio/epinio/acceptance/helpers/epinio"
	"github.com/epinio/epinio/acceptance/helpers/proc"

	//"github.com/epinio/epinio/acceptance/helpers/route53"
	"github.com/epinio/epinio/acceptance/testenv"
)

// This test uses AWS route53 to update the system domain's records
var _ = Describe("<Scenario2> GKE, Letsencrypt-staging, deploy instance(s)", func() {
	var (
		//	flags         []string
		epinioHelper epinio.Epinio
		//appName      = catalog.NewAppName()
		appNameStatic string = "apps-430007115"
		//	loadbalancer  string
		domain string
		//	zoneID        string
		instancesNum string
	//	extraEnvName  string
	//	extraEnvValue string
	//	name_exists   bool
	//	value_exists  bool
	)

	BeforeEach(func() {
		epinioHelper = epinio.NewEpinioHelper(testenv.EpinioBinaryPath())

		domain = os.Getenv("EPINIO_SYSTEM_DOMAIN")
		Expect(domain).ToNot(BeEmpty())

		//zoneID = os.Getenv("AWS_ZONE_ID")
		//Expect(zoneID).ToNot(BeEmpty())

		instancesNum = "1"

		//flags = []string{
		//	"--set", "server.disableTracking=true", // disable tracking during tests
		//	"--set", "global.domain=" + domain,
		//	"--set", "global.customTlsIssuer=letsencrypt-staging", // let epinio use existing clusterIssuer
		//}

		//extraEnvName, name_exists = os.LookupEnv("EXTRAENV_NAME")
		//extraEnvValue, value_exists = os.LookupEnv("EXTRAENV_VALUE")
		//if name_exists && value_exists {
		//	flags = append(flags, "--set", "extraEnv[0].name="+extraEnvName, "--set-string", "extraEnv[0].value="+extraEnvValue)
		//}
	})

	AfterEach(func() {
		//out, err := epinioHelper.Uninstall()
		//Expect(err).NotTo(HaveOccurred(), out)
	})
	//
	It("Installs with letsencrypt-staging cert, custom domain and pushes an app with "+instancesNum+" instances", func() {
		//By("Creating letsencrypt issuer", func() {
		//	// Create certificate secret and cluster_issuer
		//	out, err := proc.RunW("kubectl", "apply", "-f",
		//		testenv.TestAssetPath("letsencrypt-staging.yaml"))
		//	Expect(err).NotTo(HaveOccurred(), out)
		//})
		//
		//By("Checking LoadBalancer IP", func() {
		//	// Ensure that Traefik LB is not in Pending state anymore, could take time
		//	Eventually(func() string {
		//		out, err := proc.RunW("kubectl", "get", "svc", "-n", "traefik", "traefik", "--no-headers")
		//		Expect(err).NotTo(HaveOccurred(), out)
		//		return out
		//	}, "4m", "2s").ShouldNot(ContainSubstring("<pending>"))
		//
		//	out, err := proc.RunW("kubectl", "get", "service", "-n", "traefik", "traefik", "-o", "json")
		//	Expect(err).NotTo(HaveOccurred(), out)
		//
		//	// Check that an IP address for LB is configured
		//	status := &testenv.LoadBalancerHostname{}
		//	err = json.Unmarshal([]byte(out), status)
		//	Expect(err).NotTo(HaveOccurred())
		//	Expect(status.Status.LoadBalancer.Ingress).To(HaveLen(1))
		//	loadbalancer = status.Status.LoadBalancer.Ingress[0].IP
		//	Expect(loadbalancer).ToNot(BeEmpty())
		//})
		//
		//By("Updating DNS Entries", func() {
		//	change := route53.A(domain, loadbalancer, "UPSERT")
		//	out, err := route53.Update(zoneID, change, nodeTmpDir)
		//	Expect(err).NotTo(HaveOccurred(), out)
		//
		//	change = route53.A("*."+domain, loadbalancer, "UPSERT")
		//	out, err = route53.Update(zoneID, change, nodeTmpDir)
		//	Expect(err).NotTo(HaveOccurred(), out)
		//})
		//
		//// Check that DNS entry is correctly propagated
		//By("Checking that DNS entry is correctly propagated", func() {
		//	Eventually(func() string {
		//		out, err := route53.TestDnsAnswer(zoneID, domain, "A")
		//		Expect(err).NotTo(HaveOccurred(), out)
		//
		//		answer := &route53.DNSAnswer{}
		//		err = json.Unmarshal([]byte(out), answer)
		//		Expect(err).NotTo(HaveOccurred())
		//		if len(answer.RecordData) == 0 {
		//			return ""
		//		}
		//		return answer.RecordData[0]
		//	}, "5m", "2s").Should(Equal(loadbalancer))
		//})
		//
		//// Workaround to (try to!) ensure that the DNS is really propagated!
		//time.Sleep(3 * time.Minute)
		//
		//By("Installing Epinio", func() {
		//	out, err := epinioHelper.Install(flags...)
		//	Expect(err).NotTo(HaveOccurred(), out)
		//	Expect(out).To(ContainSubstring("STATUS: deployed"))
		//
		//	out, err = testenv.PatchEpinio()
		//	Expect(err).ToNot(HaveOccurred(), out)
		//})
		//
		By("Connecting to Epinio", func() {
			Eventually(func() string {
				out, _ := epinioHelper.Run("login", "-u", "admin", "-p", "password", "--trust-ca", "https://epinio."+domain)
				return out
			}, "10s", "5s").Should(ContainSubstring("Login successful"))
		})

		By("Checking Epinio info command", func() {
			Eventually(func() string {
				out, _ := epinioHelper.Run("info")
				return out
			}, "10s", "2s").Should(ContainSubstring("Epinio Server Version:"))
		})

		//		By("Pushing an app with "+instancesNum+" instances", func() {
		//			out, err := epinioHelper.Run("push",
		//				"--name", appName,
		//				"--path", testenv.AssetPath("sample-app"),
		//				"--instances", instancesNum)
		//			Expect(err).ToNot(HaveOccurred(), out)
		//
		//			Eventually(func() string {
		//				out, err := proc.Kubectl("get", "deployments",
		//					"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s", appName, testenv.DefaultWorkspace),
		//					"--namespace", testenv.DefaultWorkspace,
		//					"-o", "jsonpath={.items[].spec.replicas}")
		//				Expect(err).ToNot(HaveOccurred(), out)
		//				return out
		//			}, "30s", "1s").Should(Equal(instancesNum))
		//
		//			// Verify cluster_issuer is used
		//			out, err = proc.RunW("kubectl", "get", "certificate",
		//				"-n", testenv.DefaultWorkspace,
		//				"--selector", "app.kubernetes.io/name="+appName,
		//				"-o", "jsonpath='{.items[*].spec.issuerRef.name}'")
		//			Expect(err).NotTo(HaveOccurred(), out)
		//			Expect(out).To(Equal("'letsencrypt-production'"))
		//
		//			// Wait until all the CertRequests in all namespaces are Ready -> the CSRs have been approved and signed
		//			out, err = proc.RunW("kubectl", "wait", "--for=condition=Ready",
		//				"certificaterequest", "--selector", "app.kubernetes.io/managed-by=Helm",
		//				"--all-namespaces", "--timeout=120s")
		//			Expect(err).NotTo(HaveOccurred(), out)
		//		})

		By("Exec to running application", func() {
			// Create complete exec command
			p, err := proc.Get("", testenv.EpinioBinaryPath(), "app", "exec", appNameStatic)
			Expect(err).NotTo(HaveOccurred())

			args := strings.Fields(p.String())
			command := args[0]
			commandArgs := args[1:]

			// Debug
			fmt.Printf("\nCommand to be executed: %s\n\n", p.String())

			// Specify a timeout
			defaultExpectTimout := 5 * time.Second
			console, err := expect.NewConsole(expect.WithStdout(os.Stdout), expect.WithStdin(os.Stdin), expect.WithDefaultTimeout(defaultExpectTimout))
			Expect(err).NotTo(HaveOccurred())
			defer console.Close()

			cmd := exec.Command(command, commandArgs...)
			//fmt.Printf("\nCommand to be executed: %s\n\n", cmd.String())
			cmd.Stdin = console.Tty()
			cmd.Stdout = console.Tty()
			cmd.Stderr = console.Tty()

			err = cmd.Start()
			Expect(err).NotTo(HaveOccurred())
			defer cmd.Wait()
			// Interact with the program.
			str, err := console.ExpectString("Executing a shell")
			if err != nil {
				log.Fatalf("got %s", str)
			}
			// Wait for shell
			console.ExpectString(".*@.*" + appNameStatic + ".*\\$#")
			// Check whether running on Packeto container
			console.SendLine("cat /etc/os-release")
			str, err = console.ExpectString("Paketo Buildpacks")
			if err != nil {
				log.Fatalf("got %s", str)
			}
			console.SendLine("exit")
			console.ExpectEOF()
			Expect(true).To(BeTrue()) // Replace with your actual expectations.
		})

		//		By("Pushing an app with "+instancesNum+" instances, and not verifying certs", func() {
		//			appName2 := appName + "-ssv"
		//			out, err := epinioHelper.Run("push",
		//				"--skip-ssl-verification",
		//				"--name", appName2,
		//				"--path", testenv.AssetPath("sample-app"),
		//				"--instances", instancesNum)
		//			Expect(err).ToNot(HaveOccurred(), out)
		//
		//			Eventually(func() string {
		//				out, err := proc.Kubectl("get", "deployments",
		//					"-l", fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/part-of=%s",
		//						appName2, testenv.DefaultWorkspace),
		//					"--namespace", testenv.DefaultWorkspace,
		//					"-o", "jsonpath={.items[].spec.replicas}")
		//				Expect(err).ToNot(HaveOccurred(), out)
		//				return out
		//			}, "30s", "1s").Should(Equal(instancesNum))
		//
		//			// Verify cluster_issuer is used
		//			out, err = proc.RunW("kubectl", "get", "certificate",
		//				"-n", testenv.DefaultWorkspace,
		//				"--selector", "app.kubernetes.io/name="+appName2,
		//				"-o", "jsonpath='{.items[*].spec.issuerRef.name}'")
		//			Expect(err).NotTo(HaveOccurred(), out)
		//			Expect(out).To(Equal("'letsencrypt-production'"))
		//
		//			// Wait until all the CertRequests in all namespaces are Ready -> the CSRs have been approved and signed
		//			out, err = proc.RunW("kubectl", "wait", "--for=condition=Ready",
		//				"certificaterequest", "--selector", "app.kubernetes.io/managed-by=Helm",
		//				"--all-namespaces", "--timeout=120s")
		//			Expect(err).NotTo(HaveOccurred(), out)
		//		})
	})
})
