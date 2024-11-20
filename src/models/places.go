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

// Place is an object representing the database table.
type Place struct {
	PlaceID   int       `boil:"place_id" json:"place_id" toml:"place_id" yaml:"place_id"`
	Address   string    `boil:"address" json:"address" toml:"address" yaml:"address"`
	ZipCode   string    `boil:"zip_code" json:"zip_code" toml:"zip_code" yaml:"zip_code"`
	CreatedAt null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *placeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L placeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PlaceColumns = struct {
	PlaceID   string
	Address   string
	ZipCode   string
	CreatedAt string
	UpdatedAt string
}{
	PlaceID:   "place_id",
	Address:   "address",
	ZipCode:   "zip_code",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var PlaceTableColumns = struct {
	PlaceID   string
	Address   string
	ZipCode   string
	CreatedAt string
	UpdatedAt string
}{
	PlaceID:   "places.place_id",
	Address:   "places.address",
	ZipCode:   "places.zip_code",
	CreatedAt: "places.created_at",
	UpdatedAt: "places.updated_at",
}

// Generated where

var PlaceWhere = struct {
	PlaceID   whereHelperint
	Address   whereHelperstring
	ZipCode   whereHelperstring
	CreatedAt whereHelpernull_Time
	UpdatedAt whereHelpernull_Time
}{
	PlaceID:   whereHelperint{field: "`places`.`place_id`"},
	Address:   whereHelperstring{field: "`places`.`address`"},
	ZipCode:   whereHelperstring{field: "`places`.`zip_code`"},
	CreatedAt: whereHelpernull_Time{field: "`places`.`created_at`"},
	UpdatedAt: whereHelpernull_Time{field: "`places`.`updated_at`"},
}

// PlaceRels is where relationship names are stored.
var PlaceRels = struct {
	Seichies string
}{
	Seichies: "Seichies",
}

// placeR is where relationships are stored.
type placeR struct {
	Seichies SeichySlice `boil:"Seichies" json:"Seichies" toml:"Seichies" yaml:"Seichies"`
}

// NewStruct creates a new relationship struct
func (*placeR) NewStruct() *placeR {
	return &placeR{}
}

func (r *placeR) GetSeichies() SeichySlice {
	if r == nil {
		return nil
	}
	return r.Seichies
}

// placeL is where Load methods for each relationship are stored.
type placeL struct{}

var (
	placeAllColumns            = []string{"place_id", "address", "zip_code", "created_at", "updated_at"}
	placeColumnsWithoutDefault = []string{"address", "zip_code"}
	placeColumnsWithDefault    = []string{"place_id", "created_at", "updated_at"}
	placePrimaryKeyColumns     = []string{"place_id"}
	placeGeneratedColumns      = []string{}
)

type (
	// PlaceSlice is an alias for a slice of pointers to Place.
	// This should almost always be used instead of []Place.
	PlaceSlice []*Place
	// PlaceHook is the signature for custom Place hook methods
	PlaceHook func(context.Context, boil.ContextExecutor, *Place) error

	placeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	placeType                 = reflect.TypeOf(&Place{})
	placeMapping              = queries.MakeStructMapping(placeType)
	placePrimaryKeyMapping, _ = queries.BindMapping(placeType, placeMapping, placePrimaryKeyColumns)
	placeInsertCacheMut       sync.RWMutex
	placeInsertCache          = make(map[string]insertCache)
	placeUpdateCacheMut       sync.RWMutex
	placeUpdateCache          = make(map[string]updateCache)
	placeUpsertCacheMut       sync.RWMutex
	placeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var placeAfterSelectMu sync.Mutex
var placeAfterSelectHooks []PlaceHook

var placeBeforeInsertMu sync.Mutex
var placeBeforeInsertHooks []PlaceHook
var placeAfterInsertMu sync.Mutex
var placeAfterInsertHooks []PlaceHook

var placeBeforeUpdateMu sync.Mutex
var placeBeforeUpdateHooks []PlaceHook
var placeAfterUpdateMu sync.Mutex
var placeAfterUpdateHooks []PlaceHook

var placeBeforeDeleteMu sync.Mutex
var placeBeforeDeleteHooks []PlaceHook
var placeAfterDeleteMu sync.Mutex
var placeAfterDeleteHooks []PlaceHook

var placeBeforeUpsertMu sync.Mutex
var placeBeforeUpsertHooks []PlaceHook
var placeAfterUpsertMu sync.Mutex
var placeAfterUpsertHooks []PlaceHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Place) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Place) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Place) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Place) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Place) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Place) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Place) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Place) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Place) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range placeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPlaceHook registers your hook function for all future operations.
func AddPlaceHook(hookPoint boil.HookPoint, placeHook PlaceHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		placeAfterSelectMu.Lock()
		placeAfterSelectHooks = append(placeAfterSelectHooks, placeHook)
		placeAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		placeBeforeInsertMu.Lock()
		placeBeforeInsertHooks = append(placeBeforeInsertHooks, placeHook)
		placeBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		placeAfterInsertMu.Lock()
		placeAfterInsertHooks = append(placeAfterInsertHooks, placeHook)
		placeAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		placeBeforeUpdateMu.Lock()
		placeBeforeUpdateHooks = append(placeBeforeUpdateHooks, placeHook)
		placeBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		placeAfterUpdateMu.Lock()
		placeAfterUpdateHooks = append(placeAfterUpdateHooks, placeHook)
		placeAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		placeBeforeDeleteMu.Lock()
		placeBeforeDeleteHooks = append(placeBeforeDeleteHooks, placeHook)
		placeBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		placeAfterDeleteMu.Lock()
		placeAfterDeleteHooks = append(placeAfterDeleteHooks, placeHook)
		placeAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		placeBeforeUpsertMu.Lock()
		placeBeforeUpsertHooks = append(placeBeforeUpsertHooks, placeHook)
		placeBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		placeAfterUpsertMu.Lock()
		placeAfterUpsertHooks = append(placeAfterUpsertHooks, placeHook)
		placeAfterUpsertMu.Unlock()
	}
}

// One returns a single place record from the query.
func (q placeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Place, error) {
	o := &Place{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for places")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Place records from the query.
func (q placeQuery) All(ctx context.Context, exec boil.ContextExecutor) (PlaceSlice, error) {
	var o []*Place

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Place slice")
	}

	if len(placeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Place records in the query.
func (q placeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count places rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q placeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if places exists")
	}

	return count > 0, nil
}

// Seichies retrieves all the seichy's Seichies with an executor.
func (o *Place) Seichies(mods ...qm.QueryMod) seichyQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`seichies`.`place_id`=?", o.PlaceID),
	)

	return Seichies(queryMods...)
}

// LoadSeichies allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (placeL) LoadSeichies(ctx context.Context, e boil.ContextExecutor, singular bool, maybePlace interface{}, mods queries.Applicator) error {
	var slice []*Place
	var object *Place

	if singular {
		var ok bool
		object, ok = maybePlace.(*Place)
		if !ok {
			object = new(Place)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePlace)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePlace))
			}
		}
	} else {
		s, ok := maybePlace.(*[]*Place)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePlace)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePlace))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &placeR{}
		}
		args[object.PlaceID] = struct{}{}
	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &placeR{}
			}
			args[obj.PlaceID] = struct{}{}
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
		qm.From(`seichies`),
		qm.WhereIn(`seichies.place_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load seichies")
	}

	var resultSlice []*Seichy
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice seichies")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on seichies")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for seichies")
	}

	if len(seichyAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Seichies = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &seichyR{}
			}
			foreign.R.Place = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.PlaceID == foreign.PlaceID {
				local.R.Seichies = append(local.R.Seichies, foreign)
				if foreign.R == nil {
					foreign.R = &seichyR{}
				}
				foreign.R.Place = local
				break
			}
		}
	}

	return nil
}

// AddSeichies adds the given related objects to the existing relationships
// of the place, optionally inserting them as new records.
// Appends related to o.R.Seichies.
// Sets related.R.Place appropriately.
func (o *Place) AddSeichies(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Seichy) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PlaceID = o.PlaceID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `seichies` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"place_id"}),
				strmangle.WhereClause("`", "`", 0, seichyPrimaryKeyColumns),
			)
			values := []interface{}{o.PlaceID, rel.SeichiID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PlaceID = o.PlaceID
		}
	}

	if o.R == nil {
		o.R = &placeR{
			Seichies: related,
		}
	} else {
		o.R.Seichies = append(o.R.Seichies, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &seichyR{
				Place: o,
			}
		} else {
			rel.R.Place = o
		}
	}
	return nil
}

// Places retrieves all the records using an executor.
func Places(mods ...qm.QueryMod) placeQuery {
	mods = append(mods, qm.From("`places`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`places`.*"})
	}

	return placeQuery{q}
}

// FindPlace retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPlace(ctx context.Context, exec boil.ContextExecutor, placeID int, selectCols ...string) (*Place, error) {
	placeObj := &Place{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `places` where `place_id`=?", sel,
	)

	q := queries.Raw(query, placeID)

	err := q.Bind(ctx, exec, placeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from places")
	}

	if err = placeObj.doAfterSelectHooks(ctx, exec); err != nil {
		return placeObj, err
	}

	return placeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Place) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no places provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(placeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	placeInsertCacheMut.RLock()
	cache, cached := placeInsertCache[key]
	placeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			placeAllColumns,
			placeColumnsWithDefault,
			placeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(placeType, placeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(placeType, placeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `places` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `places` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `places` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, placePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into places")
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

	o.PlaceID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == placeMapping["place_id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PlaceID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for places")
	}

CacheNoHooks:
	if !cached {
		placeInsertCacheMut.Lock()
		placeInsertCache[key] = cache
		placeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Place.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Place) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	placeUpdateCacheMut.RLock()
	cache, cached := placeUpdateCache[key]
	placeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			placeAllColumns,
			placePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update places, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `places` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, placePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(placeType, placeMapping, append(wl, placePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update places row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for places")
	}

	if !cached {
		placeUpdateCacheMut.Lock()
		placeUpdateCache[key] = cache
		placeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q placeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for places")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for places")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PlaceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), placePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `places` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, placePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in place slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all place")
	}
	return rowsAff, nil
}

var mySQLPlaceUniqueColumns = []string{
	"place_id",
	"address",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Place) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no places provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(placeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPlaceUniqueColumns, o)

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

	placeUpsertCacheMut.RLock()
	cache, cached := placeUpsertCache[key]
	placeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			placeAllColumns,
			placeColumnsWithDefault,
			placeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			placeAllColumns,
			placePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert places, could not build update column list")
		}

		ret := strmangle.SetComplement(placeAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`places`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `places` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(placeType, placeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(placeType, placeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for places")
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

	o.PlaceID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == placeMapping["place_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(placeType, placeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for places")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for places")
	}

CacheNoHooks:
	if !cached {
		placeUpsertCacheMut.Lock()
		placeUpsertCache[key] = cache
		placeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Place record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Place) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Place provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), placePrimaryKeyMapping)
	sql := "DELETE FROM `places` WHERE `place_id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from places")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for places")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q placeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no placeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from places")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for places")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PlaceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(placeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), placePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `places` WHERE " +
		strmangle.WhereInClause(string(dialect.LQ), string(dialect.RQ), 0, placePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from place slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for places")
	}

	if len(placeAfterDeleteHooks) != 0 {
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
func (o *Place) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPlace(ctx, exec, o.PlaceID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PlaceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PlaceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), placePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `places`.* FROM `places` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, placePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PlaceSlice")
	}

	*o = slice

	return nil
}

// PlaceExists checks if the Place row exists.
func PlaceExists(ctx context.Context, exec boil.ContextExecutor, placeID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `places` where `place_id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, placeID)
	}
	row := exec.QueryRowContext(ctx, sql, placeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if places exists")
	}

	return exists, nil
}

// Exists checks if the Place row exists.
func (o *Place) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PlaceExists(ctx, exec, o.PlaceID)
}
