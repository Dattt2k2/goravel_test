package database

import (
	"github.com/goravel/framework/contracts/database/schema"
	"github.com/goravel/framework/contracts/database/seeder"

	"goravel/database/migrations"
	"goravel/database/seeders"
)

type Kernel struct {
}

func (kernel Kernel) Migrations() []schema.Migration {
	return []schema.Migration{
		&migrations.M20240915060148CreateUsersTable{},          // First, users table
		&migrations.M20251205_CreateDomainTable{},              // Domain table
		&migrations.M20251205_CreateCompanyTable{},             // Company before recruiters and working history
		&migrations.M20251205_CreateCandidateTable{},           // Candidate before applications
		&migrations.M20251205_CreateRecuiterTable{},            // Recruiters after companies
		&migrations.M20251205_CreateWorkingHistoryTable{},      // Working history after companies and candidates
		&migrations.M20251205_CreateRecuitementCampaignTable{}, // Recruitment campaign before applications
		&migrations.M20251205_CreateApplicationsTable{},        // Applications last as it depends on other tables
	}
}

func (kernel Kernel) Seeders() []seeder.Seeder {
	return []seeder.Seeder{
		&seeders.DatabaseSeeder{},
	}
}
