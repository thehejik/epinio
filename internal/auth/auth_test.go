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

// Package auth collects structures and functions around the
// generation and processing of credentials.
package auth_test

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/epinio/epinio/helpers/kubernetes"
	"github.com/epinio/epinio/internal/auth"
	"github.com/epinio/epinio/internal/auth/authfakes"
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var r *rand.Rand

var _ = Describe("Auth users", func() {
	var authService *auth.AuthService
	var fakeSecret *authfakes.FakeSecretInterface
	var fakeConfigMap *authfakes.FakeConfigMapInterface

	BeforeEach(func() {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))

		fakeSecret = &authfakes.FakeSecretInterface{}
		fakeConfigMap = &authfakes.FakeConfigMapInterface{}

		authService = &auth.AuthService{
			Logger:             logr.Discard(),
			SecretInterface:    fakeSecret,
			ConfigMapInterface: fakeConfigMap,
		}
	})

	Describe("GetUsers", func() {

		When("kubernetes returns an error", func() {
			It("returns an error and an empty slice", func() {
				fakeSecret.ListReturns(nil, errors.New("an error"))

				users, err := authService.GetUsers(context.Background())
				Expect(err).To(HaveOccurred())
				Expect(users).To(BeEmpty())
			})
		})

		When("kubernetes returns an empty list of secrets", func() {
			It("returns an empty slice", func() {
				fakeSecret.ListReturns(&corev1.SecretList{Items: []corev1.Secret{}}, nil)

				users, err := authService.GetUsers(context.Background())
				Expect(err).ToNot(HaveOccurred())
				Expect(users).To(BeEmpty())
			})
		})

		When("kubernetes returns some user secrets", func() {
			It("returns a list of users", func() {
				userSecrets := []corev1.Secret{
					newUserSecret("admin", "password", "admin", ""),
					newUserSecret("epinio", "mypass", "user", "workspace\nworkspace2"),
				}
				fakeSecret.ListReturns(&corev1.SecretList{Items: userSecrets}, nil)

				users, err := authService.GetUsers(context.Background())
				Expect(err).ToNot(HaveOccurred())
				Expect(users).To(HaveLen(2))
				Expect(users[0].Username).To(Equal("admin"))
				Expect(users[0].Namespaces).To(HaveLen(0))
				Expect(users[1].Namespaces).To(HaveLen(2))
			})
		})
	})

	Describe("RemoveNamespaceFromUsers", func() {

		When("users have the namespace", func() {
			It("will be removed", func() {
				userSecrets := []corev1.Secret{
					newUserSecret("user1", "password", "admin", ""),
					newUserSecret("user2", "password", "user", "workspace\nworkspace2"),
					newUserSecret("user3", "password", "user", "workspace"),
				}

				// setup mock
				fakeSecret.ListReturns(&corev1.SecretList{Items: userSecrets}, nil)

				fakeSecret.GetReturnsOnCall(0, &userSecrets[1], nil)
				updatedUser2 := newUserSecret("user2", "password", "user", "workspace2")
				fakeSecret.UpdateReturnsOnCall(0, &updatedUser2, nil)

				fakeSecret.GetReturnsOnCall(1, &userSecrets[2], nil)
				updatedUser3 := newUserSecret("user3", "password", "user", "")
				fakeSecret.UpdateReturnsOnCall(1, &updatedUser3, nil)

				// do test
				err := authService.RemoveNamespaceFromUsers(context.Background(), "workspace")
				Expect(err).ToNot(HaveOccurred())

				_, secretName, _ := fakeSecret.GetArgsForCall(0)
				Expect(secretName).To(Equal("user2"))

				_, secretName, _ = fakeSecret.GetArgsForCall(1)
				Expect(secretName).To(Equal("user3"))

				Expect(fakeSecret.GetCallCount()).To(Equal(2))
			})
		})
	})

	Describe("UpdateUsers", func() {

		When("updating user with different role and namespaces", func() {
			It("will be updated", func() {
				oldUser := newUserSecret("user2", "password", "user", "workspace\nworkspace2")

				// setup mock
				fakeSecret.GetReturns(&oldUser, nil)
				fakeSecret.UpdateReturns(nil, nil)

				// do test
				result, err := authService.UpdateUser(context.Background(), auth.User{
					Role:       "another-role",
					Namespaces: []string{},
				})

				Expect(err).ToNot(HaveOccurred())
				Expect(result.Role).To(Equal("another-role"))
				Expect(result.Namespaces).To(HaveLen(0))
			})
		})
	})
})

func newUserSecret(username, password, role, namespaces string) corev1.Secret {
	return corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: username,
			Labels: map[string]string{
				kubernetes.EpinioAPISecretRoleLabelKey: role,
			},
			CreationTimestamp: metav1.NewTime(newRandomDate()),
		},
		Data: map[string][]byte{
			"username":   []byte(username),
			"password":   []byte(password),
			"namespaces": []byte(namespaces),
		},
	}
}

func newRandomDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min

	sec := r.Int63n(delta) + min
	return time.Unix(sec, 0)
}
