# bstset
## Description
bstset set Bluestacks options by overwriting configuration file.

## Usage 
### Specify from command line
```sh
bstset set target... [flags]
```
`-f, --file string`   BlueStacks configuration file path (default "C:\\ProgramData\\BlueStacks_nxt\\bluestacks.conf")

Example:
```sh
bstset set 'bst.feature.macros:1'
```

Note:
`target` must be followed `name:value` format.

### Specify from JSON file
```sh
bstset load file
```

Example:
```sh
bstset load ./config.json
```

File format:
<pre>
bstConfigPath: string
  Path to BlueStacks configuration file

targets: {name: string, value: string}[]
  Array of option you want to set.
</pre>

### For more information
Run below command to get for more information.
```sh
bstset -h
```
