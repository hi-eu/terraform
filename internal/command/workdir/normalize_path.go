package workdir

import (
	"path/filepath"
)

// NormalizePath attempts to transform the given path so that it's relative
// to the working directory, which is our preferred way to present and store
// paths to files and directories within a configuration so that they can
// be portable to operations in other working directories.
func (d *Dir) NormalizePath(given string) string {
	// We need an absolute version of d.mainDir in order for our "Rel"
	// result to be reliable.
	absMain, err := filepath.Abs(d.mainDir)
	if err != nil {
		// Weird, but okay...
		return filepath.Clean(given)
	}

	if !filepath.IsAbs(given) {
		given = filepath.Join(absMain, given)
	}

	ret, err := filepath.Rel(absMain, given)
	if err != nil {
		// It's not always possible to find a relative path. For example,
		// the given path might be on an entirely separate volume
		// (e.g. drive letter or network share) on a Windows system, which
		// always requires an absolute path.
		return filepath.Clean(given)
	}

	return ret
}
