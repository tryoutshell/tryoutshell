package ui

import "github.com/common-nighthawk/go-figure"

var orgASCII = map[string]string{
	"sigstore": `
   _____ _           _
  / ____(_)         | |
 | (___  _  __ _ ___| |_ ___  _ __ ___
  \___ \| |/ _' / __| __/ _ \| '__/ _ \
  ____) | | (_| \__ \ || (_) | | |  __/
 |_____/|_|\__, |___/\__\___/|_|  \___|
            __/ |
           |___/                       `,

	"chainguard": `
   _____ _           _                               _
  / ____| |         (_)                             | |
 | |    | |__   __ _ _ _ __   __ _ _   _  __ _ _ __| |
 | |    | '_ \ / _' | | '_ \ / _' | | | |/ _' | '__| |
 | |____| | | | (_| | | | | | (_| | |_| | (_| | |  | |
  \_____|_| |_|\__,_|_|_| |_|\__, |\__,_|\__,_|_|  |_|
                              __/ |
                             |___/                    `,

	"in-toto": `
  _____        _______    _
 |_   _|      |__   __|  | |
   | |  _ __     | | ___ | |_ ___
   | | | '_ \    | |/ _ \| __/ _ \
  _| |_| | | |   | | (_) | || (_) |
 |_____|_| |_|   |_|\___/ \__\___/ `,

	"witness": `
 __          ___ _
 \ \        / (_) |
  \ \  /\  / / _| |_ _ __   ___  ___ ___
   \ \/  \/ / | | __| '_ \ / _ \/ __/ __|
    \  /\  /  | | |_| | | |  __/\__ \__ \
     \/  \/   |_|\__|_| |_|\___||___/___/`,
}

//	func GetOrgASCII(orgID string) string {
//		if ascii, ok := orgASCII[orgID]; ok {
//			return ascii
//		}
//		return orgID // Fallback to org name
//	}

// GetOrgASCII generates ASCII art for organization name
func GetOrgASCII(orgID string) string {
	// Map org IDs to display names
	displayNames := map[string]string{
		"sigstore":   "Sigstore",
		"chainguard": "Chainguard",
		"in-toto":    "In-Toto",
		"witness":    "Witness",
	}

	displayName := displayNames[orgID]
	if displayName == "" {
		displayName = orgID
	}

	// Generate ASCII art using go-figure
	// Using "standard" font for clean look
	myFigure := figure.NewFigure(displayName, "standard", true)
	return myFigure.String()
}

// GetAppBanner generates TryOutShell banner
func GetAppBanner() string {
	myFigure := figure.NewFigure("TryOutShell", "standard", true)
	return myFigure.String()
}
