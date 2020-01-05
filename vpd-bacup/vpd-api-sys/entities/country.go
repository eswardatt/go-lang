package entities

type Country struct {
	ID, Name, ShortName, PhoneCode string
}

type State struct {
	ID, Name, CountryID string
}

type City struct {
	ID, Name, StateID string
}
