package bitbu

import (
	"fmt"
)

const (
	CQLTokenUpdate = "UPDATE "
	CQLTokenSet    = "SET "
)

type BucketFieldListOptions bool

const (
	ListBucketFieldsUpdated BucketFieldListOptions = true
	ListBucketFieldsAll     BucketFieldListOptions = false
)

func genUpdateCQLstatement(bitB DataBucket) string {
	updateCQL := CQLTokenUpdate
	//fields := strings.Join(bitB.Fields(false), ",")
	assignments := ""
	for i, updatedFieldName := range bitB.Fields(ListBucketFieldsUpdated) {
		fmt.Println("inside genUpdateCQLstatement ")
		assignments = assignments + " " + updatedFieldName + " = :" + fmt.Sprintf("%d", i+1)
	}
	return updateCQL + CQLTokenSet + assignments
}
