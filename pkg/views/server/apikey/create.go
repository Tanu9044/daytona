// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package apikey

import (
	"errors"
	"log"

	"github.com/daytonaio/daytona/pkg/views"

	"github.com/charmbracelet/huh"
)

func ApiKeyCreationView(name *string, saveToDefaultProfile *bool) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Value(name).
				Validate(func(str string) error {
					if str == "" {
						return errors.New("name can not be blank")
					}
					return nil
				}),
			huh.NewConfirm().
				Title("Save to default profile automatically").
				Value(saveToDefaultProfile),
		),
	).WithTheme(views.GetCustomTheme())

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
}