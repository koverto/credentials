# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][], and this project adheres to
[Semantic Versioning][].

## Unreleased

## v0.5.1 - 2020-03-31

### Changed

- Updated to github.com/99designs/gqlgen v0.11.3
- Updated to github.com/golang/protobuf v1.3.5
- Updated to github.com/micro/go-micro v2.4.0
- Updated to golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59

## v0.5.0 - 2020-03-31

### Removed

- `CredentialResponse.success`

## v0.4.0 - 2020-03-30

### Changed

- Updated to github.com/koverto/micro@v2.0.1

## v0.3.0 - 2020-03-05

### Changed

- Updated to github.com/koverto/micro@v1.2.0
- Updated to github.com/koverto/uuid@v1.3.0

## v0.2.3 - 2020-02-17

### Changed

- Updated to github.com/koverto/uuid@v1.2.1

## v0.2.2 - 2020-02-17

### Added

- GraphQL (un)marshaling for `CredentialType`

## v0.2.1 - 2020-02-17

### Changed

- User ID only required for filtering for password verification
- Updated to github.com/koverto/uuid@v1.2.0

## v0.2.0 - 2020-02-17

### Changed

- Renamed to `credentials`

## v0.1.0 - 2020-02-17

### Added

- Protobuf API
- `Authn.Create`
- `Authn.Verify`

[keep a changelog]: https://keepachangelog.com/en/1.0.0/
[semantic versioning]: https://semver.org/spec/v2.0.0.html
