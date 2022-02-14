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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1beta3

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.ComponentsSpec": schema_pkg_apis_syndesis_v1beta3_ComponentsSpec(ref),
		"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.Syndesis":       schema_pkg_apis_syndesis_v1beta3_Syndesis(ref),
		"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisSpec":   schema_pkg_apis_syndesis_v1beta3_SyndesisSpec(ref),
		"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisStatus": schema_pkg_apis_syndesis_v1beta3_SyndesisStatus(ref),
	}
}

func schema_pkg_apis_syndesis_v1beta3_ComponentsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"oauth": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.OauthConfiguration"),
						},
					},
					"server": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.ServerConfiguration"),
						},
					},
					"meta": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.MetaConfiguration"),
						},
					},
					"database": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.DatabaseConfiguration"),
						},
					},
					"prometheus": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.PrometheusConfiguration"),
						},
					},
					"grafana": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.GrafanaConfiguration"),
						},
					},
					"upgrade": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.UpgradeConfiguration"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.DatabaseConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.GrafanaConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.MetaConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.OauthConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.PrometheusConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.ServerConfiguration", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.UpgradeConfiguration"},
	}
}

func schema_pkg_apis_syndesis_v1beta3_Syndesis(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Syndesis is the Schema for the Syndeses API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisSpec", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SyndesisStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_syndesis_v1beta3_SyndesisSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SyndesisSpec defines the desired state of Syndesis",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"backup": {
						SchemaProps: spec.SchemaProps{
							Description: "Schedule backup",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.BackupConfig"),
						},
					},
					"routeHostname": {
						SchemaProps: spec.SchemaProps{
							Description: "The external hostname to access Syndesis",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"components": {
						SchemaProps: spec.SchemaProps{
							Description: "Components is used to configure all the core components of Syndesis",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.ComponentsSpec"),
						},
					},
					"addons": {
						SchemaProps: spec.SchemaProps{
							Description: "Optional add on features that can be enabled.",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.AddonsSpec"),
						},
					},
					"forceMigration": {
						SchemaProps: spec.SchemaProps{
							Description: "Force migration of CR to new version",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"infraScheduling": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration of Affinity and Toleration for infrastructure component pods",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SchedulingSpec"),
						},
					},
					"integrationScheduling": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration of Affinity and Toleration for integrations pods",
							Default:     map[string]interface{}{},
							Ref:         ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SchedulingSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.AddonsSpec", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.BackupConfig", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.ComponentsSpec", "github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.SchedulingSpec"},
	}
}

func schema_pkg_apis_syndesis_v1beta3_SyndesisStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "SyndesisStatus defines the observed state of Syndesis",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "The phase the operator has reached, eg. INSTALLED, STARTING",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"upgradeAttempts": {
						SchemaProps: spec.SchemaProps{
							Description: "A record of the number of times and upgrade has been attempted",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"lastUpgradeFailure": {
						SchemaProps: spec.SchemaProps{
							Description: "A record of the time of the last upgrade failure",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"forceUpgrade": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"boolean"},
							Format: "",
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "Reason if the installation or upgrade failed",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"description": {
						SchemaProps: spec.SchemaProps{
							Description: "Current description of where the installation or upgrade has reached",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Description: "The currently installed version of Syndesis",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"targetVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"backup": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.BackupStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/syndesisio/syndesis/install/operator/pkg/apis/syndesis/v1beta3.BackupStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
