package bitbu

import (
	"strings"
)

const (
	CQLTokenUpdate = " UPDATE "
	CQLTokenSet    = " SET "
	CQLTokenWhere  = " WHERE "
	CQLTokenAnd    = " AND "
	CQLTokenInsert = " INSERT INTO "
	CQLTokenValues = " VALUES "
)

type BucketFieldListOptions bool

const (
	ListBucketFieldsUpdated BucketFieldListOptions = true
	ListBucketFieldsAll     BucketFieldListOptions = false
)

func genUpdateCQLstatement(bitB DataBucket) (string, []interface{}) {
	updateCQL := CQLTokenUpdate
	//fields := strings.Join(bitB.Fields(false), ",")
	assignments := ""
	var values []interface{}
	for _, updatedFieldName := range bitB.Fields(ListBucketFieldsUpdated) {

		assignments = assignments + " " + updatedFieldName + " = ?"
		v, _ := bitB.FieldValue(updatedFieldName)
		values = append(values, v)
	}

	filters := ""
	for _, field := range bitB.FilterFieldNames() {
		if len(filters) == 0 {
			filters = filters + CQLTokenWhere + field + " = ?"
			v := bitB.Filters()[field].Value
			values = append(values, v)
		} else {
			filters = filters + CQLTokenAnd + field + " = ?"
		}
	}
	return updateCQL + bitB.DataBitUsages()[0] + CQLTokenSet + assignments + filters, values
}

func genInsertCQL(b DataBucket) (queries map[string]string, values map[string][]interface{}) {
	queries = make(map[string]string)
	values = make(map[string][]interface{})
	for _, usage := range b.DataBitUsages() {
		bit := b.DataBits()[usage].(BitReader)
		query := CQLTokenInsert + bit.BitName()
		query = query + "(" + strings.Join(bit.Fields(), ",") + ")"
		query = query + CQLTokenValues

		query = query + "(" + genValuesPlaceHolderString(bit.FieldValues()...) + ")"
		queries[usage] = query
		values[usage] = bit.FieldValues()
	}

	return queries, values
}

func genValuesPlaceHolderString(args ...interface{}) string {
	valuesPlaceHolderString := ""
	if len(args) > 0 {
		valuesPlaceHolderString = "?"
	}
	if len(args) > 1 {
		valuesPlaceHolderString = valuesPlaceHolderString + strings.Repeat(",?", len(args)-1)
	}
	return valuesPlaceHolderString
}
