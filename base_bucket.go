package bitbu

type BucketState string

const (
	BucketStateNew      BucketState = "NEW"
	BucketStateModified BucketState = "MODIFIED"
)

type BaseBucket struct {
	state             BucketState
	fields            map[string]DataBucketField
	fieldNames        []string
	changedFieldNames []string
	dataBits          map[string]DataBit
	filterFields      map[string]DataBucketField
	//for update
	isUpdated bool
}

func (b *BaseBucket) AddDataBit(bitUsage string, dataBit DataBit) error {
	b.dataBits[bitUsage] = dataBit
	return nil
}
func (b *BaseBucket) AddField(fieldName string, bitFieldName string, dataBitUsage string, value interface{}) {
	b.fieldNames = append(b.fieldNames, fieldName)
	b.fields[fieldName] = DataBucketField{
		Name: fieldName,
		BitFieldNameWithUsage: []string{bitFieldName + "." + dataBitUsage},
		OldValue:              value,
		Value:                 value,
	}

}
func (b BaseBucket) Fields(changed BucketFieldListOptions) []string {
	if bool(changed) {
		return b.changedFieldNames
	}
	return b.fieldNames
}
func (b BaseBucket) DataBits() map[string]DataBit {

	return b.dataBits
}

func (b *BaseBucket) AddFilter(fieldName string, value interface{}) {

	b.filterFields[fieldName] = DataBucketField{
		Name:  fieldName,
		Value: value,
	}

}

func (b BaseBucket) Filters() map[string]DataBucketField {

	return b.filterFields
}
