package dao

import (
	"context"
	"daoplugin/entity"
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/xerrors"
)

type Group interface {
	Count(context.Context) (int64, error)
	FindAll(context.Context) (entity.Groups, error)
	FindByID(context.Context, uint64) (*entity.Group, error)
	FindByIDs(context.Context, []uint64) (entity.Groups, error)
}

type GroupImpl struct {
	tx *sql.Tx
}

func NewGroup(ctx context.Context, tx *sql.Tx) Group {
	return &GroupImpl{tx: tx}
}

// generated by eevee
func (d *GroupImpl) Count(ctx context.Context) (r int64, e error) {
	var value int64
	query := "COUNT(*) FROM `groups`"
	if err := d.tx.QueryRowContext(ctx, query).Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return value, nil
}

// generated by eevee
func (d *GroupImpl) FindAll(ctx context.Context) (r entity.Groups, e error) {
	values := entity.Groups{}
	query := "SELECT `id`, `name` FROM `groups`"
	rows, err := d.tx.QueryContext(ctx, query)
	if err != nil {
		return values, xerrors.Errorf("falure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.Group
		if err := rows.Scan(&value.ID, &value.Name); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *GroupImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.Group, e error) {
	var value entity.Group
	query := "SELECT `id`, `name` FROM `groups` WHERE `id` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0).Scan(
		&value.ID,
		&value.Name,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return &value, nil
}

// generated by eevee
func (d *GroupImpl) FindByIDs(ctx context.Context, a0 []uint64) (r entity.Groups, e error) {
	values := entity.Groups{}
	query := "SELECT `id`, `name` FROM `groups` WHERE `id` IN (%s)"
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	selectQuery := fmt.Sprintf(query, strings.Join(placeholders, ", "))
	rows, err := d.tx.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.Group
		if err := rows.Scan(
			&value.ID,
			&value.Name,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}