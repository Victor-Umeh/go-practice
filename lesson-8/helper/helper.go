package helper

const Company = "Festac Brothers"

func GenUserPronoun(userGender string) string {
	if userGender == "male" {
		return "Mr"
	} else if userGender == "female" {
		return "Mrs"
	} else {
		return "Mental health"
	}

}