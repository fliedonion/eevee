// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"context"
	"daoplugin/dao"
	"daoplugin/entity"
	"daoplugin/model"
	"database/sql"

	"golang.org/x/xerrors"
)

type Skill interface {
	ToModel(*entity.Skill) *model.Skill
	ToModels(entity.Skills) *model.Skills
	FindAll(context.Context) (*model.Skills, error)
	FindByID(context.Context, uint64) (*model.Skill, error)
	FindByIDs(context.Context, []uint64) (*model.Skills, error)
	Count(context.Context) (int64, error)
}

type SkillImpl struct {
	skillDAO dao.Skill
	repo     Repository
}

func NewSkill(ctx context.Context, tx *sql.Tx) *SkillImpl {
	return &SkillImpl{skillDAO: dao.NewSkill(ctx, tx)}
}

func (r *SkillImpl) ToModel(value *entity.Skill) *model.Skill {
	return r.createCollection(entity.Skills{value}).First()
}

func (r *SkillImpl) ToModels(values entity.Skills) *model.Skills {
	return r.createCollection(values)
}

func (r *SkillImpl) FindAll(a0 context.Context) (*model.Skills, error) {
	values, err := r.skillDAO.FindAll(a0)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindAll: %w", err)
	}
	collection := r.createCollection(values)
	return collection, nil
}

func (r *SkillImpl) FindByID(a0 context.Context, a1 uint64) (*model.Skill, error) {
	value, err := r.skillDAO.FindByID(a0, a1)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindByID: %w", err)
	}
	if value == nil {
		return nil, nil
	}
	v := r.createCollection(entity.Skills{value}).First()
	return v, nil
}

func (r *SkillImpl) FindByIDs(a0 context.Context, a1 []uint64) (*model.Skills, error) {
	values, err := r.skillDAO.FindByIDs(a0, a1)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindByIDs: %w", err)
	}
	collection := r.createCollection(values)
	return collection, nil
}

func (r *SkillImpl) Count(a0 context.Context) (r0 int64, r1 error) {
	r0, r1 = r.skillDAO.Count(a0)
	if r1 != nil {
		r1 = xerrors.Errorf("failed to Count: %w", r1)
	}
	return
}

func (r *SkillImpl) createCollection(entities entity.Skills) *model.Skills {
	values := model.NewSkills(entities)
	for i := 0; i < len(entities); i += 1 {
		values.Add(r.create(entities[i], values))
	}
	return values
}

func (r *SkillImpl) create(entity *entity.Skill, values *model.Skills) *model.Skill {
	value := model.NewSkill(entity, r.skillDAO)
	value.SetConverter(r.repo.(model.ModelConverter))
	return value
}