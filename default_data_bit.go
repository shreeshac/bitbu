package bitbu

import "errors"

var (
	ErrNoSuchField = "Error: No such field"
)

//DefaultDataBit is a default implementation of DataBit
type DefaultDataBit struct {
	//to store sql table name or NoSQL column family name
	Name string

	FieldNames []string
	//for update
	isUpdated bool
}

func (b DefaultDataBit) BitName() string {
	return b.Name
}

func (b DefaultDataBit) Fields() []string {
	return b.FieldNames
}

func (b DefaultDataBit) FieldValue(fieldName string) (interface{}, error) {
	switch fieldName {
	case "Name":
		return b.Name, nil
	}
	return nil, errors.New(ErrNoSuchField)
}

func (b DefaultDataBit) IsUpdated() bool {
	return b.isUpdated
}
func (b DefaultDataBit) SetForUpdate(t bool) error {
	b.isUpdated = t
	return nil
}
func (b DefaultDataBit) SetValue(fieldName string, value interface{}) error {
	switch fieldName {
	case "Name":
		b.Name = value.(string)
	}
	return errors.New(ErrNoSuchField)

}
