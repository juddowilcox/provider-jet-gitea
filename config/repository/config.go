package repository

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("gitea_repository", func(r *config.Resource) {

        // we need to override the default group that terrajet generated for
        // this resource, which would be "gitea"  
        r.ShortGroup = "repository"
    })
}
