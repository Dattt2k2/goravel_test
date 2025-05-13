package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateApplicationsTable struct {
}

func (r *M20251205_CreateApplicationsTable) Signature() string {
	return "2025_12_05_000000_create_applications_table"
}

func (r *M20251205_CreateApplicationsTable) Up() error {
	if !facades.Schema().HasTable("applications") {
		return facades.Schema().Create("applications", func(table schema.Blueprint) {
			table.ID("application_id") // Already defines primary key
			table.BigInteger("candidate_id").Unsigned()
			table.String("cv_url")
			table.String("description")
			table.String("status") // Only one status field
			table.Foreign("candidate_id").References("id").On("candidates").CascadeOnDelete()
			table.BigInteger("campaign_id").Unsigned()
			table.Foreign("campaign_id").References("campaign_id").On("recuitement_campaign").CascadeOnDelete()
			table.Timestamps()
		})
	}
	return nil
}

func (r *M20251205_CreateApplicationsTable) Down() error {
	return facades.Schema().DropIfExists("applications")
}
