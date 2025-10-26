package update

type State struct {
	fields []string
}

func NewState() *State {
	return &State{
		fields: make([]string, 0),
	}
}

func (u *State) HasChanges() bool {
	return len(u.fields) > 0
}

func (u *State) Fields() []string {
	return u.fields
}

func Field[T comparable](state *State, name string, fieldPtr *T, newValue *T) {
	if newValue == nil || fieldPtr == nil {
		return
	}

	if *fieldPtr != *newValue {
		state.fields = append(state.fields, name)
		*fieldPtr = *newValue
	}
}
