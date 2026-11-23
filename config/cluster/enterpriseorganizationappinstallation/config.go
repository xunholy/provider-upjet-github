package enterpriseorganizationappinstallation

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure github_enterprise_organization_app_installation resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("github_enterprise_organization_app_installation", func(r *config.Resource) {
		r.Kind = "OrganizationAppInstallation"
		r.ShortGroup = "enterprise"

		r.References["selected_repositories"] = config.Reference{
			TerraformName: "github_repository",
		}
	})
}
