package go_unit_test_bootcamp

import "testing"

// Simple test - verifies that a single repo can be indexed correctly
func TestIndexGitHubRepositoriesByOrgBasic(t *testing.T) {
	repos := []*GitHubRepository{
		{
			Organization: "liatrio",
			Repository:   "devops-bootcamp",
			Url:          "https://github.com/liatrio/devops-bootcamp",
			License:      "MIT",
		},
	}

	result := IndexGitHubRepositoriesByOrg(repos)

	if k := len(result); k != 1 {
		t.Errorf("expected result map to have one key, but it has %d keys", k)
	}

	if n := len(result["liatrio"]); n != 1 {
		t.Errorf("expected liatrio organization to have one repository, but it has %d", n)
	}
}

// ------ Test-Simple test ------

// More complicated test - verify that multiple repos are indexed correctly
func TestIndexGitHubRepositoriesByOrgMultiple(t *testing.T) {
	repos := []*GitHubRepository{
		{
			Organization: "liatrio",
			Repository:   "devops-bootcamp",
			Url:          "https://github.com/liatrio/devops-bootcamp",
			License:      "MIT",
		},
		{
			Organization: "kubernetes",
			Repository:   "kubernetes",
			Url:          "https://github.com/kubernetes/kubernetes",
			License:      "Apache-2.0",
		},
		{
			Organization: "liatrio",
			Repository:   "lead-terraform",
			Url:          "https://github.com/liatrio/lead-terraform",
			License:      "MIT",
		},
	}
	result := IndexGitHubRepositoriesByOrg(repos)

	if k := len(result); k != 2 {
		t.Errorf("expected result map to have two keys, but it has %d keys", k)
	}

	if n := len(result["liatrio"]); n != 2 {
		t.Errorf("expected liatrio organization to have two repositories, but it has %d", n)
	}

	if n := len(result["kubernetes"]); n != 1 {
		t.Errorf("expected kubernetes organization to have two repositories, but it has %d", n)
	}
}

// ----- Test-Complicated test -----

// Edge case test - verify that zero input repositories results in an empty map
func TestIndexGitHubRepositoriesByOrgEmpty(t *testing.T) {
	var repos []*GitHubRepository

	result := IndexGitHubRepositoriesByOrg(repos)

	if k := len(result); k != 0 {
		t.Errorf("expected result mapt to have zero keys, but it has %d keys", k)
	}
}

// ----- Test-Edge case test -----
