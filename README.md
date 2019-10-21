# Enpass to 1Password converter

The Enpass2onepassword it's a tool to convert records from [Enpass](https://www.enpass.io/)   to [1Password](https://1password.com). The need for this is due to the fact that the 1Password doesn't support import files from Enpass.

## Export from Enpass

Enpass -> File -> Export to json file -> Save file enpass.json (default file of enpass export for tool)

## Run

First of all, compile the source code.
```shell script
$ go build .
```

To show the arguments of tool use the help
```bash
$ ./enpass2onepassword --help
```

By default at the same directory should place the file with information from Enpass (exported values) `enpass.json`. In this case, you can to run the tool without any arguments:
```shell script
$ ./enpass2onepassword
```

In the case, if you gave the name of exported file different from `enpass.json` need to add the argument `-enpass_src_path`:
```shell script
$ ./enpass2onepassword --enpass_src_path=enpass_export.json
```

## Import to 1Password

The result file (output of the tool) might used to import in the 1Password:

1Password -> File -> Import -> Other -> Import 