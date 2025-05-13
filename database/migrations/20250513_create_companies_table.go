package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/facades"
)

type M20251205_CreateCompanyTable struct {

}

func (r *M20251205_CreateCompanyTable) Signature() string {
	return "2025_12_05_000000_create_company_table"
}


func (r *M20251205_CreateCompanyTable) Up() error {
	if !facades.Schema().HasTable("companies") {
		return facades.Schema().Create("companies", func(table schema.Blueprint) {
			table.ID("company_id")
			table.String("name") // Tên công ty
			table.String("address") // Địa chỉ công ty
			table.String("phone") // Số điện thoại công ty
			table.String("email") // Email công ty
			table.String("website") // Website công ty
			table.String("description") // Mô tả công ty
			table.String("logo") // Hình đại diện công ty
			table.String("domain")
			table.Timestamps() // Thêm các trường created_at và updated_at
		})
	}
	return nil
}

func (r *M20251205_CreateCompanyTable) Down() error {
	return facades.Schema().DropIfExists("companies")
}