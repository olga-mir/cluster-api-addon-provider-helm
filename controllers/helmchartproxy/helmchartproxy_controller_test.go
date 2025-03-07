/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package helmchartproxy

import (
	"testing"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	addonsv1alpha1 "sigs.k8s.io/cluster-api-addon-provider-helm/api/v1alpha1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/cluster-api/util/conditions"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ reconcile.Reconciler = &HelmChartProxyReconciler{}

var (
	defaultProxy = &addonsv1alpha1.HelmChartProxy{
		TypeMeta: metav1.TypeMeta{
			APIVersion: addonsv1alpha1.GroupVersion.String(),
			Kind:       "HelmChartProxy",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-hcp",
			Namespace: "test-namespace",
		},
		Spec: addonsv1alpha1.HelmChartProxySpec{
			ClusterSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"test-label": "test-value",
				},
			},
			ReleaseName:      "test-release-name",
			ChartName:        "test-chart-name",
			RepoURL:          "https://test-repo-url",
			ReleaseNamespace: "test-release-namespace",
			Version:          "test-version",
			ValuesTemplate:   "apiServerPort: {{ .Cluster.spec.clusterNetwork.apiServerPort }}",
			Options:          &addonsv1alpha1.HelmOptions{},
		},
	}

	cluster1 = &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.GroupVersion.String(),
			Kind:       "Cluster",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cluster-1",
			Namespace: "test-namespace",
			Labels: map[string]string{
				"test-label": "test-value",
			},
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: ptr.To(int32(1234)),
			},
		},
	}

	cluster2 = &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.GroupVersion.String(),
			Kind:       "Cluster",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cluster-2",
			Namespace: "test-namespace",
			Labels: map[string]string{
				"test-label":  "test-value",
				"other-label": "other-value",
			},
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: ptr.To(int32(5678)),
			},
		},
	}

	cluster3 = &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.GroupVersion.String(),
			Kind:       "Cluster",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cluster-3",
			Namespace: "test-namespace",
			Labels: map[string]string{
				"other-label": "other-value",
			},
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: ptr.To(int32(6443)),
			},
		},
	}

	cluster4 = &clusterv1.Cluster{
		TypeMeta: metav1.TypeMeta{
			APIVersion: clusterv1.GroupVersion.String(),
			Kind:       "Cluster",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-cluster-4",
			Namespace: "other-namespace",
			Labels: map[string]string{
				"other-label": "other-value",
			},
		},
		Spec: clusterv1.ClusterSpec{
			ClusterNetwork: &clusterv1.ClusterNetwork{
				APIServerPort: ptr.To(int32(6443)),
			},
		},
	}

	hrpReady1 = &addonsv1alpha1.HelmReleaseProxy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-hrp-1",
			Namespace: "test-namespace",
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion:         addonsv1alpha1.GroupVersion.String(),
					Kind:               "HelmChartProxy",
					Name:               "test-hcp",
					Controller:         ptr.To(true),
					BlockOwnerDeletion: ptr.To(true),
				},
			},
			Labels: map[string]string{
				clusterv1.ClusterNameLabel:             "test-cluster-1",
				addonsv1alpha1.HelmChartProxyLabelName: "test-hcp",
			},
		},
		Spec: addonsv1alpha1.HelmReleaseProxySpec{
			ClusterRef: corev1.ObjectReference{
				APIVersion: clusterv1.GroupVersion.String(),
				Kind:       "Cluster",
				Name:       "test-cluster-1",
				Namespace:  "test-namespace",
			},
			ReleaseName:      "test-release-name",
			ChartName:        "test-chart-name",
			RepoURL:          "https://test-repo-url",
			ReleaseNamespace: "test-release-namespace",
			Version:          "test-version",
			Values:           "apiServerPort: 1234",
			Options:          &addonsv1alpha1.HelmOptions{},
		},
		Status: addonsv1alpha1.HelmReleaseProxyStatus{
			Conditions: []clusterv1.Condition{
				{
					Type:   clusterv1.ReadyCondition,
					Status: corev1.ConditionTrue,
				},
			},
		},
	}

	hrpReady2 = &addonsv1alpha1.HelmReleaseProxy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-hrp-2",
			Namespace: "test-namespace",
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion:         addonsv1alpha1.GroupVersion.String(),
					Kind:               "HelmChartProxy",
					Name:               "test-hcp",
					Controller:         ptr.To(true),
					BlockOwnerDeletion: ptr.To(true),
				},
			},
			Labels: map[string]string{
				clusterv1.ClusterNameLabel:             "test-cluster-2",
				addonsv1alpha1.HelmChartProxyLabelName: "test-hcp",
			},
		},
		Spec: addonsv1alpha1.HelmReleaseProxySpec{
			ClusterRef: corev1.ObjectReference{
				APIVersion: clusterv1.GroupVersion.String(),
				Kind:       "Cluster",
				Name:       "test-cluster-2",
				Namespace:  "test-namespace",
			},
			ReleaseName:      "test-release-name",
			ChartName:        "test-chart-name",
			RepoURL:          "https://test-repo-url",
			ReleaseNamespace: "test-release-namespace",
			Version:          "test-version",
			Values:           "apiServerPort: 5678",
			Options:          &addonsv1alpha1.HelmOptions{},
		},
		Status: addonsv1alpha1.HelmReleaseProxyStatus{
			Conditions: []clusterv1.Condition{
				{
					Type:   clusterv1.ReadyCondition,
					Status: corev1.ConditionTrue,
				},
			},
		},
	}
)

func TestReconcileNormal(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name           string
		helmChartProxy *addonsv1alpha1.HelmChartProxy
		objects        []client.Object
		expect         func(g *WithT, c client.Client, hcp *addonsv1alpha1.HelmChartProxy)
		expectedError  string
	}{
		{
			name:           "successfully select clusters and install HelmReleaseProxies",
			helmChartProxy: defaultProxy,
			objects:        []client.Object{cluster1, cluster2, cluster3, cluster4},
			expect: func(g *WithT, c client.Client, hcp *addonsv1alpha1.HelmChartProxy) {
				g.Expect(hcp.Status.MatchingClusters).To(BeEquivalentTo([]corev1.ObjectReference{
					{
						APIVersion: clusterv1.GroupVersion.String(),
						Kind:       "Cluster",
						Name:       "test-cluster-1",
						Namespace:  "test-namespace",
					},
					{
						APIVersion: clusterv1.GroupVersion.String(),
						Kind:       "Cluster",
						Name:       "test-cluster-2",
						Namespace:  "test-namespace",
					},
				}))
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				// This is false as the HelmReleaseProxies won't be ready until the HelmReleaseProxy controller runs.
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxiesReadyCondition)).To(BeFalse())
			},
			expectedError: "",
		},
		{
			name:           "mark HelmChartProxy as ready once HelmReleaseProxies ready conditions are true",
			helmChartProxy: defaultProxy,
			objects:        []client.Object{cluster1, cluster2, hrpReady1, hrpReady2},
			expect: func(g *WithT, c client.Client, hcp *addonsv1alpha1.HelmChartProxy) {
				g.Expect(hcp.Status.MatchingClusters).To(BeEquivalentTo([]corev1.ObjectReference{
					{
						APIVersion: clusterv1.GroupVersion.String(),
						Kind:       "Cluster",
						Name:       "test-cluster-1",
						Namespace:  "test-namespace",
					},
					{
						APIVersion: clusterv1.GroupVersion.String(),
						Kind:       "Cluster",
						Name:       "test-cluster-2",
						Namespace:  "test-namespace",
					},
				}))
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxiesReadyCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, addonsv1alpha1.HelmReleaseProxiesReadyCondition)).To(BeTrue())
				g.Expect(conditions.Has(hcp, clusterv1.ReadyCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, clusterv1.ReadyCondition)).To(BeTrue())
			},
			expectedError: "",
		},
		{
			name:           "successfully delete orphaned HelmReleaseProxies",
			helmChartProxy: defaultProxy,
			objects:        []client.Object{hrpReady1, hrpReady2},
			expect: func(g *WithT, c client.Client, hcp *addonsv1alpha1.HelmChartProxy) {
				g.Expect(hcp.Status.MatchingClusters).To(BeEmpty())
				g.Expect(c.Get(ctx, client.ObjectKey{Namespace: hrpReady1.Namespace, Name: hrpReady1.Name}, &addonsv1alpha1.HelmReleaseProxy{})).ToNot(Succeed())
				g.Expect(c.Get(ctx, client.ObjectKey{Namespace: hrpReady2.Namespace, Name: hrpReady2.Name}, &addonsv1alpha1.HelmReleaseProxy{})).ToNot(Succeed())

				// Vacuously true as there are no HRPs
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, addonsv1alpha1.HelmReleaseProxySpecsUpToDateCondition)).To(BeTrue())
				g.Expect(conditions.Has(hcp, addonsv1alpha1.HelmReleaseProxiesReadyCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, addonsv1alpha1.HelmReleaseProxiesReadyCondition)).To(BeTrue())
				g.Expect(conditions.Has(hcp, clusterv1.ReadyCondition)).To(BeTrue())
				g.Expect(conditions.IsTrue(hcp, clusterv1.ReadyCondition)).To(BeTrue())
			},
			expectedError: "",
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			g := NewWithT(t)
			t.Parallel()
			request := reconcile.Request{
				NamespacedName: util.ObjectKey(tc.helmChartProxy),
			}

			tc.objects = append(tc.objects, tc.helmChartProxy)
			r := &HelmChartProxyReconciler{
				Client: fake.NewClientBuilder().
					WithScheme(fakeScheme).
					WithObjects(tc.objects...).
					WithStatusSubresource(&addonsv1alpha1.HelmChartProxy{}).
					WithStatusSubresource(&addonsv1alpha1.HelmReleaseProxy{}).
					Build(),
			}
			result, err := r.Reconcile(ctx, request)

			if tc.expectedError != "" {
				g.Expect(err).To(HaveOccurred())
				g.Expect(err).To(MatchError(tc.expectedError), err.Error())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(result).To(Equal(reconcile.Result{}))

				hcp := &addonsv1alpha1.HelmChartProxy{}
				g.Expect(r.Client.Get(ctx, request.NamespacedName, hcp)).To(Succeed())

				tc.expect(g, r.Client, hcp)
			}
		})
	}
}
