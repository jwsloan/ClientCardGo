# Screenshot Manager for System Tests

This package provides functionality similar to Capybara's screenshot feature for our Go system tests using chromedp.

## Features

- Automatically captures and names screenshots during tests
- Organizes screenshots by test name and timestamp
- Creates animated GIFs from test sequences
- Cleans up screenshots after tests if desired

## Usage

```go
func TestSignupFlow(t *testing.T) {
    // Create Chrome instance
    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    // Create screenshot manager
    manager, err := screenshot.NewManager("signup_flow")
    if err != nil {
        t.Fatal(err)
    }
    defer manager.Cleanup() // Optional: remove screenshots after test

    // Your test steps...
    chromedp.Run(ctx,
        // Navigate to page
        chromedp.Navigate("http://localhost:3000/signup"),
        
        // Take screenshot
        chromedp.ActionFunc(func(ctx context.Context) error {
            return manager.Capture(ctx, "initial_page")
        }),
        
        // More test steps...
    )

    // Create GIF from all screenshots
    if err := manager.CreateGif(); err != nil {
        t.Fatal(err)
    }
}
```

## Output

Screenshots are saved in `tmp/test_screenshots/<test_name>_<timestamp>/`:
- `001_initial_page.png`
- `002_filled_form.png`
- `003_after_submit.png`
- `animation.gif`

## Dependencies

Required Go packages:
- github.com/chromedp/chromedp
- github.com/disintegration/imaging
