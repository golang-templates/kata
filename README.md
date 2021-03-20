# Go Kata Template

A repository template for exercises using Visual Studio Code.

You can also simply clone this repository if you do not want to put your code in GitHub.

## Features

Visual Studio Code configuration:

- Formatting code, running build, running tests with code coverage and linting on file save.
- Default build task running unit tests with race detector: `F1` -> `Tasks: Run Build Task (CTRL+ALT+B)`.

[Development Containers](https://code.visualstudio.com/docs/remote/containers) support.

Continous integration via [GitHub Actions](https://github.com/features/actions).

Dependencies scanning and updating thanks to [Dependabot](https://dependabot.com).

## Setup

### Local Machine

Follow these steps if you are OK installing and using Go on your machine.

1. Install [Go](https://golang.org/doc/install).
1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Install [Go extension](https://code.visualstudio.com/docs/languages/go).
1. Clone and open this repository.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.
1. `F1` -> `Tasks: Run Build Task (CTRL+ALT+B)`.

### Development Container

Follow these steps if you do not want to install Go on your machine and you prefer to use a Development Container instead.

1. Install [Visual Studio Code](https://code.visualstudio.com/).
1. Follow [Developing inside a Container - Getting Started](https://code.visualstudio.com/docs/remote/containers#_getting-started).
1. Clone and open this repository.
1. `F1` -> `Remote-Containers: Reopen in Container`.
1. `F1` -> `Go: Install/Update Tools` -> (select all) -> OK.
1. `F1` -> `Tasks: Run Build Task (CTRL+ALT+B)`.

## Kata catalogues

- <https://codingdojo.org/kata/>
- <http://codekata.com/>
- <https://github.com/ardalis/kata-catalog>

## Contributing

Simply create an issue or a pull request.
