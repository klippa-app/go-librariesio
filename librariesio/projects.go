package librariesio

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Project represents a project on libraries.io
type Project struct {
	Description              *string    `json:"description,omitempty"`
	Forks                    *int       `json:"forks,omitempty"`
	Homepage                 *string    `json:"homepage,omitempty"`
	Keywords                 []*string  `json:"keywords,omitempty"`
	Language                 *string    `json:"language,omitempty"`
	LatestReleaseNumber      *string    `json:"latest_release_number,omitempty"`
	LatestReleasePublishedAt *time.Time `json:"latest_release_published_at,omitempty"`
	LatestStableRelease      *Release   `json:"latest_stable_release,omitempty"`
	Name                     *string    `json:"name,omitempty"`
	NormalizedLicenses       []*string  `json:"normalized_licenses,omitempty"`
	LicenseNormalized        *bool      `json:"license_normalized,omitempty"`
	Licenses                 *string    `json:"licenses,omitempty"`
	PackageManagerURL        *string    `json:"package_manager_url,omitempty"`
	Platform                 *string    `json:"platform,omitempty"`
	Rank                     *int       `json:"rank,omitempty"`
	Stars                    *int       `json:"stars,omitempty"`
	Status                   *string    `json:"status,omitempty"`
	Versions                 []*Release `json:"versions,omitempty"`

	// Dependencies are only populated for ProjectDeps
	Dependencies []*ProjectDependency `json:"dependencies,omitempty"`

	// RepositoryURL is only populated for UserProjects
	RepositoryURL *string `json:"repository_url,omitempty"`
}

// Release represents a release of the project
type Release struct {
	Number            *string    `json:"number,omitempty"`
	PublishedAt       *time.Time `json:"published_at,omitempty"`
	SPDXExpression    *string    `json:"spdx_expression,omitempty"`
	ResearchedAt      *time.Time `json:"researched_at,omitempty"`
	RepositorySources *[]string  `json:"repository_sources,omitempty"`
}

// ProjectDependency represents a dependency of the project
type ProjectDependency struct {
	Deprecated   *bool   `json:"deprecated,omitempty"`
	Latest       *string `json:"latest,omitempty"`
	LatestStable *string `json:"latest_stable,omitempty"`
	Name         *string `json:"name,omitempty"`
	Outdated     *bool   `json:"outdated,omitempty"`
	Platform     *string `json:"platform,omitempty"`
	ProjectName  *string `json:"project_name,omitempty"`
	Requirements *string `json:"requirements,omitempty"`
}

// Project returns information about a project and it's versions.
//
// GET https://libraries.io/api/:platform/:name
//
// plat is the platform/package manager of the project
// name is the name of the project on the platform
func (c *Client) Project(ctx context.Context, plat, name string) (*Project, *http.Response, error) {
	urlStr := fmt.Sprintf("%v/%v", plat, name)

	request, err := c.NewRequest("GET", urlStr, nil)

	if err != nil {
		return nil, nil, err
	}

	project := new(Project)
	response, err := c.Do(ctx, request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

// ProjectDeps returns information about a project and it's dependencies.
//
// GET https://libraries.io/api/:platform/:name/:version/dependencies
//
// plat is the platform/package manager of the project
// name is the name of the project on the platform
// ver is the version of the project - pass "latest" for current release
func (c *Client) ProjectDeps(ctx context.Context, plat, name, ver string) (*Project, *http.Response, error) {

	urlStr := fmt.Sprintf("%v/%v/%v/dependencies", plat, name, ver)

	request, err := c.NewRequest("GET", urlStr, nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(Project)

	response, err := c.Do(ctx, request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

// Search returns a slice of projects for the given search string
//
// GET https://libraries.io/api/search?q=amelia
func (c *Client) Search(ctx context.Context, q string) ([]*Project, *http.Response, error) {
	request, err := c.NewRequest("GET", "search", nil)
	if err != nil {
		return nil, nil, err
	}

	// Add query to request
	query := request.URL.Query()
	query.Set("q", q)
	request.URL.RawQuery = query.Encode()

	var projects []*Project

	response, err := c.Do(ctx, request, &projects)
	if err != nil {
		return nil, response, err
	}

	return projects, response, nil
}
