# verG

verG is a semantic versioning CLI.

## Installation

Use homebrews to install verG.

```bash
brew tap trendyol/trendyol-tap
brew install verg
```

Use go module to install verG.

```bash
go get github.com/trendyol/verg
```

## Increment Version

```bash
verg 1.0.0 --major

--> 2.0.0
```

### Flags    

#### --major (-x)
This flag increment to major version.

```bash
verg 1.0.0 --major

--> 2.0.0
```

#### --minor (-y)
This flag increment to minor version.

```bash
verg 1.0.0 --minor

--> 1.1.0
```

#### --patch (-z)
This flag increment to patch version.

```bash
verg 1.0.0 --patch

--> 1.0.1
```

#### --release (-r)
This flag increment to release version.

```bash
verg 1.0.0 -r

--> 1.0.0-RELEASE.0
```

#### --beta (-b)
This flag increment to beta version.

```bash
verg 1.0.0 -b

--> 1.0.0-BETA.0
```

#### --alpha (-a)
This flag increment to alpha version.

```bash
verg 1.0.0 -a

--> 1.0.0-ALPHA.0
```


## Compare Version 
This command compare to versions. Operators: [==] [>=] [<=] [>] [<]

```bash
verge compare "1.0.0 == 1.0.0"

--> true
```
```bash
verge compare "1.0.0 > 1.0.1"

--> false
```
