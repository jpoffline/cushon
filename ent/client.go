// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"cushon/ent/migrate"

	"cushon/ent/customer"
	"cushon/ent/deposit"
	"cushon/ent/fund"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Customer is the client for interacting with the Customer builders.
	Customer *CustomerClient
	// Deposit is the client for interacting with the Deposit builders.
	Deposit *DepositClient
	// Fund is the client for interacting with the Fund builders.
	Fund *FundClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Customer = NewCustomerClient(c.config)
	c.Deposit = NewDepositClient(c.config)
	c.Fund = NewFundClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Customer: NewCustomerClient(cfg),
		Deposit:  NewDepositClient(cfg),
		Fund:     NewFundClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Customer: NewCustomerClient(cfg),
		Deposit:  NewDepositClient(cfg),
		Fund:     NewFundClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Customer.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Customer.Use(hooks...)
	c.Deposit.Use(hooks...)
	c.Fund.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Customer.Intercept(interceptors...)
	c.Deposit.Intercept(interceptors...)
	c.Fund.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CustomerMutation:
		return c.Customer.mutate(ctx, m)
	case *DepositMutation:
		return c.Deposit.mutate(ctx, m)
	case *FundMutation:
		return c.Fund.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// CustomerClient is a client for the Customer schema.
type CustomerClient struct {
	config
}

// NewCustomerClient returns a client for the Customer from the given config.
func NewCustomerClient(c config) *CustomerClient {
	return &CustomerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customer.Hooks(f(g(h())))`.
func (c *CustomerClient) Use(hooks ...Hook) {
	c.hooks.Customer = append(c.hooks.Customer, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `customer.Intercept(f(g(h())))`.
func (c *CustomerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Customer = append(c.inters.Customer, interceptors...)
}

// Create returns a builder for creating a Customer entity.
func (c *CustomerClient) Create() *CustomerCreate {
	mutation := newCustomerMutation(c.config, OpCreate)
	return &CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Customer entities.
func (c *CustomerClient) CreateBulk(builders ...*CustomerCreate) *CustomerCreateBulk {
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CustomerClient) MapCreateBulk(slice any, setFunc func(*CustomerCreate, int)) *CustomerCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CustomerCreateBulk{err: fmt.Errorf("calling to CustomerClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CustomerCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CustomerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Customer.
func (c *CustomerClient) Update() *CustomerUpdate {
	mutation := newCustomerMutation(c.config, OpUpdate)
	return &CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomerClient) UpdateOne(cu *Customer) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomer(cu))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomerClient) UpdateOneID(id uuid.UUID) *CustomerUpdateOne {
	mutation := newCustomerMutation(c.config, OpUpdateOne, withCustomerID(id))
	return &CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Customer.
func (c *CustomerClient) Delete() *CustomerDelete {
	mutation := newCustomerMutation(c.config, OpDelete)
	return &CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CustomerClient) DeleteOne(cu *Customer) *CustomerDeleteOne {
	return c.DeleteOneID(cu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CustomerClient) DeleteOneID(id uuid.UUID) *CustomerDeleteOne {
	builder := c.Delete().Where(customer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomerDeleteOne{builder}
}

// Query returns a query builder for Customer.
func (c *CustomerClient) Query() *CustomerQuery {
	return &CustomerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCustomer},
		inters: c.Interceptors(),
	}
}

// Get returns a Customer entity by its id.
func (c *CustomerClient) Get(ctx context.Context, id uuid.UUID) (*Customer, error) {
	return c.Query().Where(customer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomerClient) GetX(ctx context.Context, id uuid.UUID) *Customer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeposits queries the deposits edge of a Customer.
func (c *CustomerClient) QueryDeposits(cu *Customer) *DepositQuery {
	query := (&DepositClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := cu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, id),
			sqlgraph.To(deposit.Table, deposit.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, customer.DepositsTable, customer.DepositsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(cu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CustomerClient) Hooks() []Hook {
	return c.hooks.Customer
}

// Interceptors returns the client interceptors.
func (c *CustomerClient) Interceptors() []Interceptor {
	return c.inters.Customer
}

func (c *CustomerClient) mutate(ctx context.Context, m *CustomerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CustomerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CustomerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CustomerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CustomerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Customer mutation op: %q", m.Op())
	}
}

// DepositClient is a client for the Deposit schema.
type DepositClient struct {
	config
}

// NewDepositClient returns a client for the Deposit from the given config.
func NewDepositClient(c config) *DepositClient {
	return &DepositClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `deposit.Hooks(f(g(h())))`.
func (c *DepositClient) Use(hooks ...Hook) {
	c.hooks.Deposit = append(c.hooks.Deposit, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `deposit.Intercept(f(g(h())))`.
func (c *DepositClient) Intercept(interceptors ...Interceptor) {
	c.inters.Deposit = append(c.inters.Deposit, interceptors...)
}

// Create returns a builder for creating a Deposit entity.
func (c *DepositClient) Create() *DepositCreate {
	mutation := newDepositMutation(c.config, OpCreate)
	return &DepositCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Deposit entities.
func (c *DepositClient) CreateBulk(builders ...*DepositCreate) *DepositCreateBulk {
	return &DepositCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DepositClient) MapCreateBulk(slice any, setFunc func(*DepositCreate, int)) *DepositCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DepositCreateBulk{err: fmt.Errorf("calling to DepositClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DepositCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DepositCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Deposit.
func (c *DepositClient) Update() *DepositUpdate {
	mutation := newDepositMutation(c.config, OpUpdate)
	return &DepositUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepositClient) UpdateOne(d *Deposit) *DepositUpdateOne {
	mutation := newDepositMutation(c.config, OpUpdateOne, withDeposit(d))
	return &DepositUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepositClient) UpdateOneID(id uuid.UUID) *DepositUpdateOne {
	mutation := newDepositMutation(c.config, OpUpdateOne, withDepositID(id))
	return &DepositUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Deposit.
func (c *DepositClient) Delete() *DepositDelete {
	mutation := newDepositMutation(c.config, OpDelete)
	return &DepositDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DepositClient) DeleteOne(d *Deposit) *DepositDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DepositClient) DeleteOneID(id uuid.UUID) *DepositDeleteOne {
	builder := c.Delete().Where(deposit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepositDeleteOne{builder}
}

// Query returns a query builder for Deposit.
func (c *DepositClient) Query() *DepositQuery {
	return &DepositQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDeposit},
		inters: c.Interceptors(),
	}
}

// Get returns a Deposit entity by its id.
func (c *DepositClient) Get(ctx context.Context, id uuid.UUID) (*Deposit, error) {
	return c.Query().Where(deposit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepositClient) GetX(ctx context.Context, id uuid.UUID) *Deposit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFund queries the fund edge of a Deposit.
func (c *DepositClient) QueryFund(d *Deposit) *FundQuery {
	query := (&FundClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deposit.Table, deposit.FieldID, id),
			sqlgraph.To(fund.Table, fund.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, deposit.FundTable, deposit.FundPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCustomer queries the customer edge of a Deposit.
func (c *DepositClient) QueryCustomer(d *Deposit) *CustomerQuery {
	query := (&CustomerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(deposit.Table, deposit.FieldID, id),
			sqlgraph.To(customer.Table, customer.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, deposit.CustomerTable, deposit.CustomerPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepositClient) Hooks() []Hook {
	return c.hooks.Deposit
}

// Interceptors returns the client interceptors.
func (c *DepositClient) Interceptors() []Interceptor {
	return c.inters.Deposit
}

func (c *DepositClient) mutate(ctx context.Context, m *DepositMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DepositCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DepositUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DepositUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DepositDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Deposit mutation op: %q", m.Op())
	}
}

// FundClient is a client for the Fund schema.
type FundClient struct {
	config
}

// NewFundClient returns a client for the Fund from the given config.
func NewFundClient(c config) *FundClient {
	return &FundClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fund.Hooks(f(g(h())))`.
func (c *FundClient) Use(hooks ...Hook) {
	c.hooks.Fund = append(c.hooks.Fund, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `fund.Intercept(f(g(h())))`.
func (c *FundClient) Intercept(interceptors ...Interceptor) {
	c.inters.Fund = append(c.inters.Fund, interceptors...)
}

// Create returns a builder for creating a Fund entity.
func (c *FundClient) Create() *FundCreate {
	mutation := newFundMutation(c.config, OpCreate)
	return &FundCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Fund entities.
func (c *FundClient) CreateBulk(builders ...*FundCreate) *FundCreateBulk {
	return &FundCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FundClient) MapCreateBulk(slice any, setFunc func(*FundCreate, int)) *FundCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FundCreateBulk{err: fmt.Errorf("calling to FundClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FundCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FundCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Fund.
func (c *FundClient) Update() *FundUpdate {
	mutation := newFundMutation(c.config, OpUpdate)
	return &FundUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FundClient) UpdateOne(f *Fund) *FundUpdateOne {
	mutation := newFundMutation(c.config, OpUpdateOne, withFund(f))
	return &FundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FundClient) UpdateOneID(id uuid.UUID) *FundUpdateOne {
	mutation := newFundMutation(c.config, OpUpdateOne, withFundID(id))
	return &FundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Fund.
func (c *FundClient) Delete() *FundDelete {
	mutation := newFundMutation(c.config, OpDelete)
	return &FundDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FundClient) DeleteOne(f *Fund) *FundDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FundClient) DeleteOneID(id uuid.UUID) *FundDeleteOne {
	builder := c.Delete().Where(fund.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FundDeleteOne{builder}
}

// Query returns a query builder for Fund.
func (c *FundClient) Query() *FundQuery {
	return &FundQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFund},
		inters: c.Interceptors(),
	}
}

// Get returns a Fund entity by its id.
func (c *FundClient) Get(ctx context.Context, id uuid.UUID) (*Fund, error) {
	return c.Query().Where(fund.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FundClient) GetX(ctx context.Context, id uuid.UUID) *Fund {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDeposits queries the deposits edge of a Fund.
func (c *FundClient) QueryDeposits(f *Fund) *DepositQuery {
	query := (&DepositClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(fund.Table, fund.FieldID, id),
			sqlgraph.To(deposit.Table, deposit.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, fund.DepositsTable, fund.DepositsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FundClient) Hooks() []Hook {
	return c.hooks.Fund
}

// Interceptors returns the client interceptors.
func (c *FundClient) Interceptors() []Interceptor {
	return c.inters.Fund
}

func (c *FundClient) mutate(ctx context.Context, m *FundMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FundCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FundUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FundUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FundDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Fund mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Customer, Deposit, Fund []ent.Hook
	}
	inters struct {
		Customer, Deposit, Fund []ent.Interceptor
	}
)
