package scope

import (
	"phone-list/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestWherePhone_ValidState(t *testing.T) {
	mockGetValidCodeRegex("foo")
	mockGetValidPhoneRegex("bar")
	s, _ := database.MockDb()
	var count int64

	s.Mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "foo" WHERE phone REGEXP $1`)).
		WithArgs("bar").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	s.DB.Table("foo").Scopes(WherePhone("", 1)).Count(&count)
}

func TestWherePhone_InvalidState(t *testing.T) {
	mockGetValidCodeRegex("foo")
	mockGetValidPhoneRegex("bar")
	s, _ := database.MockDb()
	var count int64

	s.Mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "foo" WHERE phone REGEXP $1 AND phone NOT REGEXP $2`)).
		WithArgs("foo", "bar").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	s.DB.Table("foo").Scopes(WherePhone("foo", 2)).Count(&count)
}
