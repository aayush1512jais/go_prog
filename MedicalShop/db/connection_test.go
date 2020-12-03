package db

import (
	"reflect"
	"testing"

	"github.com/aayush1512jais/go_prog/MedicalShop/apperrors"
)

func TestConnectToPostgres(t *testing.T) {

	// myEnv, err := godotenv.Read()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// uri := myEnv["POSTGRES_DATABASE_URI"]

	//-- Throwing Error while implementing above method which is:
	//-- open .env: The system cannot find the file specified.

	uri := "postgres://postgres:aayush15@localhost/medicalShop?sslmode=disable"
	tests := []struct {
		name        string
		databaseURI string
		want        apperrors.ErrorModel
	}{

		{"Successfully Connected", uri, apperrors.ErrorModel{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := ConnectToPostgres(tt.databaseURI)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToPostgres() got = %v, want %v", got, tt.want)
			}

		})
	}
}
