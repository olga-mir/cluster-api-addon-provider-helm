//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credentials) DeepCopyInto(out *Credentials) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credentials.
func (in *Credentials) DeepCopy() *Credentials {
	if in == nil {
		return nil
	}
	out := new(Credentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartProxy) DeepCopyInto(out *HelmChartProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartProxy.
func (in *HelmChartProxy) DeepCopy() *HelmChartProxy {
	if in == nil {
		return nil
	}
	out := new(HelmChartProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChartProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartProxyList) DeepCopyInto(out *HelmChartProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmChartProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartProxyList.
func (in *HelmChartProxyList) DeepCopy() *HelmChartProxyList {
	if in == nil {
		return nil
	}
	out := new(HelmChartProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmChartProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartProxySpec) DeepCopyInto(out *HelmChartProxySpec) {
	*out = *in
	in.ClusterSelector.DeepCopyInto(&out.ClusterSelector)
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(HelmOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(Credentials)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartProxySpec.
func (in *HelmChartProxySpec) DeepCopy() *HelmChartProxySpec {
	if in == nil {
		return nil
	}
	out := new(HelmChartProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmChartProxyStatus) DeepCopyInto(out *HelmChartProxyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MatchingClusters != nil {
		in, out := &in.MatchingClusters, &out.MatchingClusters
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmChartProxyStatus.
func (in *HelmChartProxyStatus) DeepCopy() *HelmChartProxyStatus {
	if in == nil {
		return nil
	}
	out := new(HelmChartProxyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInstallOptions) DeepCopyInto(out *HelmInstallOptions) {
	*out = *in
	if in.CreateNamespace != nil {
		in, out := &in.CreateNamespace, &out.CreateNamespace
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInstallOptions.
func (in *HelmInstallOptions) DeepCopy() *HelmInstallOptions {
	if in == nil {
		return nil
	}
	out := new(HelmInstallOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmOptions) DeepCopyInto(out *HelmOptions) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		*out = new(HelmInstallOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Upgrade != nil {
		in, out := &in.Upgrade, &out.Upgrade
		*out = new(HelmUpgradeOptions)
		**out = **in
	}
	if in.Uninstall != nil {
		in, out := &in.Uninstall, &out.Uninstall
		*out = new(HelmUninstallOptions)
		**out = **in
	}
	if in.EnableClientCache != nil {
		in, out := &in.EnableClientCache, &out.EnableClientCache
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmOptions.
func (in *HelmOptions) DeepCopy() *HelmOptions {
	if in == nil {
		return nil
	}
	out := new(HelmOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseProxy) DeepCopyInto(out *HelmReleaseProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseProxy.
func (in *HelmReleaseProxy) DeepCopy() *HelmReleaseProxy {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmReleaseProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseProxyList) DeepCopyInto(out *HelmReleaseProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HelmReleaseProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseProxyList.
func (in *HelmReleaseProxyList) DeepCopy() *HelmReleaseProxyList {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HelmReleaseProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseProxySpec) DeepCopyInto(out *HelmReleaseProxySpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(HelmOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(Credentials)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseProxySpec.
func (in *HelmReleaseProxySpec) DeepCopy() *HelmReleaseProxySpec {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmReleaseProxyStatus) DeepCopyInto(out *HelmReleaseProxyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmReleaseProxyStatus.
func (in *HelmReleaseProxyStatus) DeepCopy() *HelmReleaseProxyStatus {
	if in == nil {
		return nil
	}
	out := new(HelmReleaseProxyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmUninstallOptions) DeepCopyInto(out *HelmUninstallOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmUninstallOptions.
func (in *HelmUninstallOptions) DeepCopy() *HelmUninstallOptions {
	if in == nil {
		return nil
	}
	out := new(HelmUninstallOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmUpgradeOptions) DeepCopyInto(out *HelmUpgradeOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmUpgradeOptions.
func (in *HelmUpgradeOptions) DeepCopy() *HelmUpgradeOptions {
	if in == nil {
		return nil
	}
	out := new(HelmUpgradeOptions)
	in.DeepCopyInto(out)
	return out
}
