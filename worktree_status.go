package git

import "strings"

func (w *Worktree) statusIgnoreCase(path string, idx *index.Index) bool {
	if !w.isIgnoreCase() {
		return false
	}
	for _, entry := range idx.Entries {
		if strings.EqualFold(entry.Name, path) {
			return true
		}
	}
	return false
}