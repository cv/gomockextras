package gomockextras

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringContaining(t *testing.T) {
	require.Equal(t, "", "")

	t.Run("the string 'aaa'", func(t *testing.T) {
		matcher := StringContaining("aaa")

		t.Run("is not contained in nil", func(t *testing.T) {
			require.False(t, matcher.Matches(nil))
		})

		t.Run("is not contained in an empty string", func(t *testing.T) {
			require.False(t, matcher.Matches(""))
		})

		t.Run("is contained in 'baaab'", func(t *testing.T) {
			require.True(t, matcher.Matches("baaab"))
		})
	})

	t.Run("an empty string", func(t *testing.T) {
		matcher := StringContaining("")

		t.Run("is contained in an empty string", func(t *testing.T) {
			require.True(t, matcher.Matches(""))
		})

		t.Run("is contained in 'baaab'", func(t *testing.T) {
			require.True(t, matcher.Matches("baaab"))
		})
	})

	t.Run("a slice of strings", func(t *testing.T) {
		matcher := StringContaining("o")

		t.Run("uses only first value", func(t *testing.T) {
			require.True(t, matcher.Matches([]string{"bob"}))
			require.True(t, matcher.Matches([]string{"bob", "bab"}))

			require.False(t, matcher.Matches([]string{}))
			require.False(t, matcher.Matches([]string{"bab"}))
			require.False(t, matcher.Matches([]string{"bab", "bob"}))
		})
	})

	t.Run("a slice of interface{}", func(t *testing.T) {
		matcher := StringContaining("o")

		t.Run("uses only first value", func(t *testing.T) {
			require.True(t, matcher.Matches([]interface{}{"bob"}))
			require.True(t, matcher.Matches([]interface{}{"bob", "bab"}))

			require.False(t, matcher.Matches([]interface{}{}))
			require.False(t, matcher.Matches([]interface{}{"bab"}))
			require.False(t, matcher.Matches([]interface{}{"bab", "bob"}))
		})
	})

	t.Run("a struct", func(t *testing.T) {
		matcher := StringContaining("o")

		t.Run("implementing fmt.Stringer", func(t *testing.T) {
			require.True(t, matcher.Matches(stringable{foo: "foo"}))
			require.False(t, matcher.Matches(stringable{foo: "bar"}))
		})

		t.Run("not implementing fmt.Stringer", func(t *testing.T) {
			f := struct{}{}
			require.False(t, matcher.Matches(f))
		})
	})

	t.Run("a slice of structs", func(t *testing.T) {
		matcher := StringContaining("o")

		t.Run("implementing fmt.Stringer", func(t *testing.T) {
			require.True(t, matcher.Matches([]stringable{{foo: "foo"}}))
			require.False(t, matcher.Matches([]stringable{{foo: "bar"}}))
		})

		t.Run("not implementing fmt.Stringer", func(t *testing.T) {
			f := struct{}{}
			require.False(t, matcher.Matches(f))
		})
	})

	t.Run("builds reason strings", func(t *testing.T) {
		r := StringContaining("a").String()
		require.Equal(t, r, "a string containing `a`")
	})
}

type stringable struct{ foo string }

var _ fmt.Stringer = &stringable{}

func (s stringable) String() string {
	return s.foo
}
