package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func mustGetEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Starting test…")

	// ----- Config (envs override defaults) -----
	hubURL := mustGetEnv("GRID_HUB_URL", "http://localhost:4444/wd/hub")
	appName := mustGetEnv("GRID_APP_NAME", "it_admin") // must match node stereotype

	// Paths are on the **Proxmox Windows VM**:
	userDataDir := mustGetEnv("CHROME_USER_DATA_DIR",
		`C:\Users\admin\AppData\Local\Google\Chrome\User Data`)
	profileDir := mustGetEnv("CHROME_PROFILE_DIR", `Default`) //

	// ----- Chrome options -----
	chromeOptions := map[string]interface{}{
		"args": []string{
			fmt.Sprintf(`--user-data-dir=%s`, userDataDir),
			fmt.Sprintf(`--profile-directory=%s`, profileDir),
			"--start-maximized",
			"--disable-dev-shm-usage",
			"--no-default-browser-check",
			"--disable-popup-blocking",
		},
	}

	// ----- Capabilities -----
	caps := selenium.Capabilities{
		"browserName":        "chrome",
		"goog:chromeOptions": chromeOptions,
		// Route to a specific node/slot by the custom capability proxmox node advertised
		"nodename:applicationName": appName,
	}

	log.Printf("Connecting to hub: %s", hubURL)
	wd, err := selenium.NewRemote(caps, hubURL)
	if err != nil {
		log.Fatalf("Failed to start Selenium session: %v", err)
	}
	defer func() {
		_ = wd.Quit()
	}()

	target := "http://my.linewize.net"
	log.Printf("Navigating to %s ...", target)
    // Add a short delay to allow the Chrome extension to fully initialize.
    // The wait duration is arbitrary and may be adjusted as needed.
	time.Sleep(5 * time.Second)
	if err := wd.Get(target); err != nil {
		log.Fatalf("Failed to navigate: %v", err)
	}

	title, err := wd.Title()
	if err != nil {
		log.Fatalf("Failed to get title: %v", err)
	}
	log.Printf("Page title: %s", title)

	// Another small pause, then exit
	time.Sleep(2 * time.Second)
	log.Println("Done.")
}
