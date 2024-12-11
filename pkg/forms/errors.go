package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	fieldErrors := e[field]
	if len(fieldErrors) == 0 {
		return ""
	}
	return fieldErrors[0]
}
