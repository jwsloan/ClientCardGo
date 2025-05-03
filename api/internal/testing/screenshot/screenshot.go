package screenshot

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/disintegration/imaging"
)

type Manager struct {
	testName    string
	screensDir  string
	screenshots []string
}

// NewManager creates a new screenshot manager for a test
func NewManager(testName string) (*Manager, error) {
	// Clean test name for filesystem
	safeName := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			return r
		}
		return '_'
	}, testName)

	// Create directory for this test's screenshots
	timestamp := time.Now().Format("20060102_150405")
	screensDir := filepath.Join("tmp", "test_screenshots", fmt.Sprintf("%s_%s", safeName, timestamp))
	if err := os.MkdirAll(screensDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create screenshots directory: %w", err)
	}

	return &Manager{
		testName:    safeName,
		screensDir:  screensDir,
		screenshots: make([]string, 0),
	}, nil
}

// Capture takes a screenshot with chromedp and saves it
func (m *Manager) Capture(ctx context.Context, name string) error {
	// Take screenshot
	var buf []byte
	if err := chromedp.FullScreenshot(&buf, 100); err != nil {
		return fmt.Errorf("failed to take screenshot: %w", err)
	}

	// Save screenshot
	filename := fmt.Sprintf("%03d_%s.png", len(m.screenshots)+1, name)
	filepath := filepath.Join(m.screensDir, filename)
	if err := os.WriteFile(filepath, buf, 0644); err != nil {
		return fmt.Errorf("failed to save screenshot: %w", err)
	}

	m.screenshots = append(m.screenshots, filepath)
	return nil
}

// CreateGif creates an animated GIF from all screenshots taken
func (m *Manager) CreateGif() error {
	if len(m.screenshots) == 0 {
		return nil
	}

	// Load all images
	var images []image.Image
	for _, path := range m.screenshots {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open screenshot %s: %w", path, err)
		}
		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
			return fmt.Errorf("failed to decode screenshot %s: %w", path, err)
		}

		// Resize image if needed
		if len(images) > 0 {
			firstBounds := images[0].Bounds()
			if img.Bounds() != firstBounds {
				img = imaging.Resize(img, firstBounds.Dx(), firstBounds.Dy(), imaging.Lanczos)
			}
		}

		images = append(images, img)
	}

	// Create GIF
	gifPath := filepath.Join(m.screensDir, "animation.gif")
	if err := imaging.Save(images[0], gifPath); err != nil {
		return fmt.Errorf("failed to create GIF: %w", err)
	}

	return nil
}

// Cleanup removes the screenshot directory
func (m *Manager) Cleanup() error {
	return os.RemoveAll(m.screensDir)
}
