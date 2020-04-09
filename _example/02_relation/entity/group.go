// Code generated by eevee. DO NOT EDIT!

package entity

type Group struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type Groups []*Group

func (e Groups) IDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.ID)
	}
	return values
}

func (e Groups) Names() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Name)
	}
	return values
}