package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateDomainTable struct {
}

func (r *M20251205_CreateDomainTable) Signature() string {
	return "2025_12_05_000000_create_domain_table"
}

func (r *M20251205_CreateDomainTable) Up() error {
	if !facades.Schema().HasTable("domains") {
		return facades.Schema().Create("domains", func(table schema.Blueprint) {
			table.ID("id")       // This already creates a primary key, no need for table.Primary
			table.String("name") // Add the name field
			table.Unique("name") // Make name unique separately
			table.String("description")
			table.Timestamps()
		})
	}
	return nil
}

func (r *M20251205_CreateDomainTable) Down() error {
	return facades.Schema().DropIfExists("domains")
}
