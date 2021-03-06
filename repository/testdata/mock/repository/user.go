// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"/entity"
	"/model"
	"context"
	"log"
	"reflect"

	"golang.org/x/xerrors"
)

type UserMock struct {
	expect *UserExpect
}

func (r *UserMock) EXPECT() *UserExpect {
	return r.expect
}

func NewUserMock() *UserMock {
	return &UserMock{expect: NewUserExpect()}
}

type UserToModelExpect struct {
	expect        *UserExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(value *entity.User)
	value         *entity.User
	r0            *model.User
}

func (r *UserToModelExpect) Return(r0 *model.User) *UserToModelExpect {
	r.r0 = r0
	return r
}

func (r *UserToModelExpect) Do(action func(value *entity.User)) *UserToModelExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *UserToModelExpect) OutOfOrder() *UserToModelExpect {
	r.isOutOfOrder = true
	return r
}

func (r *UserToModelExpect) AnyTimes() *UserToModelExpect {
	r.isAnyTimes = true
	return r
}

func (r *UserToModelExpect) Times(n int) *UserToModelExpect {
	r.requiredTimes = n
	return r
}

func (r *UserMock) ToModel(value *entity.User) (r0 *model.User) {
	if len(r.expect.toModel) == 0 {
		log.Printf("cannot find mock method for User.ToModel")
		return
	}
	for _, exp := range r.expect.toModel {
		if !reflect.DeepEqual(exp.value, value) {
			continue
		}
		for _, action := range exp.actions {
			action(value)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			log.Printf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	log.Printf("invalid argument User value:[%+v]", value)
	return
}

func (r *UserExpect) ToModel(value *entity.User) *UserToModelExpect {
	exp := &UserToModelExpect{
		actions: []func(value *entity.User){},
		expect:  r,
		value:   value,
	}
	r.toModel = append(r.toModel, exp)
	return exp
}

type UserToModelsExpect struct {
	expect        *UserExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(values entity.Users)
	values        entity.Users
	r0            *model.Users
}

func (r *UserToModelsExpect) Return(r0 *model.Users) *UserToModelsExpect {
	r.r0 = r0
	return r
}

func (r *UserToModelsExpect) Do(action func(values entity.Users)) *UserToModelsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *UserToModelsExpect) OutOfOrder() *UserToModelsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *UserToModelsExpect) AnyTimes() *UserToModelsExpect {
	r.isAnyTimes = true
	return r
}

func (r *UserToModelsExpect) Times(n int) *UserToModelsExpect {
	r.requiredTimes = n
	return r
}

func (r *UserMock) ToModels(values entity.Users) (r0 *model.Users) {
	if len(r.expect.toModels) == 0 {
		log.Printf("cannot find mock method for User.ToModels")
		return
	}
	for _, exp := range r.expect.toModels {
		if !reflect.DeepEqual(exp.values, values) {
			continue
		}
		for _, action := range exp.actions {
			action(values)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			log.Printf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	log.Printf("invalid argument User values:[%+v]", values)
	return
}

func (r *UserExpect) ToModels(values entity.Users) *UserToModelsExpect {
	exp := &UserToModelsExpect{
		actions: []func(values entity.Users){},
		expect:  r,
		values:  values,
	}
	r.toModels = append(r.toModels, exp)
	return exp
}

type UserCreateExpect struct {
	expect        *UserExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, value *entity.User)
	ctx           context.Context
	value         *entity.User
	r0            *model.User
	r1            error
}

func (r *UserCreateExpect) Return(r0 *model.User, r1 error) *UserCreateExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *UserCreateExpect) Do(action func(ctx context.Context, value *entity.User)) *UserCreateExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *UserCreateExpect) OutOfOrder() *UserCreateExpect {
	r.isOutOfOrder = true
	return r
}

func (r *UserCreateExpect) AnyTimes() *UserCreateExpect {
	r.isAnyTimes = true
	return r
}

func (r *UserCreateExpect) Times(n int) *UserCreateExpect {
	r.requiredTimes = n
	return r
}

func (r *UserMock) Create(ctx context.Context, value *entity.User) (r0 *model.User, r1 error) {
	if len(r.expect.create) == 0 {
		r1 = xerrors.New("cannot find mock method for User.Create")
		return
	}
	for _, exp := range r.expect.create {
		if !reflect.DeepEqual(exp.ctx, ctx) {
			continue
		}
		if !reflect.DeepEqual(exp.value, value) {
			continue
		}
		for _, action := range exp.actions {
			action(ctx, value)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument User ctx:[%+v] value:[%+v]", ctx, value)
	return
}

func (r *UserExpect) Create(ctx context.Context, value *entity.User) *UserCreateExpect {
	exp := &UserCreateExpect{
		actions: []func(ctx context.Context, value *entity.User){},
		ctx:     ctx,
		expect:  r,
		value:   value,
	}
	r.create = append(r.create, exp)
	return exp
}

type UserCreatesExpect struct {
	expect        *UserExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, entities entity.Users)
	ctx           context.Context
	entities      entity.Users
	r0            *model.Users
	r1            error
}

func (r *UserCreatesExpect) Return(r0 *model.Users, r1 error) *UserCreatesExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *UserCreatesExpect) Do(action func(ctx context.Context, entities entity.Users)) *UserCreatesExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *UserCreatesExpect) OutOfOrder() *UserCreatesExpect {
	r.isOutOfOrder = true
	return r
}

func (r *UserCreatesExpect) AnyTimes() *UserCreatesExpect {
	r.isAnyTimes = true
	return r
}

func (r *UserCreatesExpect) Times(n int) *UserCreatesExpect {
	r.requiredTimes = n
	return r
}

func (r *UserMock) Creates(ctx context.Context, entities entity.Users) (r0 *model.Users, r1 error) {
	if len(r.expect.creates) == 0 {
		r1 = xerrors.New("cannot find mock method for User.Creates")
		return
	}
	for _, exp := range r.expect.creates {
		if !reflect.DeepEqual(exp.ctx, ctx) {
			continue
		}
		if !reflect.DeepEqual(exp.entities, entities) {
			continue
		}
		for _, action := range exp.actions {
			action(ctx, entities)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument User ctx:[%+v] entities:[%+v]", ctx, entities)
	return
}

func (r *UserExpect) Creates(ctx context.Context, entities entity.Users) *UserCreatesExpect {
	exp := &UserCreatesExpect{
		actions:  []func(ctx context.Context, entities entity.Users){},
		ctx:      ctx,
		entities: entities,
		expect:   r,
	}
	r.creates = append(r.creates, exp)
	return exp
}

type UserExpect struct {
	toModel  []*UserToModelExpect
	toModels []*UserToModelsExpect
	create   []*UserCreateExpect
	creates  []*UserCreatesExpect
}

func NewUserExpect() *UserExpect {
	return &UserExpect{
		create:   []*UserCreateExpect{},
		creates:  []*UserCreatesExpect{},
		toModel:  []*UserToModelExpect{},
		toModels: []*UserToModelsExpect{},
	}
}
