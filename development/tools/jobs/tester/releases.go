package tester

// List of currently supported releases
// Please always make it up to date
// When we removing support for given version, there remove
// its entry also here.
var (
	Release14 = mustParse("1.4")
	Release13 = mustParse("1.3")
	Release12 = mustParse("1.2")
)

// GetAllKymaReleaseBranches returns all supported kyma release branches
func GetAllKymaReleaseBranches() []*SupportedRelease {
	return []*SupportedRelease{
		Release14,
		Release13,
		Release12,
	}
}


