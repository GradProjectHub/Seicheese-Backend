// Code generated by SQLBoiler 4.17.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PointLog is an object representing the database table.
type PointLog struct {
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UserID    uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Point     int       `boil:"point" json:"point" toml:"point" yaml:"point"`

	R *pointLogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pointLogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PointLogColumns = struct {
	CreatedAt string
	UserID    string
	Point     string
}{
	CreatedAt: "created_at",
	UserID:    "user_id",
	Point:     "point",
}

var PointLogTableColumns = struct {
	CreatedAt string
	UserID    string
	Point     string
}{
	CreatedAt: "point_logs.created_at",
	UserID:    "point_logs.user_id",
	Point:     "point_logs.point",
}

// Generated where

var PointLogWhere = struct {
	CreatedAt whereHelpertime_Time
	UserID    whereHelperuint
	Point     whereHelperint
}{
	CreatedAt: whereHelpertime_Time{field: "`point_logs`.`created_at`"},
	UserID:    whereHelperuint{field: "`point_logs`.`user_id`"},
	Point:     whereHelperint{field: "`point_logs`.`point`"},
}

// PointLogRels is where relationship names are stored.
var PointLogRels = struct {
	User string
}{
	User: "User",
}

// pointLogR is where relationships are stored.
type pointLogR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*pointLogR) NewStruct() *pointLogR {
	return &pointLogR{}
}

func (r *pointLogR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// pointLogL is where Load methods for each relationship are stored.
type pointLogL struct{}

var (
	pointLogAllColumns            = []string{"created_at", "user_id", "point"}
	pointLogColumnsWithoutDefault = []string{"user_id", "point"}
	pointLogColumnsWithDefault    = []string{"created_at"}
	pointLogPrimaryKeyColumns     = []string{"created_at", "user_id"}
	pointLogGeneratedColumns      = []string{}
)

type (
	// PointLogSlice is an alias for a slice of pointers to PointLog.
	// This should almost always be used instead of []PointLog.
	PointLogSlice []*PointLog
	// PointLogHook is the signature for custom PointLog hook methods
	PointLogHook func(context.Context, boil.ContextExecutor, *PointLog) error

	pointLogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pointLogType                 = reflect.TypeOf(&PointLog{})
	pointLogMapping              = queries.MakeStructMapping(pointLogType)
	pointLogPrimaryKeyMapping, _ = queries.BindMapping(pointLogType, pointLogMapping, pointLogPrimaryKeyColumns)
	pointLogInsertCacheMut       sync.RWMutex
	pointLogInsertCache          = make(map[string]insertCache)
	pointLogUpdateCacheMut       sync.RWMutex
	pointLogUpdateCache          = make(map[string]updateCache)
	pointLogUpsertCacheMut       sync.RWMutex
	pointLogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var pointLogAfterSelectMu sync.Mutex
var pointLogAfterSelectHooks []PointLogHook

var pointLogBeforeInsertMu sync.Mutex
var pointLogBeforeInsertHooks []PointLogHook
var pointLogAfterInsertMu sync.Mutex
var pointLogAfterInsertHooks []PointLogHook

var pointLogBeforeUpdateMu sync.Mutex
var pointLogBeforeUpdateHooks []PointLogHook
var pointLogAfterUpdateMu sync.Mutex
var pointLogAfterUpdateHooks []PointLogHook

var pointLogBeforeDeleteMu sync.Mutex
var pointLogBeforeDeleteHooks []PointLogHook
var pointLogAfterDeleteMu sync.Mutex
var pointLogAfterDeleteHooks []PointLogHook

var pointLogBeforeUpsertMu sync.Mutex
var pointLogBeforeUpsertHooks []PointLogHook
var pointLogAfterUpsertMu sync.Mutex
var pointLogAfterUpsertHooks []PointLogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PointLog) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PointLog) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PointLog) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PointLog) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PointLog) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PointLog) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PointLog) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PointLog) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PointLog) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range pointLogAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPointLogHook registers your hook function for all future operations.
func AddPointLogHook(hookPoint boil.HookPoint, pointLogHook PointLogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		pointLogAfterSelectMu.Lock()
		pointLogAfterSelectHooks = append(pointLogAfterSelectHooks, pointLogHook)
		pointLogAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		pointLogBeforeInsertMu.Lock()
		pointLogBeforeInsertHooks = append(pointLogBeforeInsertHooks, pointLogHook)
		pointLogBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		pointLogAfterInsertMu.Lock()
		pointLogAfterInsertHooks = append(pointLogAfterInsertHooks, pointLogHook)
		pointLogAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		pointLogBeforeUpdateMu.Lock()
		pointLogBeforeUpdateHooks = append(pointLogBeforeUpdateHooks, pointLogHook)
		pointLogBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		pointLogAfterUpdateMu.Lock()
		pointLogAfterUpdateHooks = append(pointLogAfterUpdateHooks, pointLogHook)
		pointLogAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		pointLogBeforeDeleteMu.Lock()
		pointLogBeforeDeleteHooks = append(pointLogBeforeDeleteHooks, pointLogHook)
		pointLogBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		pointLogAfterDeleteMu.Lock()
		pointLogAfterDeleteHooks = append(pointLogAfterDeleteHooks, pointLogHook)
		pointLogAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		pointLogBeforeUpsertMu.Lock()
		pointLogBeforeUpsertHooks = append(pointLogBeforeUpsertHooks, pointLogHook)
		pointLogBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		pointLogAfterUpsertMu.Lock()
		pointLogAfterUpsertHooks = append(pointLogAfterUpsertHooks, pointLogHook)
		pointLogAfterUpsertMu.Unlock()
	}
}

// One returns a single pointLog record from the query.
func (q pointLogQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PointLog, error) {
	o := &PointLog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for point_logs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PointLog records from the query.
func (q pointLogQuery) All(ctx context.Context, exec boil.ContextExecutor) (PointLogSlice, error) {
	var o []*PointLog

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PointLog slice")
	}

	if len(pointLogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PointLog records in the query.
func (q pointLogQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count point_logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pointLogQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if point_logs exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *PointLog) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pointLogL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePointLog interface{}, mods queries.Applicator) error {
	var slice []*PointLog
	var object *PointLog

	if singular {
		var ok bool
		object, ok = maybePointLog.(*PointLog)
		if !ok {
			object = new(PointLog)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePointLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePointLog))
			}
		}
	} else {
		s, ok := maybePointLog.(*[]*PointLog)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePointLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePointLog))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &pointLogR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pointLogR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.user_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.PointLogs = append(foreign.R.PointLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.PointLogs = append(foreign.R.PointLogs, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the pointLog to the related item.
// Sets o.R.User to related.
// Adds o to related.R.PointLogs.
func (o *PointLog) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `point_logs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, pointLogPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.CreatedAt, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &pointLogR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			PointLogs: PointLogSlice{o},
		}
	} else {
		related.R.PointLogs = append(related.R.PointLogs, o)
	}

	return nil
}

// PointLogs retrieves all the records using an executor.
func PointLogs(mods ...qm.QueryMod) pointLogQuery {
	mods = append(mods, qm.From("`point_logs`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`point_logs`.*"})
	}

	return pointLogQuery{q}
}

// FindPointLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPointLog(ctx context.Context, exec boil.ContextExecutor, createdAt time.Time, userID uint, selectCols ...string) (*PointLog, error) {
	pointLogObj := &PointLog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `point_logs` where `created_at`=? AND `user_id`=?", sel,
	)

	q := queries.Raw(query, createdAt, userID)

	err := q.Bind(ctx, exec, pointLogObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from point_logs")
	}

	if err = pointLogObj.doAfterSelectHooks(ctx, exec); err != nil {
		return pointLogObj, err
	}

	return pointLogObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PointLog) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no point_logs provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pointLogColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pointLogInsertCacheMut.RLock()
	cache, cached := pointLogInsertCache[key]
	pointLogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pointLogAllColumns,
			pointLogColumnsWithDefault,
			pointLogColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pointLogType, pointLogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pointLogType, pointLogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `point_logs` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `point_logs` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `point_logs` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, pointLogPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into point_logs")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CreatedAt,
		o.UserID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for point_logs")
	}

CacheNoHooks:
	if !cached {
		pointLogInsertCacheMut.Lock()
		pointLogInsertCache[key] = cache
		pointLogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PointLog.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PointLog) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pointLogUpdateCacheMut.RLock()
	cache, cached := pointLogUpdateCache[key]
	pointLogUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pointLogAllColumns,
			pointLogPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update point_logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `point_logs` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, pointLogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pointLogType, pointLogMapping, append(wl, pointLogPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update point_logs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for point_logs")
	}

	if !cached {
		pointLogUpdateCacheMut.Lock()
		pointLogUpdateCache[key] = cache
		pointLogUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pointLogQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for point_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for point_logs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PointLogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pointLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `point_logs` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pointLogPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pointLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pointLog")
	}
	return rowsAff, nil
}

var mySQLPointLogUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PointLog) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no point_logs provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pointLogColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPointLogUniqueColumns, o)

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

	pointLogUpsertCacheMut.RLock()
	cache, cached := pointLogUpsertCache[key]
	pointLogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			pointLogAllColumns,
			pointLogColumnsWithDefault,
			pointLogColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			pointLogAllColumns,
			pointLogPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert point_logs, could not build update column list")
		}

		ret := strmangle.SetComplement(pointLogAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`point_logs`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `point_logs` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(pointLogType, pointLogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pointLogType, pointLogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for point_logs")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(pointLogType, pointLogMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for point_logs")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for point_logs")
	}

CacheNoHooks:
	if !cached {
		pointLogUpsertCacheMut.Lock()
		pointLogUpsertCache[key] = cache
		pointLogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PointLog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PointLog) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PointLog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pointLogPrimaryKeyMapping)
	sql := "DELETE FROM `point_logs` WHERE `created_at`=? AND `user_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from point_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for point_logs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pointLogQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pointLogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from point_logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for point_logs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PointLogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(pointLogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pointLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `point_logs` WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 0, pointLogPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pointLog slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for point_logs")
	}

	if len(pointLogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PointLog) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPointLog(ctx, exec, o.CreatedAt, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PointLogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PointLogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pointLogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `point_logs`.* FROM `point_logs` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pointLogPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PointLogSlice")
	}

	*o = slice

	return nil
}

// PointLogExists checks if the PointLog row exists.
func PointLogExists(ctx context.Context, exec boil.ContextExecutor, createdAt time.Time, userID uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `point_logs` where `created_at`=? AND `user_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, createdAt, userID)
	}
	row := exec.QueryRowContext(ctx, sql, createdAt, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if point_logs exists")
	}

	return exists, nil
}

// Exists checks if the PointLog row exists.
func (o *PointLog) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PointLogExists(ctx, exec, o.CreatedAt, o.UserID)
}
