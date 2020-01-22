package gomockextras

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/mock/gomock"
)

func StringContaining(s string) gomock.Matcher {
	return &stringContainingMatcher{substr: s}
}

type stringContainingMatcher struct {
	substr string
}

var _ gomock.Matcher = stringContainingMatcher{}

func (s stringContainingMatcher) String() string {
	return fmt.Sprintf("a string containing `%s`", s.substr)
}

func (s stringContainingMatcher) Matches(x interface{}) bool {
	if x == nil {
		return false
	}

	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		if v.Len() > 0 {
			return s.Matches(v.Index(0).Interface())
		} else {
			return false
		}

	case reflect.String:
		return strings.Contains(x.(string), s.substr)

	}

	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	if v.Type().ConvertibleTo(stringerType) {
		return s.Matches(v.Interface().(fmt.Stringer).String())
	}

	return false
}
