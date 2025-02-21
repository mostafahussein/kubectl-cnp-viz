package plugin

import (
	"fmt"
	"path/filepath"

	"github.com/playwright-community/playwright-go"
)

func captureDiagram(translateX, translateY, scale float64, outputDir string) error {
	yamlFilePath := filepath.Join(outputDir, yamlFileName)
	screenshotPath := filepath.Join(outputDir, diagramFileName)
	if err := initPlaywright(); err != nil {
		return fmt.Errorf("failed to initialize Playwright: %w", err)
	}
	defer func() {
		browser.Close()
		pw.Stop()
	}()
	context, err := browser.NewContext()
	if err != nil {
		return fmt.Errorf("failed to create a browser context: %w", err)
	}

	page, err := context.NewPage()
	if err != nil {
		return fmt.Errorf("failed to create a new page: %w", err)
	}

	if err := page.SetViewportSize(2048, 1280); err != nil {
		return fmt.Errorf("failed to set the viewport size: %w", err)
	}

	if _, err := page.Goto(networkpolicySite); err != nil {
		return fmt.Errorf("failed to navigate to the editor: %w", err)
	}

	startButton := page.Locator(`button:has-text("Start")`)
	if err := startButton.Click(); err != nil {
		return fmt.Errorf("failed to click the 'Start' button: %w", err)
	}

	fileInput := page.Locator(`input[type="file"]`)
	if err := fileInput.SetInputFiles(yamlFilePath); err != nil {
		return fmt.Errorf("failed to upload the YAML file: %w", err)
	}

	if _, err := page.Evaluate(`() => {
			document.querySelector('#app > div > div[class^="styles_panel"]').style.top = "100%";
			const slackButton = document.querySelector('#app > div > div[class^="styles_panel"] > div[class^="styles_slackButton"]');
			if (slackButton) slackButton.style.display = "none";
	}`, nil); err != nil {
		return fmt.Errorf("failed to adjust the UI elements: %w", err)
	}

	if _, err := page.Evaluate(fmt.Sprintf(`() => {
			const g = document.querySelector('#app > div > div[class^="styles_map"] > svg > g');
			if (g) g.setAttribute('transform', 'translate(%f, %f) scale(%f)');
	}`, translateX, translateY, scale), nil); err != nil {
		return fmt.Errorf("failed to transform the G element: %w", err)
	}

	_, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(screenshotPath),
	})
	if err != nil {
		return fmt.Errorf("failed to capture the screenshot: %w", err)
	}
	fmt.Printf("INFO: Screenshot saved at %s\n", screenshotPath)

	return nil
}
