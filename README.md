# swiftCodeProject

## Description
SwiftCodeProject is an application I developed in Go to parse, store, and provide access to SWIFT code data through a RESTful API, ensuring efficient querying and retrieval of banking information. It uses PostgreSQL as the database for optimized storage and fast queries.

---

## Project Scope

### SWIFT Code Overview
A SWIFT code (Bank Identifier Code - BIC) uniquely identifies a bank’s branch or headquarters for international wire transfers. Initially stored in an Excel file, this data is now structured and exposed via an API.

---

## Features

### Data Parsing
- Extracts SWIFT data from the provided Excel file.
- Identifies headquarters (codes ending in `XXX`) and links branches.
- Stores country codes and names in uppercase.

### Database
- Uses PostgreSQL for efficient storage and querying.
- Enables retrieval of individual SWIFT codes and country-specific data.

### REST API Endpoints
- **GET** `/v1/swift-codes/{swift-code}` – Retrieve details of a SWIFT code.
- **GET** `/v1/swift-codes/country/{countryISO2code}` – List all SWIFT codes for a country.
- **POST** `/v1/swift-codes` – Add a new SWIFT code entry.
- **DELETE** `/v1/swift-codes/{swift-code}` – Remove a SWIFT code.

