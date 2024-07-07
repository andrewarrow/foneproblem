package models

func Energy(nrg any) string {
	s, _ := nrg.(string)

	first := s[0]
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
