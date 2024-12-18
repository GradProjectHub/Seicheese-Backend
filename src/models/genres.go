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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Genre is an object representing the database table.
type Genre struct {
	GenreID   int       `boil:"genre_id" json:"genre_id" toml:"genre_id" yaml:"genre_id"`
	GenreName string    `boil:"genre_name" json:"genre_name" toml:"genre_name" yaml:"genre_name"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *genreR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L genreL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GenreColumns = struct {
	GenreID   string
	GenreName string
	CreatedAt string
	UpdatedAt string
}{
	GenreID:   "genre_id",
	GenreName: "genre_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var GenreTableColumns = struct {
	GenreID   string
	GenreName string
	CreatedAt string
	UpdatedAt string
}{
	GenreID:   "genres.genre_id",
	GenreName: "genres.genre_name",
	CreatedAt: "genres.created_at",
	UpdatedAt: "genres.updated_at",
}

// Generated where

var GenreWhere = struct {
	GenreID   whereHelperint
	GenreName whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	GenreID:   whereHelperint{field: "`genres`.`genre_id`"},
	GenreName: whereHelperstring{field: "`genres`.`genre_name`"},
	CreatedAt: whereHelpernull_Time{field: "`genres`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`genres`.`updated_at`"},
}

// GenreRels is where relationship names are stored.
var GenreRels = struct {
	Contents string
}{
	Contents: "Contents",
}

// genreR is where relationships are stored.
type genreR struct {
	Contents ContentSlice `boil:"Contents" json:"Contents" toml:"Contents" yaml:"Contents"`
}

// NewStruct creates a new relationship struct
func (*genreR) NewStruct() *genreR {
	return &genreR{}
}

func (r *genreR) GetContents() ContentSlice {
	if r == nil {
		return nil
	}
	return r.Contents
}

// genreL is where Load methods for each relationship are stored.
type genreL struct{}

var (
	genreAllColumns            = []string{"genre_id", "genre_name", "created_at", "updated_at"}
	genreColumnsWithoutDefault = []string{"genre_name"}
	genreColumnsWithDefault    = []string{"genre_id", "created_at", "updated_at"}
	genrePrimaryKeyColumns     = []string{"genre_id"}
	genreGeneratedColumns      = []string{}
)

type (
	// GenreSlice is an alias for a slice of pointers to Genre.
	// This should almost always be used instead of []Genre.
	GenreSlice []*Genre
	// GenreHook is the signature for custom Genre hook methods
	GenreHook func(context.Context, boil.ContextExecutor, *Genre) error

	genreQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	genreType                 = reflect.TypeOf(&Genre{})
	genreMapping              = queries.MakeStructMapping(genreType)
	genrePrimaryKeyMapping, _ = queries.BindMapping(genreType, genreMapping, genrePrimaryKeyColumns)
	genreInsertCacheMut       sync.RWMutex
	genreInsertCache          = make(map[string]insertCache)
	genreUpdateCacheMut       sync.RWMutex
	genreUpdateCache          = make(map[string]updateCache)
	genreUpsertCacheMut       sync.RWMutex
	genreUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var genreAfterSelectMu sync.Mutex
var genreAfterSelectHooks []GenreHook

var genreBeforeInsertMu sync.Mutex
var genreBeforeInsertHooks []GenreHook
var genreAfterInsertMu sync.Mutex
var genreAfterInsertHooks []GenreHook

var genreBeforeUpdateMu sync.Mutex
var genreBeforeUpdateHooks []GenreHook
var genreAfterUpdateMu sync.Mutex
var genreAfterUpdateHooks []GenreHook

var genreBeforeDeleteMu sync.Mutex
var genreBeforeDeleteHooks []GenreHook
var genreAfterDeleteMu sync.Mutex
var genreAfterDeleteHooks []GenreHook

var genreBeforeUpsertMu sync.Mutex
var genreBeforeUpsertHooks []GenreHook
var genreAfterUpsertMu sync.Mutex
var genreAfterUpsertHooks []GenreHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Genre) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Genre) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Genre) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Genre) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Genre) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Genre) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Genre) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Genre) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Genre) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range genreAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGenreHook registers your hook function for all future operations.
func AddGenreHook(hookPoint boil.HookPoint, genreHook GenreHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		genreAfterSelectMu.Lock()
		genreAfterSelectHooks = append(genreAfterSelectHooks, genreHook)
		genreAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		genreBeforeInsertMu.Lock()
		genreBeforeInsertHooks = append(genreBeforeInsertHooks, genreHook)
		genreBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		genreAfterInsertMu.Lock()
		genreAfterInsertHooks = append(genreAfterInsertHooks, genreHook)
		genreAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		genreBeforeUpdateMu.Lock()
		genreBeforeUpdateHooks = append(genreBeforeUpdateHooks, genreHook)
		genreBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		genreAfterUpdateMu.Lock()
		genreAfterUpdateHooks = append(genreAfterUpdateHooks, genreHook)
		genreAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		genreBeforeDeleteMu.Lock()
		genreBeforeDeleteHooks = append(genreBeforeDeleteHooks, genreHook)
		genreBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		genreAfterDeleteMu.Lock()
		genreAfterDeleteHooks = append(genreAfterDeleteHooks, genreHook)
		genreAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		genreBeforeUpsertMu.Lock()
		genreBeforeUpsertHooks = append(genreBeforeUpsertHooks, genreHook)
		genreBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		genreAfterUpsertMu.Lock()
		genreAfterUpsertHooks = append(genreAfterUpsertHooks, genreHook)
		genreAfterUpsertMu.Unlock()
	}
}

// One returns a single genre record from the query.
func (q genreQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Genre, error) {
	o := &Genre{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for genres")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Genre records from the query.
func (q genreQuery) All(ctx context.Context, exec boil.ContextExecutor) (GenreSlice, error) {
	var o []*Genre

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Genre slice")
	}

	if len(genreAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Genre records in the query.
func (q genreQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count genres rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q genreQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if genres exists")
	}

	return count > 0, nil
}

// Contents retrieves all the content's Contents with an executor.
func (o *Genre) Contents(mods ...qm.QueryMod) contentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`contents`.`genre_id`=?", o.GenreID),
	)

	return Contents(queryMods...)
}

// LoadContents allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (genreL) LoadContents(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGenre interface{}, mods queries.Applicator) error {
	var slice []*Genre
	var object *Genre

	if singular {
		var ok bool
		object, ok = maybeGenre.(*Genre)
		if !ok {
			object = new(Genre)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGenre)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGenre))
			}
		}
	} else {
		s, ok := maybeGenre.(*[]*Genre)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGenre)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGenre))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &genreR{}
		}
		args[object.GenreID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &genreR{}
			}
			args[obj.GenreID] = struct{}{}
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
		qm.From(`contents`),
		qm.WhereIn(`contents.genre_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load contents")
	}

	var resultSlice []*Content
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice contents")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on contents")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for contents")
	}

	if len(contentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Contents = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &contentR{}
			}
			foreign.R.Genre = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.GenreID == foreign.GenreID {
				local.R.Contents = append(local.R.Contents, foreign)
				if foreign.R == nil {
					foreign.R = &contentR{}
				}
				foreign.R.Genre = local
				break
			}
		}
	}

	return nil
}

// AddContents adds the given related objects to the existing relationships
// of the genre, optionally inserting them as new records.
// Appends related to o.R.Contents.
// Sets related.R.Genre appropriately.
func (o *Genre) AddContents(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Content) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.GenreID = o.GenreID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `contents` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"genre_id"}),
				strmangle.WhereClause("`", "`", 0, contentPrimaryKeyColumns),
			)
			values := []interface{}{o.GenreID, rel.ContentID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.GenreID = o.GenreID
		}
	}

	if o.R == nil {
		o.R = &genreR{
			Contents: related,
		}
	} else {
		o.R.Contents = append(o.R.Contents, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &contentR{
				Genre: o,
			}
		} else {
			rel.R.Genre = o
		}
	}
	return nil
}

// Genres retrieves all the records using an executor.
func Genres(mods ...qm.QueryMod) genreQuery {
	mods = append(mods, qm.From("`genres`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`genres`.*"})
	}

	return genreQuery{q}
}

// FindGenre retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGenre(ctx context.Context, exec boil.ContextExecutor, genreID int, selectCols ...string) (*Genre, error) {
	genreObj := &Genre{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `genres` where `genre_id`=?", sel,
	)

	q := queries.Raw(query, genreID)

	err := q.Bind(ctx, exec, genreObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from genres")
	}

	if err = genreObj.doAfterSelectHooks(ctx, exec); err != nil {
		return genreObj, err
	}

	return genreObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Genre) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no genres provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genreColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	genreInsertCacheMut.RLock()
	cache, cached := genreInsertCache[key]
	genreInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			genreAllColumns,
			genreColumnsWithDefault,
			genreColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(genreType, genreMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(genreType, genreMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `genres` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `genres` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `genres` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, genrePrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into genres")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.GenreID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == genreMapping["genre_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.GenreID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for genres")
	}

CacheNoHooks:
	if !cached {
		genreInsertCacheMut.Lock()
		genreInsertCache[key] = cache
		genreInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Genre.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Genre) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	genreUpdateCacheMut.RLock()
	cache, cached := genreUpdateCache[key]
	genreUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			genreAllColumns,
			genrePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update genres, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `genres` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, genrePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(genreType, genreMapping, append(wl, genrePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update genres row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for genres")
	}

	if !cached {
		genreUpdateCacheMut.Lock()
		genreUpdateCache[key] = cache
		genreUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q genreQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for genres")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GenreSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `genres` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, genrePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in genre slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all genre")
	}
	return rowsAff, nil
}

var mySQLGenreUniqueColumns = []string{
	"genre_id",
	"genre_name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Genre) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no genres provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(genreColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGenreUniqueColumns, o)

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

	genreUpsertCacheMut.RLock()
	cache, cached := genreUpsertCache[key]
	genreUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			genreAllColumns,
			genreColumnsWithDefault,
			genreColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			genreAllColumns,
			genrePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert genres, could not build update column list")
		}

		ret := strmangle.SetComplement(genreAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`genres`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `genres` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(genreType, genreMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(genreType, genreMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for genres")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.GenreID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == genreMapping["genre_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(genreType, genreMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for genres")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for genres")
	}

CacheNoHooks:
	if !cached {
		genreUpsertCacheMut.Lock()
		genreUpsertCache[key] = cache
		genreUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Genre record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Genre) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Genre provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), genrePrimaryKeyMapping)
	sql := "DELETE FROM `genres` WHERE `genre_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for genres")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q genreQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no genreQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from genres")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for genres")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GenreSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(genreBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `genres` WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 0, genrePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from genre slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for genres")
	}

	if len(genreAfterDeleteHooks) != 0 {
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
func (o *Genre) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGenre(ctx, exec, o.GenreID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GenreSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GenreSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), genrePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `genres`.* FROM `genres` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, genrePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GenreSlice")
	}

	*o = slice

	return nil
}

// GenreExists checks if the Genre row exists.
func GenreExists(ctx context.Context, exec boil.ContextExecutor, genreID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `genres` where `genre_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, genreID)
	}
	row := exec.QueryRowContext(ctx, sql, genreID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if genres exists")
	}

	return exists, nil
}

// Exists checks if the Genre row exists.
func (o *Genre) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GenreExists(ctx, exec, o.GenreID)
}
