package bitbu

import (
	"strings"
	"testing"
)

// type DataBucketTest struct {
// 	Name string
// 	DefaultDataBucket
// }

func TestGenUpdateCQLstatement(t *testing.T) {
	defaultDataBucket := NewDefaultDataBucket()
	updateCQL, _ := genUpdateCQLstatement(&defaultDataBucket)
	if len(updateCQL) == 0 {
		t.Error("Update CQL string can't be empty")
	}
	if !strings.HasPrefix(updateCQL, CQLTokenUpdate) {
		t.Errorf("Update query is not starting with 'Update '. updateCQL-'%s'", updateCQL)
	}

	if strings.Count(updateCQL, CQLTokenSet) < 1 {
		t.Errorf("Update CQL should contain SET keyword. CQL='%s'", updateCQL)
	}
	err := defaultDataBucket.SetFieldValue("Name", "shreesha")
	if err != nil {
		t.Fatal(err)
	}
	updateCQL, _ = genUpdateCQLstatement(&defaultDataBucket)
	if strings.Count(updateCQL, "Name = ?") != 1 {
		t.Errorf("Update CQL should contain update coloumn 'Name'. CQL='%s'", updateCQL)

	}

}
