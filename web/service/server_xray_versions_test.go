package service

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetXrayVersionsPrefersConfiguredDefault(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[
			{"tag_name":"v26.4.25"},
			{"tag_name":"v26.1.23"},
			{"tag_name":"v26.5.3"},
			{"tag_name":"v26.4.30"}
		]`))
	}))
	defer srv.Close()

	oldClient := xrayVersionsClient
	oldURL := xrayReleasesURL
	xrayVersionsClient = &http.Client{Timeout: 2 * time.Second}
	xrayReleasesURL = srv.URL
	defer func() {
		xrayVersionsClient = oldClient
		xrayReleasesURL = oldURL
	}()

	got, err := (&ServerService{}).GetXrayVersions()
	if err != nil {
		t.Fatalf("GetXrayVersions() error = %v", err)
	}

	if len(got) != 3 {
		t.Fatalf("GetXrayVersions() len = %d, want 3, values=%v", len(got), got)
	}
	if got[0] != preferredDefaultXrayVersion {
		t.Fatalf("GetXrayVersions()[0] = %q, want %q", got[0], preferredDefaultXrayVersion)
	}
	if got[1] != "v26.4.25" || got[2] != "v26.4.30" {
		t.Fatalf("unexpected order after preferred version: %v", got)
	}
}

func TestGetXrayVersionsKeepsOrderWhenPreferredMissing(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[
			{"tag_name":"v26.4.25"},
			{"tag_name":"v26.4.30"}
		]`))
	}))
	defer srv.Close()

	oldClient := xrayVersionsClient
	oldURL := xrayReleasesURL
	xrayVersionsClient = &http.Client{Timeout: 2 * time.Second}
	xrayReleasesURL = srv.URL
	defer func() {
		xrayVersionsClient = oldClient
		xrayReleasesURL = oldURL
	}()

	got, err := (&ServerService{}).GetXrayVersions()
	if err != nil {
		t.Fatalf("GetXrayVersions() error = %v", err)
	}

	if len(got) != 2 {
		t.Fatalf("GetXrayVersions() len = %d, want 2", len(got))
	}
	if got[0] != "v26.4.25" || got[1] != "v26.4.30" {
		t.Fatalf("unexpected order: %v", got)
	}
}
