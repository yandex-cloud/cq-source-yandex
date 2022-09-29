// Code generated by codegen; DO NOT EDIT.

package serverless

import (
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/yandex-cloud/cq-provider-yandex/client"
)

func Functions() *schema.Table {
	return &schema.Table{
		Name:      "yandex_serverless_functions",
		Resolver:  fetchFunctions,
		Multiplex: client.MultiplexBy(client.Folders),
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
				Name:     "folder_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FolderId"),
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: client.ResolveProtoTimestamp("CreatedAt"),
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
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "log_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LogGroupId"),
			},
			{
				Name:     "http_invoke_url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("HttpInvokeUrl"),
			},
			{
				Name:     "status",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status"),
			},
		},
	}
}
