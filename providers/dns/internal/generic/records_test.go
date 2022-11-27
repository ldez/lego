package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRecordName_success(t *testing.T) {
	t.Parallel()

	cases := []struct{ fqdn, authZone, expected string }{
		// typical case: sub-domain is controlled on (grand-) parent domain
		{
			fqdn:     "foo.example.com.",
			authZone: "example.com.",
			expected: "foo",
		},
		{
			fqdn:     "bar.foo.example.com.",
			authZone: "foo.example.com.",
			expected: "bar",
		},
		{
			fqdn:     "bar.foo.example.com.",
			authZone: "example.com.",
			expected: "bar.foo",
		},
		// eTLDs are an edge case, but GetRecordName won't treat it
		// special
		{
			fqdn:     "example.com.au.",
			authZone: "au.",
			expected: "example.com",
		},
		// in practice, this shouldn't happen, as fqdn will always
		// be a sub-domain of authZone
		{
			fqdn:     "bar.foo.example.com.",
			authZone: "foo.example.org.",
			expected: "bar",
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.fqdn, func(t *testing.T) {
			t.Parallel()

			name, err := GetRecordName(tc.fqdn, tc.authZone)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, name)
		})
	}
}

func TestGetRecordName_failure(t *testing.T) {
	t.Parallel()

	cases := []struct{ name, fqdn, authZone, err string }{
		{
			name:     "arguments switched",
			fqdn:     "example.com.",
			authZone: "foo.example.com.",
			err:      "-5 is lower than the length of the fqdn (fqdn: example.com., authZone: foo.example.com.)",
		},
		{
			name:     "different sub-domain of same length",
			fqdn:     "foo.example.com.",
			authZone: "bar.example.com.",
			err:      "-1 is lower than the length of the fqdn (fqdn: foo.example.com., authZone: bar.example.com.)",
		},
	}

	for i := range cases {
		tc := cases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			name, err := GetRecordName(tc.fqdn, tc.authZone)
			require.EqualError(t, err, tc.err)
			assert.Equal(t, "", name)
		})
	}
}
