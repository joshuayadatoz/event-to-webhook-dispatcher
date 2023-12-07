# Event-to-Webhook Dispatcher
This small Go application simulates the processing of mock insurance policy events and dispatches notifications to a specified webhook URL. 

## Getting Started
These instructions will get the project up and running on your local machine for development and testing purposes.

## Prerequisites
Go (version 1.21.5)

Internet access (for sending data to the webhook)

## Installing
### Clone the Repository:
Clone this repository to your local machine, or download the source code.

### Navigate to the Project Directory:
Open a terminal and navigate to the directory where you cloned the repository.

### Running the Application
To run the application, execute the following command in the terminal within the project directory:

~~~ bash
go run .
~~~ 
This will start the server and begin sending mock event data to the configured webhook URL every 2 seconds, for a total of 3 events. Each event tagged with the current time, featuring a random count of bikes, and a corresponding calculated total value.

### Testing the Application
The project includes a suite of automated tests. To run these tests, execute the following command in the project directory:

~~~ bash
go test ./... -v
~~~ 
This command runs all tests in the project and outputs the results, indicating whether each test passed or failed.

### Viewing Webhook Data
To view the data received by the webhook, visit [this link](https://webhook.site/#!/2898cf90-f0a5-4718-b5cd-9ab40fa2f498/7a51d7db-7d67-4cf0-95ec-4be3d7c3c898/1). Here, you will see the incoming requests from the application.