# go-lang config loader
Parse generic JSON into go-lang data-structures

Basic usage, for a JSON file `/path/to/file.json` with the contents:
```
{
  "contents": "users",
  "users": [
    {
      "name": "Phil Brookes",
      "github": "philbrookes"
    }
  ]
}
```
Parse this like so:
```
var configLoader = NewConfigLoader()
config, err := configLoader.GetConfigFor("/path/to/file.json")
```

Then pull out values like so:
```
username := config.GetNested("users").GetNested("0").GetValue("name")
```
