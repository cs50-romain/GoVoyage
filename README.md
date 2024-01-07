# GoVoyage

This is a simple web application that allows users to compare plane tickets based on their travel details. The tool fetches flight data from a third-party API and displays relevant information such as airline, departure time, arrival time, and price.

## Table of Contents
- [Features](#features)
- [Tools/Technologies](#tools-technologies)
- [Usage](#usage)
- [Optional Enhancements](#optional-enhancements)
- [Contributing](#contributing)
- [License](#license)

## Features

- User-friendly form for entering travel details (start date, end date, destination).
- Dynamic retrieval of flight data from a flight data provider API (serpapi).
- Display of flight information, including airline, departure time, arrival time, and price.
- Basic error handling for API requests and user input validation.
- Responsive and visually appealing interface.

## Tools/Technologies

- HTML, CSS
- Go, HTMX
- Flight Data API: serpapi which scrapes resutls from Google Flights

## Usage

Visit the online version of the project at [govoyage.onrender.com](https://govoyage.onrender.com/) to use the Plane Ticket Comparison Tool. Simply fill out the form with your travel details and click the "Search Flights" button to view available flights.

## Optional Enhancements

Feel free to enhance the project with the following features:

- Create my own scraper to scrape results from Google Flights
- Details modal or additional information page for each flight.
- Better UI for flight results.
- Better user error message.
- Sorting and filtering options for a better user experience.

## Contributing

Contributions are welcome! If you have ideas for improvements or new features, feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
