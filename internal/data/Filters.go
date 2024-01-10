package data

import (
	"github.com/xuche123/bookwise/internal/validator"
	"strings"
)

type Filter struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func ValidateFilters(v *validator.Validator, f Filter) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must not be greater than hundred")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}

func (f Filter) sortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filter) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (f Filter) limit() int {
	return f.PageSize
}
func (f Filter) offset() int {
	return (f.Page - 1) * f.PageSize
}
