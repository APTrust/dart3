# README

## About

This is the rewrite of DART. We're moving away from Electron toward a more
maintainable code base.

## Live Development

To run in live development mode, run `wails dev` in the project directory. 
This will run a Vite development server that will provide very fast hot reload 
of your frontend changes. If you want to develop in a browser and have access 
to your Go methods, there is also a dev server that runs on http://localhost:34115. 
Connect to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Testing

You can test the core Go code by running `go test -p 1 ./...` from the top 
level directory. The `-p 1` flag is required because many tests save data to
an in-memory SQLite database, and then clear the table when they're done.

If you run tests in parallel, some tests will be deleting data inserted by
other tests, and you'll get random failures.

Front end tests will come later, when the front end has enough functionality
to test.
