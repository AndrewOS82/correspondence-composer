# Correspondence Composer
## Composing Rules!

## About

This is an application which listens for incoming events indicating that correspondence should be generated for a client. It fetches relevant data for the correspondence from Zinnia enterprise APIs and then generates XML that can be consumed by other services which prepare and send the correspondence.

## Setup

### Installing Go on your machine

Follow [these instructions](https://go.dev/doc/install) for installing go on Linux, Mac or Windows, then proceed to the "Run the app" section.

### Setting up configs and credentials

A sample env file has been provided as `.env.sample`. To run the app locally, you'll need to copy this into a `.env` file and update username, password and endpoints as necessary.

To make successful requests to the rules engine you'll need to have valid credentials for the environment that the endpoints are hitting.

## Run Instructions

You can run this application in two ways, by installing Go on your machine or via Docker. If you've installed Go on your machine start up app with this command:

`make run`

### Using Docker to run the application

If you already have Docker set up and running use the following commands to run the app with Docker:

```
make docker-build
make localdev
```

## Running Kafka locally with Docker

If you already have Docker set up you can run the following commands to run an instance of Kafka locally with a built in UI that allows you to publish messages with a click of the button.

Run `docker-compose -f kafka.yml up` to start up the instance.

You'll also need to run `echo '127.0.0.1 kafka0' | sudo tee -a /etc/hosts` the first time you spint it up to fix local connectivity.

Visit `localhost:8080` in the browser to see the Kafka UI and set up topics and publish messages.

The current topic in this MVP is "correspondence.test.one" and publishing a message with the value "anniversary" will run the example composer process.

// TODO: add step-by-step instructions on publishing messages to test functionality of the application.

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

This command will read the xsd file saved as `<XsdFileName>` in the xsds directory and output Go types in a file named `<output_file_name>` in the `models/generated` directory.
