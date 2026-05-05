<<<<<<< HEAD
package controller

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("ServerBootConfiguration Controller", func() {
	Context("When reconciling a resource", func() {

		It("should successfully reconcile the resource", func() {

			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.
		})
=======
// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	metalv1alpha1 "github.com/ironcore-dev/metal-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	. "sigs.k8s.io/controller-runtime/pkg/envtest/komega"
)

var _ = Describe("ServerBootConfiguration Controller", func() {
	ns := SetupTest(nil)

	var server *metalv1alpha1.Server

	BeforeEach(func(ctx SpecContext) {
		By("Creating a Server object")
		server = &metalv1alpha1.Server{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-severbootconfig-",
				Annotations: map[string]string{
					metalv1alpha1.OperationAnnotation: metalv1alpha1.OperationAnnotationIgnore,
				},
			},
			Spec: metalv1alpha1.ServerSpec{
				SystemUUID: "38947555-7742-3448-3784-823347823834",
			},
		}
		Expect(k8sClient.Create(ctx, server)).To(Succeed())
	})

	AfterEach(func(ctx SpecContext) {
		Expect(k8sClient.Delete(ctx, server)).To(Succeed())
		EnsureCleanState()
	})

	It("Should successfully add the boot configuration ref to server", func(ctx SpecContext) {
		By("By creating a server boot configuration")
		config := &metalv1alpha1.ServerBootConfiguration{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns.Name,
				Name:      server.Name,
			},
			Spec: metalv1alpha1.ServerBootConfigurationSpec{
				ServerRef: v1.LocalObjectReference{Name: server.Name},
				Image:     "foo:latest",
			},
		}
		Expect(k8sClient.Create(ctx, config)).To(Succeed())

		Eventually(Object(config)).Should(SatisfyAll(
			HaveField("Status.State", metalv1alpha1.ServerBootConfigurationStatePending),
		))
>>>>>>> tmp-original-05-05-26-00-31
	})
})
