// Code generated by eevee. DO NOT EDIT!

package entity

import (
	"go.knocknote.io/rapidash"
	"golang.org/x/xerrors"
	"time"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Age       int       `json:"age"`
	SkillID   uint64    `json:"skillID"`
	SkillRank int       `json:"skillRank"`
	GroupID   uint64    `json:"groupID"`
	WorldID   uint64    `json:"worldID"`
	FieldID   uint64    `json:"fieldID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Users []*User

func (e Users) IDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.ID)
	}
	return values
}

func (e Users) Names() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Name)
	}
	return values
}

func (e Users) Sexes() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Sex)
	}
	return values
}

func (e Users) Ages() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.Age)
	}
	return values
}

func (e Users) SkillIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.SkillID)
	}
	return values
}

func (e Users) SkillRanks() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.SkillRank)
	}
	return values
}

func (e Users) GroupIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.GroupID)
	}
	return values
}

func (e Users) WorldIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.WorldID)
	}
	return values
}

func (e Users) FieldIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.FieldID)
	}
	return values
}

func (e Users) CreatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.CreatedAt)
	}
	return values
}

func (e Users) UpdatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.UpdatedAt)
	}
	return values
}

func (e *User) Struct() *rapidash.Struct {
	return rapidash.NewStruct("users").
		FieldUint64("id").
		FieldString("name").
		FieldString("sex").
		FieldInt("age").
		FieldUint64("skill_id").
		FieldInt("skill_rank").
		FieldUint64("group_id").
		FieldUint64("world_id").
		FieldUint64("field_id").
		FieldTime("created_at").
		FieldTime("updated_at")
}

func (e *User) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.String("name", e.Name)
	enc.String("sex", e.Sex)
	enc.Int("age", e.Age)
	enc.Uint64("skill_id", e.SkillID)
	enc.Int("skill_rank", e.SkillRank)
	enc.Uint64("group_id", e.GroupID)
	enc.Uint64("world_id", e.WorldID)
	enc.Uint64("field_id", e.FieldID)
	enc.Time("created_at", e.CreatedAt)
	enc.Time("updated_at", e.UpdatedAt)
	if err := enc.Error(); err != nil {
		return xerrors.Errorf("failed to encode: %w", err)
	}
	return nil
}

func (e *Users) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return xerrors.Errorf("failed to encode: %w", err)
		}
	}
	return nil
}

func (e *User) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.Name = dec.String("name")
	e.Sex = dec.String("sex")
	e.Age = dec.Int("age")
	e.SkillID = dec.Uint64("skill_id")
	e.SkillRank = dec.Int("skill_rank")
	e.GroupID = dec.Uint64("group_id")
	e.WorldID = dec.Uint64("world_id")
	e.FieldID = dec.Uint64("field_id")
	e.CreatedAt = dec.Time("created_at")
	e.UpdatedAt = dec.Time("updated_at")
	if err := dec.Error(); err != nil {
		return xerrors.Errorf("failed to decode: %w", err)
	}
	return nil
}

func (e *Users) DecodeRapidash(dec rapidash.Decoder) error {
	decLen := dec.Len()
	values := make(Users, decLen)
	for i := 0; i < decLen; i++ {
		var v User
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return xerrors.Errorf("failed to decode: %w", err)
		}
		values[i] = &v
	}
	*e = values
	return nil
}