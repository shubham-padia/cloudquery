// Code generated by codegen; DO NOT EDIT.

package codegen

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	"github.com/cloudquery/plugin-sdk/schema"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
)

func Teams() *schema.Table {
	return &schema.Table{
		Name:        "heroku_teams",
		Description: "https://devcenter.heroku.com/articles/platform-api-reference#team-attributes",
		Resolver:    fetchTeams,
		Columns: []schema.Column{
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "credit_card_collections",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("CreditCardCollections"),
			},
			{
				Name:     "default",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Default"),
			},
			{
				Name:     "enterprise_account",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EnterpriseAccount"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "identity_provider",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("IdentityProvider"),
			},
			{
				Name:     "membership_limit",
				Type:     schema.TypeFloat,
				Resolver: schema.PathResolver("MembershipLimit"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "provisioned_licenses",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ProvisionedLicenses"),
			},
			{
				Name:     "role",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Role"),
			},
			{
				Name:     "type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Type"),
			},
			{
				Name:     "updated_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("UpdatedAt"),
			},
		},
	}
}

func fetchTeams(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	nextRange := &heroku.ListRange{
		Field: "id",
		Max:   1000,
	}
	// Roundtripper middleware in client/pagination.go
	// sets the nextRange value after each request
	for nextRange.Max != 0 {
		ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
		v, err := c.Heroku.TeamList(ctxWithRange, nextRange)
		if err != nil {
			return errors.WithStack(err)
		}
		res <- v
	}
	return nil
}
