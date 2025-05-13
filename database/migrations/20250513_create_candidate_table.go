package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateCandidateTable struct {
}

func (r *M20251205_CreateCandidateTable) Signature() string {
	return "2025_12_05_000000_create_candidate_table"
}

func (r *M20251205_CreateCandidateTable) Up() error {
	return facades.Schema().Create("candidates", func(table schema.Blueprint) {
		// Sử dụng BigInteger và đánh dấu là Unsigned cho cột id
		table.BigInteger("id").Unsigned()
		// Thiết lập id là primary key
		table.Primary("id")
		// Liên kết id của candidates với id của bảng users
		table.Foreign("id").References("id").On("users").CascadeOnDelete()
		// Các trường dữ liệu khác
		table.Date("dob")
		table.String("school")
		table.Date("year_enrollment")
		table.Date("year_graduation")
		table.Boolean("is_graduated").Default(false)
		table.Integer("experience_years").Default(0)
		table.Timestamps() // Thêm các trường created_at và updated_at
	})
	
}

// Down Reverse the migrations.
func (r *M20251205_CreateCandidateTable) Down() error {
	return facades.Schema().DropIfExists("candidates")
}