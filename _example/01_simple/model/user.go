// Code generated by eevee. DO NOT EDIT!

package model

import (
	"context"
	"encoding/json"
	"simple/dao"
	"simple/entity"
	"sort"
	"strconv"

	"golang.org/x/xerrors"
)

type UserFinder interface {
	FindAll(context.Context) (*Users, error)
	FindByID(context.Context, uint64) (*User, error)
	FindByIDs(context.Context, []uint64) (*Users, error)
}

type User struct {
	*entity.User
	userDAO          dao.User
	isAlreadyCreated bool
	savedValue       entity.User
	conv             ModelConverter
}

type Users struct {
	values []*User
}

type UsersCollection []*Users

func NewUser(value *entity.User, userDAO dao.User) *User {
	return &User{
		User:    value,
		userDAO: userDAO,
	}
}

func NewUsers(entities entity.Users) *Users {
	return &Users{values: make([]*User, 0, len(entities))}
}

func (m *Users) newUsers(values []*User) *Users {
	return &Users{values: values}
}

func (m *Users) Each(iter func(*User)) {
	if m == nil {
		return
	}
	for _, value := range m.values {
		iter(value)
	}
}

func (m *Users) EachIndex(iter func(int, *User)) {
	if m == nil {
		return
	}
	for idx, value := range m.values {
		iter(idx, value)
	}
}

func (m *Users) EachWithError(iter func(*User) error) error {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if err := iter(value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Users) EachIndexWithError(iter func(int, *User) error) error {
	if m == nil {
		return nil
	}
	for idx, value := range m.values {
		if err := iter(idx, value); err != nil {
			return xerrors.Errorf("failed to iteration: %w", err)
		}
	}
	return nil
}

func (m *Users) Map(mapFunc func(*User) *User) *Users {
	if m == nil {
		return nil
	}
	mappedValues := []*User{}
	for _, value := range m.values {
		mappedValue := mapFunc(value)
		if mappedValue != nil {
			mappedValues = append(mappedValues, mappedValue)
		}
	}
	return m.newUsers(mappedValues)
}

func (m *Users) Any(cond func(*User) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if cond(value) {
			return true
		}
	}
	return false
}

func (m *Users) Some(cond func(*User) bool) bool {
	return m.Any(cond)
}

func (m *Users) IsIncluded(cond func(*User) bool) bool {
	return m.Any(cond)
}

func (m *Users) All(cond func(*User) bool) bool {
	if m == nil {
		return false
	}
	for _, value := range m.values {
		if !cond(value) {
			return false
		}
	}
	return true
}

func (m *Users) Sort(compare func(*User, *User) bool) {
	if m == nil {
		return
	}
	sort.Slice(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Users) SortStable(compare func(*User, *User) bool) {
	if m == nil {
		return
	}
	sort.SliceStable(m.values, func(i, j int) bool {
		return compare(m.values[i], m.values[j])
	})
}

func (m *Users) Find(cond func(*User) bool) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if cond(value) {
			return value
		}
	}
	return nil
}

func (m *Users) Filter(filter func(*User) bool) *Users {
	if m == nil {
		return nil
	}
	filteredValues := []*User{}
	for _, value := range m.values {
		if filter(value) {
			filteredValues = append(filteredValues, value)
		}
	}
	return m.newUsers(filteredValues)
}

func (m *Users) IsEmpty() bool {
	if m == nil {
		return true
	}
	if len(m.values) == 0 {
		return true
	}
	return false
}

func (m *Users) At(idx int) *User {
	if m == nil {
		return nil
	}
	if idx < 0 {
		return nil
	}
	if len(m.values) > idx {
		return m.values[idx]
	}
	return nil
}

func (m *Users) First() *User {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[0]
	}
	return nil
}

func (m *Users) Last() *User {
	if m == nil {
		return nil
	}
	if len(m.values) > 0 {
		return m.values[len(m.values)-1]
	}
	return nil
}

func (m *Users) Compact() *Users {
	if m == nil {
		return nil
	}
	compactedValues := []*User{}
	for _, value := range m.values {
		if value == nil {
			continue
		}
		compactedValues = append(compactedValues, value)
	}
	return m.newUsers(compactedValues)
}

func (m *Users) Add(args ...*User) *Users {
	if m == nil {
		return nil
	}
	for _, value := range args {
		m.values = append(m.values, value)
	}
	return m
}

func (m *Users) Merge(args ...*Users) *Users {
	if m == nil {
		return nil
	}
	for _, arg := range args {
		for _, value := range arg.values {
			m.values = append(m.values, value)
		}
	}
	return m
}

func (m *Users) Len() int {
	if m == nil {
		return 0
	}
	return len(m.values)
}

func (m *UsersCollection) Merge() *Users {
	if m == nil {
		return nil
	}
	if len(*m) == 0 {
		return nil
	}
	if len(*m) == 1 {
		return (*m)[0]
	}
	values := []*User{}
	for _, collection := range *m {
		for _, value := range collection.values {
			values = append(values, value)
		}
	}
	return (*m)[0].newUsers(values)
}

func (m *User) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '{')
	buf = append(buf, "\"id\":"...)
	buf = strconv.AppendUint(buf, m.ID, 10)
	buf = append(buf, ',')
	buf = append(buf, "\"name\":"...)
	buf = append(buf, strconv.Quote(m.Name)...)
	buf = append(buf, '}')
	return buf, nil
}

func (m *User) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	isWritten := false
	buf = append(buf, '{')
	if option.Exists("id") {
		buf = append(buf, "\"id\":"...)
		buf = strconv.AppendUint(buf, m.ID, 10)
		isWritten = true
	}
	if option.Exists("name") {
		if isWritten {
			buf = append(buf, ',')
		}
		buf = append(buf, "\"name\":"...)
		buf = append(buf, strconv.Quote(m.Name)...)
		isWritten = true
	}
	buf = append(buf, '}')
	return buf, nil
}

func (m *Users) ToJSON(ctx context.Context) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSON(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *Users) ToJSONWithOption(ctx context.Context, option *RenderOption) ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	buf := []byte{}
	buf = append(buf, '[')
	for idx, value := range m.values {
		if idx != 0 {
			buf = append(buf, ',')
		}
		bytes, err := value.ToJSONWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to JSON: %w", err)
		}
		buf = append(buf, bytes...)
	}
	buf = append(buf, ']')
	return buf, nil
}

func (m *User) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *User) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Users) MarshalJSON() ([]byte, error) {
	bytes, err := m.ToJSON(context.Background())
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *Users) MarshalJSONContext(ctx context.Context) ([]byte, error) {
	bytes, err := m.ToJSON(ctx)
	if err != nil {
		return nil, xerrors.Errorf("cannot render to JSON: %w", err)
	}
	return bytes, nil
}

func (m *User) UnmarshalJSON(bytes []byte) error {
	var value struct {
		*entity.User
	}
	if err := json.Unmarshal(bytes, &value); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.User = value.User
	return nil
}

func (m *Users) UnmarshalJSON(bytes []byte) error {
	var values []*User
	if err := json.Unmarshal(bytes, &values); err != nil {
		return xerrors.Errorf("failed to unmarshal: %w", err)
	}
	m.values = values
	return nil
}

func (m *User) ToMap(ctx context.Context) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	value["id"] = m.ID
	value["name"] = m.Name
	return value, nil
}

func (m *User) ToMapWithOption(ctx context.Context, option *RenderOption) (map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := map[string]interface{}{}
	if option.Exists("id") {
		value["id"] = m.ID
	}
	if option.Exists("name") {
		value["name"] = m.Name
	}
	return value, nil
}

func (m *Users) ToMap(ctx context.Context) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMap(ctx)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *Users) ToMapWithOption(ctx context.Context, option *RenderOption) ([]map[string]interface{}, error) {
	if m == nil {
		return nil, nil
	}
	if r, ok := interface{}(m).(BeforeRenderer); ok {
		if err := r.BeforeRender(ctx); err != nil {
			return nil, xerrors.Errorf("failed to BeforeRender: %w", err)
		}
	}
	value := []map[string]interface{}{}
	for _, v := range m.values {
		mapValue, err := v.ToMapWithOption(ctx, option)
		if err != nil {
			return nil, xerrors.Errorf("cannot render to map: %w", err)
		}
		value = append(value, mapValue)
	}
	return value, nil
}

func (m *User) SetConverter(conv ModelConverter) {
	m.conv = conv
}

func (m *User) Create(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	if m.isAlreadyCreated {
		return xerrors.New("this instance has already created")
	}
	if err := m.userDAO.Create(ctx, m.User); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	m.savedValue = *m.User
	m.isAlreadyCreated = true
	return nil
}

func (m *User) Update(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	isRequiredUpdate := false
	if m.savedValue.ID != m.ID {
		isRequiredUpdate = true
	}
	if m.savedValue.Name != m.Name {
		isRequiredUpdate = true
	}
	if !isRequiredUpdate {
		return nil
	}
	if err := m.userDAO.Update(ctx, m.User); err != nil {
		return xerrors.Errorf("failed to Update: %w", err)
	}
	m.savedValue = *m.User
	return nil
}

func (m *User) Delete(ctx context.Context) error {
	if m.userDAO == nil {
		// for testing
		return nil
	}
	if err := m.userDAO.DeleteByID(ctx, m.ID); err != nil {
		return xerrors.Errorf("failed to Delete: %w", err)
	}
	return nil
}

func (m *User) SetAlreadyCreated(isAlreadyCreated bool) {
	m.isAlreadyCreated = isAlreadyCreated
}

func (m *User) SetSavedValue(savedValue *entity.User) {
	m.savedValue = *savedValue
}

func (m *User) Save(ctx context.Context) error {
	if m.isAlreadyCreated {
		if err := m.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}
	if err := m.Create(ctx); err != nil {
		return xerrors.Errorf("failed to Create: %w", err)
	}
	return nil
}

func (m *Users) Create(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Create(ctx); err != nil {
			return xerrors.Errorf("failed to Create: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) Update(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Update(ctx); err != nil {
			return xerrors.Errorf("failed to Update: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) Save(ctx context.Context) error {
	if err := m.EachWithError(func(v *User) error {
		if err := v.Save(ctx); err != nil {
			return xerrors.Errorf("failed to Save: %w", err)
		}
		return nil
	}); err != nil {
		return xerrors.Errorf("interrupt iteration for Users: %w", err)
	}
	return nil
}

func (m *Users) UniqueID() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[uint64]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.ID]; exists {
			return false
		}
		filterMap[value.ID] = struct{}{}
		return true
	})
}

func (m *Users) GroupByID() map[uint64]*Users {
	if m == nil {
		return nil
	}
	values := map[uint64]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.ID]; !exists {
			values[value.ID] = &Users{}
		}
		values[value.ID].Add(value)
	}
	return values
}

func (m *Users) IDs() []uint64 {
	if m == nil {
		return nil
	}
	values := []uint64{}
	for _, value := range m.values {
		values = append(values, value.ID)
	}
	return values
}

func (m *Users) UniqueName() *Users {
	if m == nil {
		return nil
	}
	filterMap := map[string]struct{}{}
	return m.Filter(func(value *User) bool {
		if _, exists := filterMap[value.Name]; exists {
			return false
		}
		filterMap[value.Name] = struct{}{}
		return true
	})
}

func (m *Users) GroupByName() map[string]*Users {
	if m == nil {
		return nil
	}
	values := map[string]*Users{}
	for _, value := range m.values {
		if _, exists := values[value.Name]; !exists {
			values[value.Name] = &Users{}
		}
		values[value.Name].Add(value)
	}
	return values
}

func (m *Users) Names() []string {
	if m == nil {
		return nil
	}
	values := []string{}
	for _, value := range m.values {
		values = append(values, value.Name)
	}
	return values
}

func (m *Users) FirstByID(a0 uint64) *User {
	if m == nil {
		return nil
	}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		return value
	}
	return nil
}

func (m *Users) FilterByID(a0 uint64) *Users {
	if m == nil {
		return nil
	}
	values := []*User{}
	for _, value := range m.values {
		if value.ID != a0 {
			continue
		}
		values = append(values, value)
	}
	return m.newUsers(values)
}
