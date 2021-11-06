package repository

import (
	"phone-list/app/model"
	"phone-list/database"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_List(t *testing.T) {
	mockScopeWhere()
	s, err := database.MockDb()
	if err != nil {
		t.Errorf("failed to mock DB %v", err)
	}
	defer s.Conn.Close()

	customer := model.Customer{
		ID:    1,
		Name:  "foo",
		Phone: "bar",
	}

	s.Mock.MatchExpectationsInOrder(false)

	s.Mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "customer" WHERE country = $1 AND state = $2 LIMIT 2 OFFSET 1`)).
		WithArgs("foo", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "phone"}).
			AddRow(customer.ID, customer.Name, customer.Phone))

	repo := NewCustomerRepository(s.DB)
	list, err := repo.List(2, 1, "foo", 1)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(list))

	err = s.Mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestCustomerRepository_Count(t *testing.T) {
	mockScopeWhere()
	s, err := database.MockDb()
	if err != nil {
		t.Errorf("failed to mock DB %v", err)
	}
	defer s.Conn.Close()

	s.Mock.MatchExpectationsInOrder(false)

	s.Mock.
		ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "customer" WHERE country = $1 AND state = $2`)).
		WithArgs("foo", 1).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

	repo := NewCustomerRepository(s.DB)
	count, err := repo.Count("foo", 1)
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}

	assert.Equal(t, 1, count)

	err = s.Mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Failed to meet expectations, got error: %v", err)
	}
}

func TestNewCustomerRepository(t *testing.T) {
	s, err := database.MockDb()
	if err != nil {
		t.Errorf("failed to mock DB %v", err)
	}
	repo := NewCustomerRepository(s.DB).(*customerRepository)

	assert.Equal(t, s.DB, repo.db)
}
