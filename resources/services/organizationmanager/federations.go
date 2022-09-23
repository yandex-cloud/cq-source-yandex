// Code generated by codegen; DO NOT EDIT.

package organizationmanager

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Federations() *schema.Table {
	return &schema.Table{
		Name:      "yandex_organizationmanager_federations",
		Resolver:  fetchFederations,
		Multiplex: client.MultiplexBy(client.Organizations),
		Columns: []schema.Column{
			{
				Name:        "id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Id"),
				Description: `Resource ID`,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "organization_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("OrganizationId"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
			},
			{
				Name:     "cookie_max_age",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CookieMaxAge"),
			},
			{
				Name:     "auto_create_account_on_login",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoCreateAccountOnLogin"),
			},
			{
				Name:     "issuer",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Issuer"),
			},
			{
				Name:     "sso_binding",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SsoBinding"),
			},
			{
				Name:     "sso_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SsoUrl"),
			},
			{
				Name:     "security_settings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SecuritySettings"),
			},
			{
				Name:     "case_insensitive_name_ids",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CaseInsensitiveNameIds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
		},
	}
}
