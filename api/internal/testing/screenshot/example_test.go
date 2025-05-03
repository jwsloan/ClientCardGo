package screenshot_test

import (
	"context"
	"testing"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/jwsloan/clientcard/api/internal/testing/screenshot"
)

func TestScreenshotManager(t *testing.T) {
	// Skip in short mode
	if testing.Short() {
		t.Skip("skipping screenshot test in short mode")
	}

	// Create Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Create screenshot manager
	manager, err := screenshot.NewManager("example_signup_flow")
	if err != nil {
		t.Fatalf("failed to create screenshot manager: %v", err)
	}
	defer manager.Cleanup()

	// Example test flow
	err = chromedp.Run(ctx,
		// Visit signup page
		chromedp.Navigate("http://localhost:3000/signup"),
		chromedp.Sleep(time.Second), // Wait for page load
		chromedp.ActionFunc(func(ctx context.Context) error {
			return manager.Capture(ctx, "signup_page")
		}),

		// Fill form
		chromedp.SendKeys(`input[name="email"]`, "test@example.com"),
		chromedp.SendKeys(`input[name="password"]`, "password123"),
		chromedp.SendKeys(`input[name="password_confirmation"]`, "password123"),
		chromedp.Sleep(time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return manager.Capture(ctx, "filled_form")
		}),

		// Submit form
		chromedp.Click(`button[type="submit"]`),
		chromedp.Sleep(time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return manager.Capture(ctx, "after_submit")
		}),
	)

	if err != nil {
		t.Fatalf("test flow failed: %v", err)
	}

	// Create GIF from screenshots
	if err := manager.CreateGif(); err != nil {
		t.Fatalf("failed to create GIF: %v", err)
	}
}
