package strategy

import (
	"net/http"
	"strings"

	"github.com/block/cachew/internal/httputil"
)

func validateGitHubReleaseRequest(restParam string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, name := range []string{"org", "repo", restParam} {
			value := r.PathValue(name)
			if err := httputil.ValidateRelativePath(value); err != nil || (name != restParam && strings.Contains(value, "/")) {
				http.Error(w, "invalid release path", http.StatusBadRequest)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
