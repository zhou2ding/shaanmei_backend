// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:                 db,
		CarrierSeriesDatum: newCarrierSeriesDatum(db, opts...),
		MineSeriesDatum:    newMineSeriesDatum(db, opts...),
		PredictionDatum:    newPredictionDatum(db, opts...),
		RouteDatum:         newRouteDatum(db, opts...),
		ShipmentPrice:      newShipmentPrice(db, opts...),
		TimeseriesDatum:    newTimeseriesDatum(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	CarrierSeriesDatum carrierSeriesDatum
	MineSeriesDatum    mineSeriesDatum
	PredictionDatum    predictionDatum
	RouteDatum         routeDatum
	ShipmentPrice      shipmentPrice
	TimeseriesDatum    timeseriesDatum
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CarrierSeriesDatum: q.CarrierSeriesDatum.clone(db),
		MineSeriesDatum:    q.MineSeriesDatum.clone(db),
		PredictionDatum:    q.PredictionDatum.clone(db),
		RouteDatum:         q.RouteDatum.clone(db),
		ShipmentPrice:      q.ShipmentPrice.clone(db),
		TimeseriesDatum:    q.TimeseriesDatum.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:                 db,
		CarrierSeriesDatum: q.CarrierSeriesDatum.replaceDB(db),
		MineSeriesDatum:    q.MineSeriesDatum.replaceDB(db),
		PredictionDatum:    q.PredictionDatum.replaceDB(db),
		RouteDatum:         q.RouteDatum.replaceDB(db),
		ShipmentPrice:      q.ShipmentPrice.replaceDB(db),
		TimeseriesDatum:    q.TimeseriesDatum.replaceDB(db),
	}
}

type queryCtx struct {
	CarrierSeriesDatum *carrierSeriesDatumDo
	MineSeriesDatum    *mineSeriesDatumDo
	PredictionDatum    *predictionDatumDo
	RouteDatum         *routeDatumDo
	ShipmentPrice      *shipmentPriceDo
	TimeseriesDatum    *timeseriesDatumDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		CarrierSeriesDatum: q.CarrierSeriesDatum.WithContext(ctx),
		MineSeriesDatum:    q.MineSeriesDatum.WithContext(ctx),
		PredictionDatum:    q.PredictionDatum.WithContext(ctx),
		RouteDatum:         q.RouteDatum.WithContext(ctx),
		ShipmentPrice:      q.ShipmentPrice.WithContext(ctx),
		TimeseriesDatum:    q.TimeseriesDatum.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
