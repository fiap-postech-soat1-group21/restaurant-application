package main

import (
	"flag"
	"log"
	repository "soat1-challenge1/internal/repositories"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type order struct {
	ID        uuid.UUID `gorm:"primaryKey;autoIncrement"`
	Confirmed bool      `gorm:"not null"`
	Paid      bool      `gorm:"not null"`
	StatusID  string    `gorm:"not null"`
	ClientID  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	rollback := flag.Bool("rollback", false, "use when migration should rollback")
	flag.Parse()

	db := repository.NewRepository()

	migration := []*gormigrate.Migration{
		// create orders table
		{
			ID: "20230604",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&order{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&order{})
			},
		},
	}

	m := gormigrate.New(db.Conn, gormigrate.DefaultOptions, migration)

	if *rollback {
		if err := m.RollbackMigration(migration[0]); err != nil {
			log.Fatalf("Could not rollback: %v", err)
		}
		log.Printf("Migration rollback did run successfully")
		return
	}

	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}
	log.Printf("Migration did run successfully")
}
