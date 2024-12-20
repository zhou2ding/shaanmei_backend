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

func newRouteDatum(db *gorm.DB, opts ...gen.DOOption) routeDatum {
	_routeDatum := routeDatum{}

	_routeDatum.routeDatumDo.UseDB(db, opts...)
	_routeDatum.routeDatumDo.UseModel(&model.RouteDatum{})

	tableName := _routeDatum.routeDatumDo.TableName()
	_routeDatum.ALL = field.NewAsterisk(tableName)
	_routeDatum.ID = field.NewInt32(tableName, "id")
	_routeDatum.RouteName = field.NewString(tableName, "route_name")
	_routeDatum.MineName = field.NewString(tableName, "mine_name")
	_routeDatum.DestinationRegion = field.NewString(tableName, "destination_region")
	_routeDatum.RouteDistance = field.NewFloat32(tableName, "route_distance")
	_routeDatum.AvgShipmentUnitPrice = field.NewFloat32(tableName, "avg_shipment_unit_price")
	_routeDatum.ShipmentVolume = field.NewFloat32(tableName, "shipment_volume")
	_routeDatum.Timestamp = field.NewTime(tableName, "timestamp")

	_routeDatum.fillFieldMap()

	return _routeDatum
}

type routeDatum struct {
	routeDatumDo routeDatumDo

	ALL                  field.Asterisk
	ID                   field.Int32
	RouteName            field.String
	MineName             field.String
	DestinationRegion    field.String
	RouteDistance        field.Float32
	AvgShipmentUnitPrice field.Float32
	ShipmentVolume       field.Float32
	Timestamp            field.Time

	fieldMap map[string]field.Expr
}

func (r routeDatum) Table(newTableName string) *routeDatum {
	r.routeDatumDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r routeDatum) As(alias string) *routeDatum {
	r.routeDatumDo.DO = *(r.routeDatumDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *routeDatum) updateTableName(table string) *routeDatum {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt32(table, "id")
	r.RouteName = field.NewString(table, "route_name")
	r.MineName = field.NewString(table, "mine_name")
	r.DestinationRegion = field.NewString(table, "destination_region")
	r.RouteDistance = field.NewFloat32(table, "route_distance")
	r.AvgShipmentUnitPrice = field.NewFloat32(table, "avg_shipment_unit_price")
	r.ShipmentVolume = field.NewFloat32(table, "shipment_volume")
	r.Timestamp = field.NewTime(table, "timestamp")

	r.fillFieldMap()

	return r
}

func (r *routeDatum) WithContext(ctx context.Context) *routeDatumDo {
	return r.routeDatumDo.WithContext(ctx)
}

func (r routeDatum) TableName() string { return r.routeDatumDo.TableName() }

func (r routeDatum) Alias() string { return r.routeDatumDo.Alias() }

func (r routeDatum) Columns(cols ...field.Expr) gen.Columns { return r.routeDatumDo.Columns(cols...) }

func (r *routeDatum) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *routeDatum) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 8)
	r.fieldMap["id"] = r.ID
	r.fieldMap["route_name"] = r.RouteName
	r.fieldMap["mine_name"] = r.MineName
	r.fieldMap["destination_region"] = r.DestinationRegion
	r.fieldMap["route_distance"] = r.RouteDistance
	r.fieldMap["avg_shipment_unit_price"] = r.AvgShipmentUnitPrice
	r.fieldMap["shipment_volume"] = r.ShipmentVolume
	r.fieldMap["timestamp"] = r.Timestamp
}

func (r routeDatum) clone(db *gorm.DB) routeDatum {
	r.routeDatumDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r routeDatum) replaceDB(db *gorm.DB) routeDatum {
	r.routeDatumDo.ReplaceDB(db)
	return r
}

type routeDatumDo struct{ gen.DO }

func (r routeDatumDo) Debug() *routeDatumDo {
	return r.withDO(r.DO.Debug())
}

func (r routeDatumDo) WithContext(ctx context.Context) *routeDatumDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r routeDatumDo) ReadDB() *routeDatumDo {
	return r.Clauses(dbresolver.Read)
}

func (r routeDatumDo) WriteDB() *routeDatumDo {
	return r.Clauses(dbresolver.Write)
}

func (r routeDatumDo) Session(config *gorm.Session) *routeDatumDo {
	return r.withDO(r.DO.Session(config))
}

func (r routeDatumDo) Clauses(conds ...clause.Expression) *routeDatumDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r routeDatumDo) Returning(value interface{}, columns ...string) *routeDatumDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r routeDatumDo) Not(conds ...gen.Condition) *routeDatumDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r routeDatumDo) Or(conds ...gen.Condition) *routeDatumDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r routeDatumDo) Select(conds ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r routeDatumDo) Where(conds ...gen.Condition) *routeDatumDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r routeDatumDo) Order(conds ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r routeDatumDo) Distinct(cols ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r routeDatumDo) Omit(cols ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r routeDatumDo) Join(table schema.Tabler, on ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r routeDatumDo) LeftJoin(table schema.Tabler, on ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r routeDatumDo) RightJoin(table schema.Tabler, on ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r routeDatumDo) Group(cols ...field.Expr) *routeDatumDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r routeDatumDo) Having(conds ...gen.Condition) *routeDatumDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r routeDatumDo) Limit(limit int) *routeDatumDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r routeDatumDo) Offset(offset int) *routeDatumDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r routeDatumDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *routeDatumDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r routeDatumDo) Unscoped() *routeDatumDo {
	return r.withDO(r.DO.Unscoped())
}

func (r routeDatumDo) Create(values ...*model.RouteDatum) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r routeDatumDo) CreateInBatches(values []*model.RouteDatum, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r routeDatumDo) Save(values ...*model.RouteDatum) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r routeDatumDo) First() (*model.RouteDatum, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteDatum), nil
	}
}

func (r routeDatumDo) Take() (*model.RouteDatum, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteDatum), nil
	}
}

func (r routeDatumDo) Last() (*model.RouteDatum, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteDatum), nil
	}
}

func (r routeDatumDo) Find() ([]*model.RouteDatum, error) {
	result, err := r.DO.Find()
	return result.([]*model.RouteDatum), err
}

func (r routeDatumDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RouteDatum, err error) {
	buf := make([]*model.RouteDatum, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r routeDatumDo) FindInBatches(result *[]*model.RouteDatum, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r routeDatumDo) Attrs(attrs ...field.AssignExpr) *routeDatumDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r routeDatumDo) Assign(attrs ...field.AssignExpr) *routeDatumDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r routeDatumDo) Joins(fields ...field.RelationField) *routeDatumDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r routeDatumDo) Preload(fields ...field.RelationField) *routeDatumDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r routeDatumDo) FirstOrInit() (*model.RouteDatum, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteDatum), nil
	}
}

func (r routeDatumDo) FirstOrCreate() (*model.RouteDatum, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RouteDatum), nil
	}
}

func (r routeDatumDo) FindByPage(offset int, limit int) (result []*model.RouteDatum, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r routeDatumDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r routeDatumDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r routeDatumDo) Delete(models ...*model.RouteDatum) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *routeDatumDo) withDO(do gen.Dao) *routeDatumDo {
	r.DO = *do.(*gen.DO)
	return r
}
