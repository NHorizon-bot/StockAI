package config

import "testing"

func TestMySQLConfigDSN(t *testing.T) {
	cfg := MySQLConfig{
		Host:     "db.internal",
		Port:     "3306",
		User:     "app",
		Password: "secret",
		Database: "stockai",
	}

	got := cfg.DSN()
	want := "app:secret@tcp(db.internal:3306)/stockai?charset=utf8mb4&parseTime=True&loc=Local"
	if got != want {
		t.Fatalf("DSN() = %q, want %q", got, want)
	}
}
