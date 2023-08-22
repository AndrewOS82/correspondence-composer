# Correspondence Composer

## About

This is an application which listens for incoming events indicating that correspondence should be generated for a client. It fetches relevant data for the correspondence from Zinnia enterprise APIs and then generates XML that can be consumed by other services which prepare and send the correspondence.

## Setup

### Installing Go on your machine

Follow [these instructions](https://go.dev/doc/install) for installing go on Linux, Mac or Windows, then proceed to the "Run the app" section.


## Run Instructions

You can run this application in two ways, by installing Go on your machine or via Docker. If you've installed Go on your machine start up app with this command:

`make run`

### Using Docker

If you already have Docker set up and running use the following commands to run the app with Docker:

```
make docker-build
make localdev
```

## Contributing

To contribute to the project:

* Create a branch off main
* Make sure that your code is linted and tested
* Open a pull request

### Linting

Follow the instructions [here](https://golangci-lint.run/usage/install/#local-installation) to install `golangci-lint`. Then run `make lint`.

### Testing

Run `make test`.


## Regenerating types from XSDs

We are using the [xgen](https://github.com/xuri/xgen) library to generate Go types from the XSD file so you will need to first install that library with:

`go install github.com/xuri/xgen/cmd/xgen@latest`

Then to generate (or regenerate) Go types from an xsd file run the following command (do not include file extensions in the file names):

`make generate-xsd-types xsd=<XsdFileName> output=<output_file_name>`

This command will read the xsd file saved as `<XsdFileName>` in the xsds directory and output Go types in a file named `<output_file_name>` in the schemas directory.
