# Creating a Golang Project

## No dependencies

> [Project](../projects/hello)

1. Create a new folder for your project, such as `mkdir hello`.

2. Enable dependency tracking, `go mod init <name>`.

3. Declare a `main` package.

4. Run code by running the following command inside the project folder: `go run .`

## With dependencies

> [Project](../projects/proverb)

If you want to use an external module such as `rsc.io/quote`, you can do the following:

1. Search for desired package in https://pkg.go.dev/

2. Import the `rsc.io/quote` package and add a call to any of its functions.

3. Add new module requirements and sums. Go will add the `quote` module as a requirement, as well as a `go.sum` for use in authenticating the module with `go mod tidy`.
