# newuuids

Pronounced _[njuɪdz](http://ipa-reader.xyz/?text=nju%C9%AAdz&voice=Emma)_. Batch replace UUIDs in text.

## CLI

### Precompiled releases
You can download a (x64) pre-compiled release from [the releases page on GitHub](https://github.com/motevets/newuuids/releases).

## Build it yourself
1. Install golang (testsed with 1.16)
1. `git clone https://github.com/motevets/newuuids`
1. `cd newuuids`
1. `go build cmd/newuuids`

### Usage

Reads line-by-line from STDIN, replaces UUIDs with new random UUIDs, and prints result to STDOUT.
If it encounters a UUID that its replaced before, it will replace that UUID with the same UUID is used before.const

Example:

    cat /path/to/file/with/uuids | newuuids

## Go Library

Located in [pkg/uuidbump](./pkg/uuidbump).

## License

MIT. See [LICENSE](./LICENSE)
