package helpers

import (
	"ewallet-notification/seeder"

	"github.com/sirupsen/logrus"
)

func RunSeeder() {
	if DB == nil {
		logrus.Fatal("database connection not initialized, call SetupMySQL() first")
	}
	seeder.SeedTemplates(DB)
	logrus.Info("database seeding completed...")
}
