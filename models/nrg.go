package models

type NrgPair struct {
	Id    string
	Label string
}

var AllEnergies = []string{"f18", "m18", "h18", "f25", "m25", "h25", "f30", "m30", "h30", "f40", "m40", "h40"}

func AllEnergiesHuman() []NrgPair {
	items := []NrgPair{}
	for _, nrg := range AllEnergies {
		p := NrgPair{nrg, Energy(nrg)}
		items = append(items, p)
	}
	return items
}
func EnergyMap() map[string]bool {
	m := map[string]bool{}

	for _, ngr := range AllEnergies {
		m[ngr] = true
	}
	return m
}

func Energy(nrg any) string {
	s, _ := nrg.(string)

	first := s[0:1]
	flavor := ""
	if first == "m" {
		flavor = "Male"
	} else if first == "f" {
		flavor = "Female"
	} else {
		flavor = "Human"
	}

	second := s[1:]
	if second == "18" {
		return flavor + " 18-24"
	} else if second == "25" {
		return flavor + " 25-29"
	} else if second == "30" {
		return flavor + " 30-39"
	} else if second == "40" {
		return flavor + " 40-49"
	}

	return s
}
