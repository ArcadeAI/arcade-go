// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"
	"strconv"

	"github.com/ArcadeAI/arcade-go/internal/apijson"
	"github.com/ArcadeAI/arcade-go/internal/requestconfig"
	"github.com/ArcadeAI/arcade-go/option"
)

type OffsetPage[T any] struct {
	Items      []T            `json:"items"`
	TotalCount int64          `json:"total_count"`
	Offset     int64          `json:"offset"`
	JSON       offsetPageJSON `json:"-"`
	cfg        *requestconfig.RequestConfig
	res        *http.Response
}

// offsetPageJSON contains the JSON metadata for the struct [OffsetPage[T]]
type offsetPageJSON struct {
	Items       apijson.Field
	TotalCount  apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OffsetPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r offsetPageJSON) RawJSON() string {
	return r.raw
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OffsetPage[T]) GetNextPage() (res *OffsetPage[T], err error) {
	if len(r.Items) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)

	next := r.Offset

	if next < r.TotalCount && next != 0 {
		err = cfg.Apply(option.WithQuery("offset", strconv.FormatInt(next, 10)))
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffsetPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OffsetPage[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OffsetPageAutoPager[T any] struct {
	page *OffsetPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewOffsetPageAutoPager[T any](page *OffsetPage[T], err error) *OffsetPageAutoPager[T] {
	return &OffsetPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OffsetPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Items) == 0 {
		return false
	}
	if r.idx >= len(r.page.Items) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Items) == 0 {
			return false
		}
	}
	r.cur = r.page.Items[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffsetPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *OffsetPageAutoPager[T]) Err() error {
	return r.err
}

func (r *OffsetPageAutoPager[T]) Index() int {
	return r.run
}
