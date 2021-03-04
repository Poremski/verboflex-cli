package verboflex

import "runtime"

// GetAppVersion returns the current version number.
func GetAppVersion() string {
	const appVersion = "1.0.0"

	switch runtime.GOOS {
	case "darwin":
		return appVersion
	case "linux":
		return appVersion
	case "windows":
		return appVersion
	default:
		return appVersion
	}
}
