package enumerable

import (
	"testing"
)

func TestJoin(t *testing.T) {
	enumerator := Join(
		Slice([]Person{{"Foo"}, {"Bar"}}),
		Slice([]Pet{{"FooPet", "Foo"}, {"BarPet", "Bar"}}),
		func(person Person) string { return person.Name },
		func(pet Pet) string { return pet.Owner },
		func(person Person, pet Pet) string {
			return person.Name + ":" + pet.Name
		},
	).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, "Foo:FooPet", enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, "Bar:BarPet", enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}
