package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateRecuitementCampaignTable struct {
}

func (r *M20251205_CreateRecuitementCampaignTable) Signature() string {
	return "2025_12_05_000000_create_recuitement_campaign_table"
}

func (r *M20251205_CreateRecuitementCampaignTable) Up() error {
	if !facades.Schema().HasTable("recuitement_campaign") {
		return facades.Schema().Create("recuitement_campaign", func(table schema.Blueprint) {
			table.ID("campaign_id")                    // This already creates a primary key
			table.BigInteger("recuiter_id").Unsigned() // Use BigInteger instead of ID
			table.Foreign("recuiter_id").References("user_id").On("recuiters").CascadeOnDelete()
			table.String("name")
			table.String("description")
			table.Decimal("salary_from")
			table.Decimal("salary_to")
			table.String("currency")
			table.String("domain")
			table.Foreign("domain").References("id").On("domains").CascadeOnDelete()
			table.Integer("experience_years")
			table.Boolean("urgent").Default(false)
			table.Date("start_date")
			table.Date("end_date")
			table.BigInteger("company_id").Unsigned()
			table.Foreign("company_id").References("company_id").On("companies").CascadeOnDelete()
			table.Timestamps()
		})
	}
	return nil
}

func (r *M20251205_CreateRecuitementCampaignTable) Down() error {
	return facades.Schema().DropIfExists("recuitement_campaign") // Fix table name consistency
}
