# Internal Packages

This directory contains the internal packages for the alfred-dndbeyond-monster-workflow application.

## Package Structure

### `monster/`
Contains data types and models for D&D monster information:
- `Monster`: Represents a D&D monster with all its attributes
- `ResultSet`: Mirrors the response from the dnddeutsch.de API

### `sources/`
Provides source book filtering for D&D monsters:
- `ExcludedSources`: List of source books to exclude from results (typically non-official or third-party)
- `ContainsAny()`: Helper function to check if any string from one list is present in another

### `icons/`
Provides icon management for D&D monster types:
- Icon definitions for all monster types (aberration, beast, celestial, etc.)
- `ForType()`: Returns the appropriate icon for a given monster type

## Design Principles

These packages follow Go best practices for 2024:
- Clear separation of concerns
- Comprehensive test coverage
- Package-level documentation
- Exported identifiers follow Go naming conventions
- Internal-only packages that cannot be imported by external projects
