// Copyright 2025 okTurtles Foundation. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_25

import (
	"xorm.io/xorm"
)

func AddSubjectToRepository(x *xorm.Engine) error {
	type Repository struct {
		Subject string `xorm:"VARCHAR(255) NOT NULL DEFAULT ''"`
	}

	if err := x.Sync(new(Repository)); err != nil {
		return err
	}

	// Set subject to the current name for all existing repositories
	_, err := x.Exec("UPDATE repository SET subject = name WHERE subject = '' OR subject IS NULL")
	return err
}
