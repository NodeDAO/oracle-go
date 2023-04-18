// desc:
// @author renshiwei
// Date: 2023/4/6 11:28

package httptool

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net"
	"net/http"
	"time"
)

type HttpTool struct {
	client  *http.Client
	timeout time.Duration

	extraHeaders map[string]string
}

// Error represents an http error.
type Error struct {
	Method     string
	Endpoint   string
	StatusCode int
	Data       []byte
}

// timeout for tests.
var defaultTimeout = 60 * time.Second

func (e Error) Error() string {
	return fmt.Sprintf("%s failed with status %d: %s", e.Method, e.StatusCode, e.Data)
}

func New(ctx context.Context, timeout time.Duration) (*HttpTool, error) {
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   timeout,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        64,
			MaxConnsPerHost:     64,
			MaxIdleConnsPerHost: 64,
			IdleConnTimeout:     600 * time.Second,
		},
	}

	h := &HttpTool{
		client:  client,
		timeout: timeout,
		//extraHeaders: nil,
	}

	return h, nil
}

func (h *HttpTool) GetRequest(ctx context.Context, endpoint string) (io.Reader, error) {
	var timeout time.Duration
	if h.timeout == 0 {
		timeout = defaultTimeout
	} else {
		timeout = h.timeout
	}

	opCtx, cancel := context.WithTimeout(ctx, timeout)
	req, err := http.NewRequestWithContext(opCtx, http.MethodGet, endpoint, nil)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to create GET request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// golang http post EOF 错误
	//req.Close = true

	resp, err := h.client.Do(req)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to call GET endpoint")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		// Nothing found.  This is not an error, so we return nil on both counts.
		cancel()
		return nil, errors.Wrapf(err, "Failed to request. status: %d", http.StatusNotFound)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		cancel()
		return nil, errors.Wrap(err, "failed to read GET response")
	}

	statusFamily := resp.StatusCode / 100
	if statusFamily != 2 {
		cancel()
		return nil, Error{
			Method:     http.MethodGet,
			StatusCode: resp.StatusCode,
			Endpoint:   endpoint,
			Data:       data,
		}
	}
	cancel()

	return bytes.NewReader(data), nil
}
