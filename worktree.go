package git

import (
	"strings"
	"github.com/go-git/go-git/v5/plumbing/format/index"
)

func (w *Worktree) isIgnoreCase() bool {
	cfg, err := w.r.Config()
	if err != nil {
		return false
	}
	return cfg.Core.IgnoreCase
}

func (w *Worktree) removeStaleEntry(idx *index.Index, path string) {
	if !w.isIgnoreCase() {
		return
	}
	for _, entry := range idx.Entries {
		if strings.EqualFold(entry.Name, path) && entry.Name != path {
			idx.Remove(entry.Name)
		}
	}
}