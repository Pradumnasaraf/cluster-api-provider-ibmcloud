/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"

	"sigs.k8s.io/controller-runtime/pkg/conversion"

	utilconversion "sigs.k8s.io/cluster-api/util/conversion"

	infrav1beta1 "sigs.k8s.io/cluster-api-provider-ibmcloud/api/v1beta1"
)

func (src *IBMPowerVSCluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSCluster)

	return Convert_v1alpha4_IBMPowerVSCluster_To_v1beta1_IBMPowerVSCluster(src, dst, nil)
}

func (dst *IBMPowerVSCluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSCluster)

	return Convert_v1beta1_IBMPowerVSCluster_To_v1alpha4_IBMPowerVSCluster(src, dst, nil)
}

func (src *IBMPowerVSClusterList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSClusterList)

	return Convert_v1alpha4_IBMPowerVSClusterList_To_v1beta1_IBMPowerVSClusterList(src, dst, nil)
}

func (dst *IBMPowerVSClusterList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSClusterList)

	return Convert_v1beta1_IBMPowerVSClusterList_To_v1alpha4_IBMPowerVSClusterList(src, dst, nil)
}

func (src *IBMPowerVSMachine) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSMachine)

	if err := Convert_v1alpha4_IBMPowerVSMachine_To_v1beta1_IBMPowerVSMachine(src, dst, nil); err != nil {
		return err
	}

	if err := Convert_v1alpha4_IBMPowerVSResourceReference_To_v1beta1_IBMPowerVSResourceReference(&src.Spec.Image, dst.Spec.Image, nil); err != nil {
		return err
	}
	return nil
}

func (dst *IBMPowerVSMachine) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSMachine)

	if err := Convert_v1beta1_IBMPowerVSMachine_To_v1alpha4_IBMPowerVSMachine(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	if src.Spec.Image == nil && src.Spec.ImageRef != nil {
		dst.Spec.Image.Name = &src.Spec.ImageRef.Name
	}

	if src.Spec.Image != nil && src.Spec.ImageRef == nil {
		dst.Spec.Image.Name = src.Spec.Image.Name
		dst.Spec.Image.ID = src.Spec.Image.ID
	}

	return nil
}

func (src *IBMPowerVSMachineList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSMachineList)

	return Convert_v1alpha4_IBMPowerVSMachineList_To_v1beta1_IBMPowerVSMachineList(src, dst, nil)
}

func (dst *IBMPowerVSMachineList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSMachineList)

	return Convert_v1beta1_IBMPowerVSMachineList_To_v1alpha4_IBMPowerVSMachineList(src, dst, nil)
}

func (src *IBMPowerVSMachineTemplate) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSMachineTemplate)

	if err := Convert_v1alpha4_IBMPowerVSMachineTemplate_To_v1beta1_IBMPowerVSMachineTemplate(src, dst, nil); err != nil {
		return err
	}
	if err := Convert_v1alpha4_IBMPowerVSResourceReference_To_v1beta1_IBMPowerVSResourceReference(&src.Spec.Template.Spec.Image, dst.Spec.Template.Spec.Image, nil); err != nil {
		return err
	}
	return nil
}

func (dst *IBMPowerVSMachineTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSMachineTemplate)

	if err := Convert_v1beta1_IBMPowerVSMachineTemplate_To_v1alpha4_IBMPowerVSMachineTemplate(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	if src.Spec.Template.Spec.Image == nil && src.Spec.Template.Spec.ImageRef != nil {
		dst.Spec.Template.Spec.Image.Name = &src.Spec.Template.Spec.ImageRef.Name
	}

	if src.Spec.Template.Spec.Image != nil && src.Spec.Template.Spec.ImageRef == nil {
		dst.Spec.Template.Spec.Image.Name = src.Spec.Template.Spec.Image.Name
		dst.Spec.Template.Spec.Image.ID = src.Spec.Template.Spec.Image.ID
	}

	return nil
}

func (src *IBMPowerVSMachineTemplateList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1beta1.IBMPowerVSMachineTemplateList)

	return Convert_v1alpha4_IBMPowerVSMachineTemplateList_To_v1beta1_IBMPowerVSMachineTemplateList(src, dst, nil)
}

func (dst *IBMPowerVSMachineTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1beta1.IBMPowerVSMachineTemplateList)

	return Convert_v1beta1_IBMPowerVSMachineTemplateList_To_v1alpha4_IBMPowerVSMachineTemplateList(src, dst, nil)
}

// Convert_v1beta1_IBMPowerVSMachineStatus_To_v1alpha4_IBMPowerVSMachineStatus is an autogenerated conversion function.
// Requires manual conversion as FailureReason, FailureMessage and Conditions does not exist in v1alpha4 version of IBMPowerVSMachineStatus.
func Convert_v1beta1_IBMPowerVSMachineStatus_To_v1alpha4_IBMPowerVSMachineStatus(in *infrav1beta1.IBMPowerVSMachineStatus, out *IBMPowerVSMachineStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_IBMPowerVSMachineStatus_To_v1alpha4_IBMPowerVSMachineStatus(in, out, s)
}

// Convert_v1beta1_IBMPowerVSMachineSpec_To_v1alpha4_IBMPowerVSMachineSpec is an autogenerated conversion function.
// Requires manual conversion as ImageRef does not exist in v1alpha4 version of IBMPowerVSMachineSpec.
func Convert_v1beta1_IBMPowerVSMachineSpec_To_v1alpha4_IBMPowerVSMachineSpec(in *infrav1beta1.IBMPowerVSMachineSpec, out *IBMPowerVSMachineSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_IBMPowerVSMachineSpec_To_v1alpha4_IBMPowerVSMachineSpec(in, out, s)
}

func Convert_v1alpha4_IBMPowerVSMachineSpec_To_v1beta1_IBMPowerVSMachineSpec(in *IBMPowerVSMachineSpec, out *infrav1beta1.IBMPowerVSMachineSpec, s apiconversion.Scope) error {
	return autoConvert_v1alpha4_IBMPowerVSMachineSpec_To_v1beta1_IBMPowerVSMachineSpec(in, out, s)
}

// Convert_v1beta1_IBMPowerVSMachineTemplateStatus_To_v1alpha4_IBMPowerVSMachineTemplateStatus is an autogenerated conversion function.
func Convert_v1beta1_IBMPowerVSMachineTemplateStatus_To_v1alpha4_IBMPowerVSMachineTemplateStatus(in *infrav1beta1.IBMPowerVSMachineTemplateStatus, out *IBMPowerVSMachineTemplateStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_IBMPowerVSMachineTemplateStatus_To_v1alpha4_IBMPowerVSMachineTemplateStatus(in, out, s)
}

// Convert_v1beta1_IBMPowerVSResourceReference_To_v1alpha4_IBMPowerVSResourceReference is an autogenerated conversion function.
func Convert_v1beta1_IBMPowerVSResourceReference_To_v1alpha4_IBMPowerVSResourceReference(in *infrav1beta1.IBMPowerVSResourceReference, out *IBMPowerVSResourceReference, s apiconversion.Scope) error {
	return autoConvert_v1beta1_IBMPowerVSResourceReference_To_v1alpha4_IBMPowerVSResourceReference(in, out, s)
}
