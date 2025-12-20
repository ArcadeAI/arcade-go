// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package arcadego_test

import (
	"context"
	"os"
	"testing"

	"github.com/ArcadeAI/arcade-go"
	"github.com/ArcadeAI/arcade-go/internal/testutil"
	"github.com/ArcadeAI/arcade-go/option"
)

func TestManualPagination(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := arcadego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Admin.UserConnections.List(context.TODO(), arcadego.AdminUserConnectionListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, userConnection := range page.Items {
		t.Logf("%+v\n", userConnection.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, userConnection := range page.Items {
			t.Logf("%+v\n", userConnection.ID)
		}
	}
}
