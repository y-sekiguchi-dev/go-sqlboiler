// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Child is an object representing the database table.
type Child struct {
	PersonID  int64     `boil:"person_id" json:"person_id" toml:"person_id" yaml:"person_id"`
	SubNo     int8      `boil:"sub_no" json:"sub_no" toml:"sub_no" yaml:"sub_no"`
	FirstName string    `boil:"first_name" json:"first_name" toml:"first_name" yaml:"first_name"`
	LastName  string    `boil:"last_name" json:"last_name" toml:"last_name" yaml:"last_name"`
	Birthday  time.Time `boil:"birthday" json:"birthday" toml:"birthday" yaml:"birthday"`

	R *childR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L childL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ChildColumns = struct {
	PersonID  string
	SubNo     string
	FirstName string
	LastName  string
	Birthday  string
}{
	PersonID:  "person_id",
	SubNo:     "sub_no",
	FirstName: "first_name",
	LastName:  "last_name",
	Birthday:  "birthday",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperint8 struct{ field string }

func (w whereHelperint8) EQ(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint8) NEQ(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint8) LT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint8) LTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint8) GT(x int8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint8) GTE(x int8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint8) IN(slice []int8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ChildWhere = struct {
	PersonID  whereHelperint64
	SubNo     whereHelperint8
	FirstName whereHelperstring
	LastName  whereHelperstring
	Birthday  whereHelpertime_Time
}{
	PersonID:  whereHelperint64{field: "`children`.`person_id`"},
	SubNo:     whereHelperint8{field: "`children`.`sub_no`"},
	FirstName: whereHelperstring{field: "`children`.`first_name`"},
	LastName:  whereHelperstring{field: "`children`.`last_name`"},
	Birthday:  whereHelpertime_Time{field: "`children`.`birthday`"},
}

// ChildRels is where relationship names are stored.
var ChildRels = struct {
	Person string
}{
	Person: "Person",
}

// childR is where relationships are stored.
type childR struct {
	Person *Person
}

// NewStruct creates a new relationship struct
func (*childR) NewStruct() *childR {
	return &childR{}
}

// childL is where Load methods for each relationship are stored.
type childL struct{}

var (
	childAllColumns            = []string{"person_id", "sub_no", "first_name", "last_name", "birthday"}
	childColumnsWithoutDefault = []string{"person_id", "sub_no", "first_name", "last_name", "birthday"}
	childColumnsWithDefault    = []string{}
	childPrimaryKeyColumns     = []string{"person_id", "sub_no"}
)

type (
	// ChildSlice is an alias for a slice of pointers to Child.
	// This should generally be used opposed to []Child.
	ChildSlice []*Child
	// ChildHook is the signature for custom Child hook methods
	ChildHook func(context.Context, boil.ContextExecutor, *Child) error

	childQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	childType                 = reflect.TypeOf(&Child{})
	childMapping              = queries.MakeStructMapping(childType)
	childPrimaryKeyMapping, _ = queries.BindMapping(childType, childMapping, childPrimaryKeyColumns)
	childInsertCacheMut       sync.RWMutex
	childInsertCache          = make(map[string]insertCache)
	childUpdateCacheMut       sync.RWMutex
	childUpdateCache          = make(map[string]updateCache)
	childUpsertCacheMut       sync.RWMutex
	childUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var childBeforeInsertHooks []ChildHook
var childBeforeUpdateHooks []ChildHook
var childBeforeDeleteHooks []ChildHook
var childBeforeUpsertHooks []ChildHook

var childAfterInsertHooks []ChildHook
var childAfterSelectHooks []ChildHook
var childAfterUpdateHooks []ChildHook
var childAfterDeleteHooks []ChildHook
var childAfterUpsertHooks []ChildHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Child) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Child) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Child) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Child) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Child) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Child) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Child) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Child) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Child) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range childAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChildHook registers your hook function for all future operations.
func AddChildHook(hookPoint boil.HookPoint, childHook ChildHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		childBeforeInsertHooks = append(childBeforeInsertHooks, childHook)
	case boil.BeforeUpdateHook:
		childBeforeUpdateHooks = append(childBeforeUpdateHooks, childHook)
	case boil.BeforeDeleteHook:
		childBeforeDeleteHooks = append(childBeforeDeleteHooks, childHook)
	case boil.BeforeUpsertHook:
		childBeforeUpsertHooks = append(childBeforeUpsertHooks, childHook)
	case boil.AfterInsertHook:
		childAfterInsertHooks = append(childAfterInsertHooks, childHook)
	case boil.AfterSelectHook:
		childAfterSelectHooks = append(childAfterSelectHooks, childHook)
	case boil.AfterUpdateHook:
		childAfterUpdateHooks = append(childAfterUpdateHooks, childHook)
	case boil.AfterDeleteHook:
		childAfterDeleteHooks = append(childAfterDeleteHooks, childHook)
	case boil.AfterUpsertHook:
		childAfterUpsertHooks = append(childAfterUpsertHooks, childHook)
	}
}

// OneG returns a single child record from the query using the global executor.
func (q childQuery) OneG(ctx context.Context) (*Child, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single child record from the query.
func (q childQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Child, error) {
	o := &Child{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for children")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Child records from the query using the global executor.
func (q childQuery) AllG(ctx context.Context) (ChildSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Child records from the query.
func (q childQuery) All(ctx context.Context, exec boil.ContextExecutor) (ChildSlice, error) {
	var o []*Child

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Child slice")
	}

	if len(childAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Child records in the query, and panics on error.
func (q childQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Child records in the query.
func (q childQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count children rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q childQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q childQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if children exists")
	}

	return count > 0, nil
}

// Person pointed to by the foreign key.
func (o *Child) Person(mods ...qm.QueryMod) personQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`person_id` = ?", o.PersonID),
	}

	queryMods = append(queryMods, mods...)

	query := Persons(queryMods...)
	queries.SetFrom(query.Query, "`persons`")

	return query
}

// LoadPerson allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (childL) LoadPerson(ctx context.Context, e boil.ContextExecutor, singular bool, maybeChild interface{}, mods queries.Applicator) error {
	var slice []*Child
	var object *Child

	if singular {
		object = maybeChild.(*Child)
	} else {
		slice = *maybeChild.(*[]*Child)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &childR{}
		}
		args = append(args, object.PersonID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &childR{}
			}

			for _, a := range args {
				if a == obj.PersonID {
					continue Outer
				}
			}

			args = append(args, obj.PersonID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`persons`), qm.WhereIn(`persons.person_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Person")
	}

	var resultSlice []*Person
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Person")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for persons")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for persons")
	}

	if len(childAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Person = foreign
		if foreign.R == nil {
			foreign.R = &personR{}
		}
		foreign.R.Children = append(foreign.R.Children, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PersonID == foreign.PersonID {
				local.R.Person = foreign
				if foreign.R == nil {
					foreign.R = &personR{}
				}
				foreign.R.Children = append(foreign.R.Children, local)
				break
			}
		}
	}

	return nil
}

// SetPersonG of the child to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Children.
// Uses the global database handle.
func (o *Child) SetPersonG(ctx context.Context, insert bool, related *Person) error {
	return o.SetPerson(ctx, boil.GetContextDB(), insert, related)
}

// SetPerson of the child to the related item.
// Sets o.R.Person to related.
// Adds o to related.R.Children.
func (o *Child) SetPerson(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Person) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `children` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"person_id"}),
		strmangle.WhereClause("`", "`", 0, childPrimaryKeyColumns),
	)
	values := []interface{}{related.PersonID, o.PersonID, o.SubNo}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PersonID = related.PersonID
	if o.R == nil {
		o.R = &childR{
			Person: related,
		}
	} else {
		o.R.Person = related
	}

	if related.R == nil {
		related.R = &personR{
			Children: ChildSlice{o},
		}
	} else {
		related.R.Children = append(related.R.Children, o)
	}

	return nil
}

// Children retrieves all the records using an executor.
func Children(mods ...qm.QueryMod) childQuery {
	mods = append(mods, qm.From("`children`"))
	return childQuery{NewQuery(mods...)}
}

// FindChildG retrieves a single record by ID.
func FindChildG(ctx context.Context, personID int64, subNo int8, selectCols ...string) (*Child, error) {
	return FindChild(ctx, boil.GetContextDB(), personID, subNo, selectCols...)
}

// FindChild retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChild(ctx context.Context, exec boil.ContextExecutor, personID int64, subNo int8, selectCols ...string) (*Child, error) {
	childObj := &Child{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `children` where `person_id`=? AND `sub_no`=?", sel,
	)

	q := queries.Raw(query, personID, subNo)

	err := q.Bind(ctx, exec, childObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from children")
	}

	return childObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Child) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Child) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no children provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(childColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	childInsertCacheMut.RLock()
	cache, cached := childInsertCache[key]
	childInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			childAllColumns,
			childColumnsWithDefault,
			childColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(childType, childMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(childType, childMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `children` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `children` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `children` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, childPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into children")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PersonID,
		o.SubNo,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for children")
	}

CacheNoHooks:
	if !cached {
		childInsertCacheMut.Lock()
		childInsertCache[key] = cache
		childInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Child record using the global executor.
// See Update for more documentation.
func (o *Child) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Child.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Child) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	childUpdateCacheMut.RLock()
	cache, cached := childUpdateCache[key]
	childUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			childAllColumns,
			childPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update children, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `children` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, childPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(childType, childMapping, append(wl, childPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update children row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for children")
	}

	if !cached {
		childUpdateCacheMut.Lock()
		childUpdateCache[key] = cache
		childUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q childQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q childQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for children")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for children")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ChildSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChildSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `children` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, childPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in child slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all child")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Child) UpsertG(ctx context.Context, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateColumns, insertColumns)
}

var mySQLChildUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Child) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no children provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(childColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLChildUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	childUpsertCacheMut.RLock()
	cache, cached := childUpsertCache[key]
	childUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			childAllColumns,
			childColumnsWithDefault,
			childColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			childAllColumns,
			childPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert children, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "children", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `children` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(childType, childMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(childType, childMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for children")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(childType, childMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for children")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for children")
	}

CacheNoHooks:
	if !cached {
		childUpsertCacheMut.Lock()
		childUpsertCache[key] = cache
		childUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Child record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Child) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Child record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Child) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Child provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), childPrimaryKeyMapping)
	sql := "DELETE FROM `children` WHERE `person_id`=? AND `sub_no`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from children")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for children")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q childQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no childQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from children")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for children")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ChildSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChildSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(childBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `children` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, childPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from child slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for children")
	}

	if len(childAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Child) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Child provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Child) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindChild(ctx, exec, o.PersonID, o.SubNo)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChildSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ChildSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChildSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ChildSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), childPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `children`.* FROM `children` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, childPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChildSlice")
	}

	*o = slice

	return nil
}

// ChildExistsG checks if the Child row exists.
func ChildExistsG(ctx context.Context, personID int64, subNo int8) (bool, error) {
	return ChildExists(ctx, boil.GetContextDB(), personID, subNo)
}

// ChildExists checks if the Child row exists.
func ChildExists(ctx context.Context, exec boil.ContextExecutor, personID int64, subNo int8) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `children` where `person_id`=? AND `sub_no`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, personID, subNo)
	}
	row := exec.QueryRowContext(ctx, sql, personID, subNo)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if children exists")
	}

	return exists, nil
}
