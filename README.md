# go-librariesio


go-librariesio is a Go client library for accessing the
[libraries.io][libraries.io] API.


## Installation

``go get -u github.com/hackebrot/go-librariesio``


## libraries.io API

Connecting to the [libraries.io API][api] with **go-librariesio** requires
a [private API key][api_key].

## Usage

```go
// Create new API client with your API key
c := librariesio.NewClient("... your API key ...")

// Create a new context (with a timeout if you want)
ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
defer cancel()

// Request information about a project using the client
project, _, err := c.GetProject(ctx, "pypi", "cookiecutter")

if err != nil {
    fmt.Fprintf(os.Stderr, "%v\n", err)
    os.Exit(1)
}

// All structs for API resources use pointer values.
// If you expect fields to not be returned by the API
// make sure to check for nil values before deferencing.
fmt.Fprintf(
    os.Stdout,
    "name: %s\nversion: %s\nlanguage: %s\n",
    *project.Name,
    *project.LatestReleaseNumber,
    *project.Language,
)
```

## License

Distributed under the terms of the [MIT License][MIT], amelia is free and
open source software.


## Code of Conduct

Please note that this project is released with a
[Contributor Code of Conduct][Code of Conduct].

By participating in this project you agree to abide by its terms.

[api_key]: https://libraries.io/account
[api]: https://libraries.io/api
[Code of Conduct]: code_of_conduct.md
[libraries.io]: https://libraries.io/
[MIT]: LICENSE
