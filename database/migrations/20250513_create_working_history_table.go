package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateWorkingHistoryTable struct {

}

func (r *M20251205_CreateWorkingHistoryTable) Signature() string {
	return "2025_12_05_000000_create_working_history_table"
}

func (r *M20251205_CreateWorkingHistoryTable) Up() error {
    if !facades.Schema().HasTable("working_histories") {
        return facades.Schema().Create("working_histories", func(table schema.Blueprint) {
            table.ID("id") // Use the default ID name instead of a custom one
            // Or if you need a custom ID name:
            // table.BigIncrements("working_history_id") 
            table.BigInteger("candidate_id").Unsigned()
            table.BigInteger("company_id").Unsigned()
            table.String("position")
            table.String("description")
            table.Timestamp("start_date")
            table.Timestamp("end_date").Nullable()
            table.Foreign("candidate_id").References("id").On("candidates").CascadeOnDelete()
            table.Foreign("company_id").References("company_id").On("companies").CascadeOnDelete()
            table.Timestamps()
        })
    }
    return nil
} 

func (r *M20251205_CreateWorkingHistoryTable) Down() error {
	return facades.Schema().DropIfExists("working_histories")
}