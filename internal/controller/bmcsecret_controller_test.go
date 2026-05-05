<<<<<<< HEAD
=======
// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

>>>>>>> tmp-original-05-05-26-00-31
package controller

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("BMCSecret Controller", func() {
	Context("When reconciling a resource", func() {

<<<<<<< HEAD
		It("should successfully reconcile the resource", func() {

=======
		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default",
		}
		bmcsecret := &metalv1alpha1.BMCSecret{}

		BeforeEach(func(ctx SpecContext) {
			By("creating the custom resource for the Kind BMCSecret")
			err := k8sClient.Get(ctx, typeNamespacedName, bmcsecret)
			if err != nil && errors.IsNotFound(err) {
				resource := &metalv1alpha1.BMCSecret{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		AfterEach(func(ctx SpecContext) {
			// TODO(user): Cleanup logic after each test, like removing the resource instance.
			resource := &metalv1alpha1.BMCSecret{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource instance BMCSecret")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})

		It("should successfully reconcile the resource", func(ctx SpecContext) {
			By("Reconciling the created resource")
			controllerReconciler := &BMCSecretReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err := controllerReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespacedName,
			})
			Expect(err).To(Succeed())
>>>>>>> tmp-original-05-05-26-00-31
			// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
			// Example: If you expect a certain status condition after reconciliation, verify it here.
		})
	})
})
