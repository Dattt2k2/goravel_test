package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateRecuiterTable struct {
}

func (r *M20251205_CreateRecuiterTable) Signature() string {
	return "2025_12_05_000000_create_recuiter_table"
}

func (r *M20251205_CreateRecuiterTable) Up() error {
	if !facades.Schema().HasTable("recuiters") {
		return facades.Schema().Create("recuiters", func(table schema.Blueprint) {
			// Use BigInteger for user_id since it's a foreign key
			table.BigInteger("user_id").Unsigned()
			table.Primary("user_id")
			table.String("avatar_url")
			table.Foreign("user_id").References("id").On("users").CascadeOnDelete()
			table.BigInteger("company_id").Unsigned() // Add the company_id column
			table.Foreign("company_id").References("company_id").On("companies").CascadeOnDelete()
			table.Timestamps()
		})
	}
	return nil
}

func (r *M20251205_CreateRecuiterTable) Down() error {
	return facades.Schema().DropIfExists("recuiters")
}
