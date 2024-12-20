// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"shaanmei_backend/dao/model"
)

func newTimeseriesDatum(db *gorm.DB, opts ...gen.DOOption) timeseriesDatum {
	_timeseriesDatum := timeseriesDatum{}

	_timeseriesDatum.timeseriesDatumDo.UseDB(db, opts...)
	_timeseriesDatum.timeseriesDatumDo.UseModel(&model.TimeseriesDatum{})

	tableName := _timeseriesDatum.timeseriesDatumDo.TableName()
	_timeseriesDatum.ALL = field.NewAsterisk(tableName)
	_timeseriesDatum.ID = field.NewInt32(tableName, "id")
	_timeseriesDatum.AvgShipmentPrice = field.NewFloat32(tableName, "avg_shipment_price")
	_timeseriesDatum.CoalType = field.NewString(tableName, "coal_type")
	_timeseriesDatum.ShipmentVolume = field.NewFloat32(tableName, "shipment_volume")
	_timeseriesDatum.Timestamp = field.NewTime(tableName, "timestamp")
	_timeseriesDatum.CarrierName = field.NewString(tableName, "carrier_name")
	_timeseriesDatum.MineName = field.NewString(tableName, "mine_name")

	_timeseriesDatum.fillFieldMap()

	return _timeseriesDatum
}

type timeseriesDatum struct {
	timeseriesDatumDo timeseriesDatumDo

	ALL              field.Asterisk
	ID               field.Int32
	AvgShipmentPrice field.Float32
	CoalType         field.String
	ShipmentVolume   field.Float32
	Timestamp        field.Time
	CarrierName      field.String
	MineName         field.String

	fieldMap map[string]field.Expr
}

func (t timeseriesDatum) Table(newTableName string) *timeseriesDatum {
	t.timeseriesDatumDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t timeseriesDatum) As(alias string) *timeseriesDatum {
	t.timeseriesDatumDo.DO = *(t.timeseriesDatumDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *timeseriesDatum) updateTableName(table string) *timeseriesDatum {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.AvgShipmentPrice = field.NewFloat32(table, "avg_shipment_price")
	t.CoalType = field.NewString(table, "coal_type")
	t.ShipmentVolume = field.NewFloat32(table, "shipment_volume")
	t.Timestamp = field.NewTime(table, "timestamp")
	t.CarrierName = field.NewString(table, "carrier_name")
	t.MineName = field.NewString(table, "mine_name")

	t.fillFieldMap()

	return t
}

func (t *timeseriesDatum) WithContext(ctx context.Context) *timeseriesDatumDo {
	return t.timeseriesDatumDo.WithContext(ctx)
}

func (t timeseriesDatum) TableName() string { return t.timeseriesDatumDo.TableName() }

func (t timeseriesDatum) Alias() string { return t.timeseriesDatumDo.Alias() }

func (t timeseriesDatum) Columns(cols ...field.Expr) gen.Columns {
	return t.timeseriesDatumDo.Columns(cols...)
}

func (t *timeseriesDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *timeseriesDatum) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 7)
	t.fieldMap["id"] = t.ID
	t.fieldMap["avg_shipment_price"] = t.AvgShipmentPrice
	t.fieldMap["coal_type"] = t.CoalType
	t.fieldMap["shipment_volume"] = t.ShipmentVolume
	t.fieldMap["timestamp"] = t.Timestamp
	t.fieldMap["carrier_name"] = t.CarrierName
	t.fieldMap["mine_name"] = t.MineName
}

func (t timeseriesDatum) clone(db *gorm.DB) timeseriesDatum {
	t.timeseriesDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t timeseriesDatum) replaceDB(db *gorm.DB) timeseriesDatum {
	t.timeseriesDatumDo.ReplaceDB(db)
	return t
}

type timeseriesDatumDo struct{ gen.DO }

func (t timeseriesDatumDo) Debug() *timeseriesDatumDo {
	return t.withDO(t.DO.Debug())
}

func (t timeseriesDatumDo) WithContext(ctx context.Context) *timeseriesDatumDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t timeseriesDatumDo) ReadDB() *timeseriesDatumDo {
	return t.Clauses(dbresolver.Read)
}

func (t timeseriesDatumDo) WriteDB() *timeseriesDatumDo {
	return t.Clauses(dbresolver.Write)
}

func (t timeseriesDatumDo) Session(config *gorm.Session) *timeseriesDatumDo {
	return t.withDO(t.DO.Session(config))
}

func (t timeseriesDatumDo) Clauses(conds ...clause.Expression) *timeseriesDatumDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t timeseriesDatumDo) Returning(value interface{}, columns ...string) *timeseriesDatumDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t timeseriesDatumDo) Not(conds ...gen.Condition) *timeseriesDatumDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t timeseriesDatumDo) Or(conds ...gen.Condition) *timeseriesDatumDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t timeseriesDatumDo) Select(conds ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t timeseriesDatumDo) Where(conds ...gen.Condition) *timeseriesDatumDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t timeseriesDatumDo) Order(conds ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t timeseriesDatumDo) Distinct(cols ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t timeseriesDatumDo) Omit(cols ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t timeseriesDatumDo) Join(table schema.Tabler, on ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t timeseriesDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t timeseriesDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t timeseriesDatumDo) Group(cols ...field.Expr) *timeseriesDatumDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t timeseriesDatumDo) Having(conds ...gen.Condition) *timeseriesDatumDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t timeseriesDatumDo) Limit(limit int) *timeseriesDatumDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t timeseriesDatumDo) Offset(offset int) *timeseriesDatumDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t timeseriesDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *timeseriesDatumDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t timeseriesDatumDo) Unscoped() *timeseriesDatumDo {
	return t.withDO(t.DO.Unscoped())
}

func (t timeseriesDatumDo) Create(values ...*model.TimeseriesDatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t timeseriesDatumDo) CreateInBatches(values []*model.TimeseriesDatum, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t timeseriesDatumDo) Save(values ...*model.TimeseriesDatum) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t timeseriesDatumDo) First() (*model.TimeseriesDatum, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeseriesDatum), nil
	}
}

func (t timeseriesDatumDo) Take() (*model.TimeseriesDatum, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeseriesDatum), nil
	}
}

func (t timeseriesDatumDo) Last() (*model.TimeseriesDatum, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeseriesDatum), nil
	}
}

func (t timeseriesDatumDo) Find() ([]*model.TimeseriesDatum, error) {
	result, err := t.DO.Find()
	return result.([]*model.TimeseriesDatum), err
}

func (t timeseriesDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TimeseriesDatum, err error) {
	buf := make([]*model.TimeseriesDatum, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t timeseriesDatumDo) FindInBatches(result *[]*model.TimeseriesDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t timeseriesDatumDo) Attrs(attrs ...field.AssignExpr) *timeseriesDatumDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t timeseriesDatumDo) Assign(attrs ...field.AssignExpr) *timeseriesDatumDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t timeseriesDatumDo) Joins(fields ...field.RelationField) *timeseriesDatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t timeseriesDatumDo) Preload(fields ...field.RelationField) *timeseriesDatumDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t timeseriesDatumDo) FirstOrInit() (*model.TimeseriesDatum, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeseriesDatum), nil
	}
}

func (t timeseriesDatumDo) FirstOrCreate() (*model.TimeseriesDatum, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TimeseriesDatum), nil
	}
}

func (t timeseriesDatumDo) FindByPage(offset int, limit int) (result []*model.TimeseriesDatum, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t timeseriesDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t timeseriesDatumDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t timeseriesDatumDo) Delete(models ...*model.TimeseriesDatum) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *timeseriesDatumDo) withDO(do gen.Dao) *timeseriesDatumDo {
	t.DO = *do.(*gen.DO)
	return t
}
