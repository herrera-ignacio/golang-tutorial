# Creating a module

1. Create a module
2. Call code from another module
3. Return and handle errors
4. Return a random greeting (handle data in slices )
5. Return greetings for multiple people (store key/value pairs in a map)
6. Add a test
7. Compile and install application

## Calling another module

For production use, you'd publish your modules from its repository, where Go tools could find it to download it. If you are developing locally, you need to redirect Go tools from its module path (where the module isn't) to the local directory.

```
go mod edit -replace=example.com/greetings=../greetings
go mod tidy
```

The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency.

Your `go.mod` file should look like this:

```
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```

After that you can just run your code `go run .`