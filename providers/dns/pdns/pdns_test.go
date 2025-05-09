package pdns

import (
	"net/url"
	"testing"

	"github.com/go-acme/lego/v4/platform/tester"
	"github.com/stretchr/testify/require"
)

const envDomain = envNamespace + "DOMAIN"

var envTest = tester.NewEnvTest(
	EnvAPIURL,
	EnvAPIKey).
	WithDomain(envDomain)

func TestNewDNSProvider(t *testing.T) {
	testCases := []struct {
		desc     string
		envVars  map[string]string
		expected string
	}{
		{
			desc: "success",
			envVars: map[string]string{
				EnvAPIKey: "123",
				EnvAPIURL: "http://example.com",
			},
		},
		{
			desc: "missing credentials",
			envVars: map[string]string{
				EnvAPIKey: "",
				EnvAPIURL: "",
			},
			expected: "pdns: some credentials information are missing: PDNS_API_KEY,PDNS_API_URL",
		},
		{
			desc: "missing api key",
			envVars: map[string]string{
				EnvAPIKey: "",
				EnvAPIURL: "http://example.com",
			},
			expected: "pdns: some credentials information are missing: PDNS_API_KEY",
		},
		{
			desc: "missing API URL",
			envVars: map[string]string{
				EnvAPIKey: "123",
				EnvAPIURL: "",
			},
			expected: "pdns: some credentials information are missing: PDNS_API_URL",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			defer envTest.RestoreEnv()
			envTest.ClearEnv()

			envTest.Apply(test.envVars)

			p, err := NewDNSProvider()

			if test.expected == "" {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.NotNil(t, p.config)
			} else {
				require.EqualError(t, err, test.expected)
			}
		})
	}
}

func TestNewDNSProviderConfig(t *testing.T) {
	testCases := []struct {
		desc             string
		apiKey           string
		customAPIVersion int
		host             *url.URL
		expected         string
	}{
		{
			desc:   "success",
			apiKey: "123",
			host:   mustParse("http://example.com"),
		},
		{
			desc:             "success custom API version",
			apiKey:           "123",
			customAPIVersion: 1,
			host:             mustParse("http://example.com"),
		},
		{
			desc:     "missing credentials",
			expected: "pdns: API key missing",
		},
		{
			desc:     "missing API key",
			apiKey:   "",
			host:     mustParse("http://example.com"),
			expected: "pdns: API key missing",
		},
		{
			desc:     "missing host",
			apiKey:   "123",
			expected: "pdns: API URL missing",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			config := NewDefaultConfig()
			config.APIKey = test.apiKey
			config.Host = test.host
			config.APIVersion = test.customAPIVersion

			p, err := NewDNSProviderConfig(config)

			if test.expected == "" {
				require.NoError(t, err)
				require.NotNil(t, p)
				require.NotNil(t, p.config)
			} else {
				require.EqualError(t, err, test.expected)
			}
		})
	}
}

func TestLivePresentAndCleanup(t *testing.T) {
	if !envTest.IsLiveTest() {
		t.Skip("skipping live test")
	}

	envTest.RestoreEnv()
	provider, err := NewDNSProvider()
	require.NoError(t, err)

	err = provider.Present(envTest.GetDomain(), "", "123d==")
	require.NoError(t, err)
	err = provider.Present(envTest.GetDomain(), "", "123e==")
	require.NoError(t, err)

	err = provider.CleanUp(envTest.GetDomain(), "", "123d==")
	require.NoError(t, err)
	err = provider.CleanUp(envTest.GetDomain(), "", "123e==")
	require.NoError(t, err)
}

func mustParse(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return u
}
