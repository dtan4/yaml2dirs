# yaml2dirs

```sh-session
$ cat dirs.yaml
- japan
  - nagoya
  - osaka
  - tokyo
    - shibuya
    - shinjuku
- malaysia
  - kuala_lumpur
- singapore

$ yaml2dir dirs.yaml

$ tree --charset=ascii .
.
|-- japan
|   |-- nagoya
|   |-- osaka
|   `-- tokyo
|       |-- shibuya
|       `-- shinjuku
|-- malaysia
|   `-- kuala_lumpur
`-- singapore

9 directories, 0 files
```

## Author

Daisuke Fujita ([@dtan4](https://github.com/dtan4))
