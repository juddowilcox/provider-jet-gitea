/*
Copyright 2021 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/terrajet/pkg/controller"

	hook "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/git/hook"
	org "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/gitea/org"
	repository "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/gitea/repository"
	team "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/gitea/team"
	user "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/gitea/user"
	app "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/oauth2/app"
	providerconfig "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/providerconfig"
	key "github.com/crossplane-contrib/provider-jet-gitea/internal/controller/public/key"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hook.Setup,
		org.Setup,
		repository.Setup,
		team.Setup,
		user.Setup,
		app.Setup,
		providerconfig.Setup,
		key.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
