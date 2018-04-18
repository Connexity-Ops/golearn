Lyrics Lookup
======

Query a REST API for lyrics by band & song, unmarshal the returned JSON and print the lyrics.

## Example
```go
fmt.Println(GetLyrics("The Champs", "Tequila"))
```
Returns:
```go
"Tequila\n\nTequila\n\nTequila"
```

## Testing
```
 go test
```

## Notes
The public API is located at https://api.lyrics.ovh/v1/<BandName>/<SongName> and requires no key.