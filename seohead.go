// Package seohead is the SEO Head togo plugin: Runtime <head> manager: editable title/meta/OG/Twitter, Person & WebSite JSON-LD, canonical, hreflang alternates and the markdown-twin link for answer engines.
//
// It self-registers a provider on blank-import and mounts its routes onto the
// kernel. The concrete implementation is ported from the fadymondy.com app under
// internal/server — this scaffold wires the provider + a health route.
package seohead

import (
	"net/http"

	"github.com/togo-framework/togo"
)

func init() {
	togo.RegisterProviderFunc("seo-head", togo.PriorityService, func(k *togo.Kernel) error {
		k.Router.Get("/api/seo-head/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"plugin":"seo-head","status":"ok"}`))
		})
		return nil
	})
}
