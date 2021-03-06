// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package cat

import (
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// AliasesOption is a non-required Aliases option that gets applied to an HTTP request.
type AliasesOption func(r *transport.Request)

// WithAliasesName - a comma-separated list of alias names to return.
func WithAliasesName(name []string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesFormat - a short version of the Accept header, e.g. json, yaml.
func WithAliasesFormat(format string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesH - comma-separated list of column names to display.
func WithAliasesH(h []string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesHelp - return help information.
func WithAliasesHelp(help bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesLocal - return local information, do not retrieve the state from master node (default: false).
func WithAliasesLocal(local bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesMasterTimeout - explicit operation timeout for connection to master node.
func WithAliasesMasterTimeout(masterTimeout time.Duration) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesS - comma-separated list of column names or column aliases to sort by.
func WithAliasesS(s []string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesV - verbose mode. Display column headers.
func WithAliasesV(v bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesErrorTrace - include the stack trace of returned errors.
func WithAliasesErrorTrace(errorTrace bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesFilterPath - a comma-separated list of filters used to reduce the respone.
func WithAliasesFilterPath(filterPath []string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesHuman - return human readable values for statistics.
func WithAliasesHuman(human bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesIgnore - ignores the specified HTTP status codes.
func WithAliasesIgnore(ignore []int) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesPretty - pretty format the returned JSON response.
func WithAliasesPretty(pretty bool) AliasesOption {
	return func(r *transport.Request) {
	}
}

// WithAliasesSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithAliasesSourceParam(sourceParam string) AliasesOption {
	return func(r *transport.Request) {
	}
}

// Aliases - see https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cat-alias.html for more info.
//
// options: optional parameters.
func (c *Cat) Aliases(options ...AliasesOption) (*AliasesResponse, error) {
	req := c.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := c.transport.Do(req)
	return &AliasesResponse{resp}, err
}

// AliasesResponse is the response for Aliases.
type AliasesResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *AliasesResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
