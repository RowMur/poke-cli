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

### Ideas for further improvement

- [ ] Update the CLI to support the "up" arrow to cycle through previous commands
- [ ] Simulate battles between pokemon
- [ ] Add more unit tests
- [ ] Refactor your code to organize it better and make it more testable
- [ ] Keep pokemon in a "party" and allow them to level up
- [ ] Allow for pokemon that are caught to evolve after a set amount of time
- [ ] Persist a user's Pokedex to disk so they can save progress between sessions
- [ ] Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- [ ] Random encounters with wild pokemon
- [ ] Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon
