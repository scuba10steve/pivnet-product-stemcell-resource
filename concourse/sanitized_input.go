package concourse

// SanitizedSource : sanitize the input source to prevent credential leaking
func SanitizedSource(source Source) map[string]string {
	s := make(map[string]string)

	if source.APIToken != "" {
		s[source.APIToken] = "***REDACTED-PIVNET_API_TOKEN***"
	}

	return s
}