package bitbu

import "testing"

func TestNewFieldNameWithBitUsage(t *testing.T) {
	field := "Name"
	usage := "Users"
	want := "Name.Users"
	result := NewFieldNameWithUsage(field, usage)

	if string(result) != want {
		t.Errorf("output of NewFieldNameWithUsage is wrong. want:'%s',result:'%s' ", want, result)
	}

	if n, _ := result.Name(); n != field {
		t.Errorf("FieldNameWithBitUsage.Name() is not behaving correctly. want:'%s',result:'%s' ", field, n)

	}
	if u, _ := result.BitUsage(); u != usage {
		t.Errorf("FieldNameWithBitUsage.BitUsage() is not behaving correctly. want:'%s',result:'%s' ", field, u)

	}

}
func TestFieldNameWithBitUsageValid(t *testing.T) {
	field := "Name"
	usage := "Users"
	validFieldNameWithUsage := NewFieldNameWithUsage(field, usage)
	inValidFieldNameWithUsage := FieldNameWithUsage("Name,Users")
	if ok, _ := validFieldNameWithUsage.Valid(); !ok {
		t.Errorf("FieldNameWithUsage.Valid() is not behaving correctly.'%s':'%v' should be true",
			string(validFieldNameWithUsage), ok)
	}
	if ok, _ := inValidFieldNameWithUsage.Valid(); ok {
		t.Errorf("FieldNameWithUsage.Valid() is not behaving correctly.'%s':'%v' should be false",
			string(inValidFieldNameWithUsage), ok)
	}
}

func TestBucketSetValue(t *testing.T) {
	b := NewDefaultDataBucket()
	b.SetFieldValue("Name", "shreesha")
	if b.Name != "shreesha" {
		t.Error("SetFieldValue is not setting the value")
	}
}
