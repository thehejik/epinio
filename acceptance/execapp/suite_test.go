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

package install_test2

import (
	"fmt"
	"os"
	"testing"

	"github.com/epinio/epinio/acceptance/testenv"

	//. "github.com/epinio/epinio/acceptance/helpers/matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(FailWithReport)
	RunSpecs(t, "Acceptance Suite")
}

var (
	nodeTmpDir string
	// Lets see if ok with init
	env testenv.EpinioEnv
)

var _ = SynchronizedBeforeSuite(func() []byte {
	//ingressController := os.Getenv("INGRESS_CONTROLLER")

	By("Installing and configuring the prerequisites", func() {
		testenv.SetRoot("../..")
		testenv.SetupEnv()

		env = testenv.New(nodeTmpDir, testenv.Root(), "admin", "password", "", "")
	})

	return []byte{}
}, func(_ []byte) {
	testenv.SetRoot("../..")
	testenv.SetupEnv()

	Expect(os.Getenv("KUBECONFIG")).ToNot(BeEmpty(), "KUBECONFIG environment variable should not be empty")
})

var _ = SynchronizedAfterSuite(func() {
}, func() { // Runs only on one node after all are done
	if testenv.SkipCleanup() {
		fmt.Printf("Found '%s', skipping all cleanup", testenv.SkipCleanupPath())
	} else {
		// Delete left-overs no matter what
		//defer func() { _, _ = testenv.CleanupTmp() }()
		fmt.Print("\nCleanup skipped - skipping anyway")
	}
})

var _ = AfterEach(func() {
	testenv.AfterEachSleep()
})

func FailWithReport(message string, callerSkip ...int) {
	// NOTE: Use something like the following if you need to debug failed tests
	// fmt.Println("\nA test failed. You may find the following information useful for debugging:")
	// fmt.Println("The cluster pods: ")
	// out, err := proc.Kubectl("get pods --all-namespaces")
	// if err != nil {
	// 	fmt.Print(err.Error())
	// } else {
	// 	fmt.Print(out)
	// }

	// Ensures the correct line numbers are reported
	Fail(message, callerSkip[0]+1)
}
