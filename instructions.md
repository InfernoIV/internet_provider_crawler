The 'go test' command will looks for and runs functions that end in _test.go
Build your program: 'go build -o crawler'
Run your program: './crawler'

Using 'go build' and './crawler' is a bit cumbersome. Building is best for production code, because it means the person running your program doesn't need to have Go installed.
However, for development, you can use 'go run .' to compile and run your program in one step.


https://dave.cheney.net/2019/05/07/prefer-table-driven-tests
https://en.wikipedia.org/wiki/Edge_case
https://en.wikipedia.org/wiki/Test-driven_development
