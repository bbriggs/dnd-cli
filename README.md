# D&D CLI

A CLI for interacting with your character

## Installation

`go get -u github.com/bbriggs/dnd-cli`

or if you have `dep` installed:

`git clone https://github.com/bbriggs/dnd-cli.git && cd dnd-cli && dep ensure && go run ./main.go`

## Configuration

The project doesn't actually do anything yet except return some very basic values. In fact, it's not useful at all yet. That said, here is an example configuration to start your character (to be saved in the local directory as `config.yaml`:

```yaml
name: Titanius Anglesmith
hp: 20
ac: 13
speed: 5
str: 8
dex: 14
con: 9
int: 12
wis: 12
cha: 17
items:
  - cigar
  - matches
  - Löbrau Beer
  - Amy's wallet
alignment: LG
```

### Planned work
[ ] Save character to a database (SQLite)
[ ] Update HP and other stats interactively
[ ] Initialize/roll new characters (from scratch and templates)
[ ] Experience tracking
[ ] Item tracking that's actually useful
[ ] Money/finance tracking
[ ] Magic tracking (did you use that spell today?)
[ ] Rest status
[ ] Integration with D&D 5e API for searching items, spells, and more
