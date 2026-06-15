/*
Copyright 2026 The Crossplane Authors.

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

// Package kube contains shared helpers for crossplane CLI commands that talk
// to a Kubernetes cluster.
package kube

import "k8s.io/client-go/rest"

// ImpersonationFlags are the kubectl-compatible privilege-elevation flags
// (--as, --as-group, --as-uid). Embed it into a command's Kong flag struct with
// the `embed:""` tag, then call Apply on the command's *rest.Config before
// building its client.
type ImpersonationFlags struct {
	As      string   `name:"as"       help:"Username to impersonate for the operation. User could be a regular user or a service account in a namespace."`
	AsGroup []string `name:"as-group" help:"Group to impersonate for the operation, this flag can be repeated to specify multiple groups." sep:"none"`
	AsUID   string   `name:"as-uid"   help:"UID to impersonate for the operation."`
}

// Apply sets impersonation on the given rest.Config. Unset fields and a nil cfg
// are no-ops, so it is always safe to call.
func (f ImpersonationFlags) Apply(cfg *rest.Config) {
	if cfg == nil {
		return
	}

	if f.As != "" {
		cfg.Impersonate.UserName = f.As
	}

	if f.AsUID != "" {
		cfg.Impersonate.UID = f.AsUID
	}

	if len(f.AsGroup) > 0 {
		cfg.Impersonate.Groups = f.AsGroup
	}
}
