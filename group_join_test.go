package enumerable

import (
	"strings"
	"testing"
)

func TestGroupJoin(t *testing.T) {
	enumerator := GroupJoin(
		Slice([]Person{{"Foo"}, {"Bar"}}),
		Slice([]Pet{{"FooPet", "Foo"}, {"BarPet", "Bar"}}),
		func(person Person) string { return person.Name },
		func(pet Pet) string { return pet.Owner },
		func(person Person, pets IEnumerable[Pet]) string {
			return person.Name + ":" + strings.Join(ToSlice(Select(pets, func(pet Pet) string { return pet.Name })), ",")
		},
	).Enumerator()
	equal(t, true, enumerator.MoveNext())
	equal(t, "Foo:FooPet", enumerator.Current())
	equal(t, true, enumerator.MoveNext())
	equal(t, "Bar:BarPet", enumerator.Current())
	equal(t, false, enumerator.MoveNext())
}
