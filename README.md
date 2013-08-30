Increments version numbers in version.json *and* generates a funny-ish codename for each version.

Version data is kept in *version.json*, in the current directory where increment_version is being run.
The version.json file is generated automatically, if it does not exist in the current directory.

Sample version.json file looks like this:

```json
{
	"full": "0.0.1",
	"major": 0,
	"minor": 0,
	"dot": 1,
	"codename": "Muddy Sponge"
}
```

Running increment_version from command line outputs the next version:

```
increment_version
```
0.0.2 'Trimmer Polishes'

version.json is overwritten and the resulting sample version.json looks like this:

```json
{
	"full": "0.0.2",
	"major": 0,
	"minor": 0,
	"dot": 1,
	"codename": "Trimmer Polishes"
}
```

Install
=======
```
go install github.com/tanel/increment_version
```

Use
===
```
increment_version
```
