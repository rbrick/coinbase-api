package coinbase

import (
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Order string

const (
	Ascending  Order = "asc"
	Descending Order = "desc"
)

type PaginatedList interface {
	Next() PaginatedList
	Previous() PaginatedList
	Values() []interface{}
}

type Pagination struct {
	//Limit specifies the number of entries that will be listed
	Limit int `json:"limit,omitempty" argument:"true"`
	//Order is the order of the entries
	Order Order `json:"order,omitempty" argument:"true"`
	//StartingAfter is a resource ID which defines your place in the list
	StartingAfter string `json:"starting_after,omitempty" argument:"true"`
	//EndingBefore is a resource ID which defines your place in the list
	EndingBefore string `json:"ending_before,omitempty" argument:"true"`
	//NextURI is the next page in the list
	NextURI string `json:"next_uri,omitempty"`
	//PreviousURI is the previous page in the list
	PreviousURI string `json:"previous_uri,omitempty"`
}

func (p *Pagination) Encode() url.Values {
	values := url.Values{}
	t := reflect.TypeOf(*p)  // look up the type
	v := reflect.ValueOf(*p) // look up the value

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i) // look up the field

		// check if the field has an argument tags
		if _, ok := f.Tag.Lookup("argument"); ok {
			// Get the entry/url name
			entryName := strings.Split(f.Tag.Get("json"), ",")[0]
			// if it does, get the field value
			valueField := v.Field(i)

			// Check if it's a non nil, non zero value
			if valueField.IsValid() {
				value := valueField.Interface()

				switch v := value.(type) {
				case Order:
					values.Add(entryName, string(v))
				case string:
					if v != "" {
						values.Add(entryName, v)
					}
				case int:
					values.Add(entryName, strconv.Itoa(v))
				}
			}
		}
	}

	return values
}
