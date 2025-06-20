// Code generated by ent, DO NOT EDIT.

package deposit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the deposit type in the database.
	Label = "deposit"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeFund holds the string denoting the fund edge name in mutations.
	EdgeFund = "fund"
	// EdgeCustomer holds the string denoting the customer edge name in mutations.
	EdgeCustomer = "customer"
	// FundFieldID holds the string denoting the ID field of the Fund.
	FundFieldID = "id"
	// CustomerFieldID holds the string denoting the ID field of the Customer.
	CustomerFieldID = "id"
	// Table holds the table name of the deposit in the database.
	Table = "deposits"
	// FundTable is the table that holds the fund relation/edge. The primary key declared below.
	FundTable = "fund_deposits"
	// FundInverseTable is the table name for the Fund entity.
	// It exists in this package in order to avoid circular dependency with the "fund" package.
	FundInverseTable = "funds"
	// CustomerTable is the table that holds the customer relation/edge. The primary key declared below.
	CustomerTable = "customer_deposits"
	// CustomerInverseTable is the table name for the Customer entity.
	// It exists in this package in order to avoid circular dependency with the "customer" package.
	CustomerInverseTable = "customers"
)

// Columns holds all SQL columns for deposit fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldCreatedAt,
}

var (
	// FundPrimaryKey and FundColumn2 are the table columns denoting the
	// primary key for the fund relation (M2M).
	FundPrimaryKey = []string{"fund_id", "deposit_id"}
	// CustomerPrimaryKey and CustomerColumn2 are the table columns denoting the
	// primary key for the customer relation (M2M).
	CustomerPrimaryKey = []string{"customer_id", "deposit_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Deposit queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmount orders the results by the amount field.
func ByAmount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmount, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByFundCount orders the results by fund count.
func ByFundCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFundStep(), opts...)
	}
}

// ByFund orders the results by fund terms.
func ByFund(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFundStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCustomerCount orders the results by customer count.
func ByCustomerCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCustomerStep(), opts...)
	}
}

// ByCustomer orders the results by customer terms.
func ByCustomer(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCustomerStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newFundStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FundInverseTable, FundFieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, FundTable, FundPrimaryKey...),
	)
}
func newCustomerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CustomerInverseTable, CustomerFieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, CustomerTable, CustomerPrimaryKey...),
	)
}
