// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_19 //nolint

import (
	"xorm.io/xorm"
)

// AddRepoTypeToRepositoryTable : add RepoType column, setting existing rows to CardTypeTextOnly
func AddRepoTypeToRepositoryTable(x *xorm.Engine) error {
	type Repository struct {
		RepoType int `xorm:"NOT NULL DEFAULT 0"`
	}

	return x.Sync(new(Repository))
}
