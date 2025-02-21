package plugin

import (
	"fmt"

	"github.com/playwright-community/playwright-go"
)

var pw *playwright.Playwright
var browser playwright.Browser

func initPlaywright() error {
	var err error
	pw, err = playwright.Run()
	if err != nil {
		return fmt.Errorf("failed to start Playwright: %w", err)
	}

	browser, err = pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return fmt.Errorf("failed to launch the browser: %w", err)
	}
	return nil
}
