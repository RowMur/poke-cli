# Poke CLI

A command line interface to explore the world of Pokemon. The purpose of this project is for me to learn the basics of Go.

### Usage

You'll need Go installed on your machine. Run `go install` in the root of the repository. Then, on the terminal run `poke-cli` to enter the CLI. Once you are in the CLI, the following commands are available...

| Command | Parameters | Output                                                   |
| ------- | ---------- | -------------------------------------------------------- |
| help    | n/a        | prints a list of available commands                      |
| exit    | n/a        | exits the CLI                                            |
| map     | n/a        | displays the next list of locations                      |
| mapBack | n/a        | displays the previous list of locations                  |
| explore | location   | prints a list of available pokemon in the given location |
| catch   | pokemon    | attempts to catch the given pokemon                      |
| inspect | pokemon    | prints details of the given pokemon                      |
| pokedex | n/a        | prints the list of caught pokemon                        |
