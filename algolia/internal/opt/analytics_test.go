// Code generated by go generate. DO NOT EDIT.

package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAnalytics(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected *opt.AnalyticsOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.Analytics(true),
		},
		{
			opts:     []interface{}{opt.Analytics(true)},
			expected: opt.Analytics(true),
		},
		{
			opts:     []interface{}{opt.Analytics(false)},
			expected: opt.Analytics(false),
		},
	} {
		var (
			in  = ExtractAnalytics(c.opts...)
			out opt.AnalyticsOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, *c.expected, out)
	}
}
