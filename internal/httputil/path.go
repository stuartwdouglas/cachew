package httputil

import (
	"strings"

	"github.com/alecthomas/errors"
)

// ValidateRelativePath prevents decoded URL components from escaping their upstream route.
func ValidateRelativePath(path string) error {
	if strings.ContainsAny(path, "\\\x00") {
		return errors.Errorf("unsupported path %q", path)
	}
	for segment := range strings.SplitSeq(strings.TrimSuffix(path, "/"), "/") {
		if segment == "" || segment == "." || segment == ".." {
			return errors.Errorf("unsupported path segment %q", segment)
		}
	}
	return nil
}
