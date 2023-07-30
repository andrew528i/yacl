# 🌀 YACL - Yet Another Config Loader

YACL is a package designed to simplify configuration management in Go applications. It allows you to load and merge configuration from various sources such as environment variables, command-line flags, and multiple file formats into an arbitrary config structure instance. Also, it is designed to work out of the box with just a single line of code to load your configuration.

## 🌟 Features

- Multiple sources to load config hierarchically:
  - 🚩 **Command-line flags**
  - 🌐 **Environment variables**
  - 📜 **YAML files**
  - 📄 **JSON files**
  - 💾 **Binary files**
  - 🆕 **Default config** instance (optional)
- Automatically **merges configs' fields** with non-default values from different sources
- Supports **nested structs**
- Supports **slices**
- No need to bind variables to config struct by hand
- Easy-to-use single line **convenient API**
- **Aliases** with the tag `yacl: "newFieldName"`


## 📝 Supported types

- string
- bool
- uint32, uin64
- int32, int64
- float64
- []string
- []bool
- []uint32, []uint64
- []int32, []int64
- []float64

---

## 💻 Installation

```go
go get github.com/andrew528i/yacl
```

## 🧑‍💻 Usage

Here's an example of how to use YACL to load configuration from multiple sources with just a single line of code.

```go
import (
	"fmt"

	"github.com/andrew528i/yacl"
)

type DatabaseConfig struct {
	Hostname string
	Port     uint
}

type Config struct {
	Database     DatabaseConfig
	CanRestart   bool
	Tags         []string
	Temperatures []float64
}

func DefaultConfig() *Config {
	// Optional fallback struct with default values
	return &Config{
		Database:     DatabaseConfig{Hostname: "localhost-default"},
		CanRestart:   true,
		Tags:         []string{"hello", "world"},
		Temperatures: nil,
	}
}

func main() {
	cfg, err := yacl.Parse[Config](DefaultConfig()) // DefaultConfig() is optional
	// cfg, err := yacl.Parse[Config]()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
}
```

### YAML file

config.yaml content:
```yaml
database:
  hostname: localhost-yaml
  port: 5432
can_restart: true
tags:
 - hello
 - world
temperatures:
 - 12.75
 - 17.49929
 - 27.46
```

```bash
$ ./app

&{Database:{Hostname:localhost-yaml Port:5432} CanRestart:true Tags:[hello world] Temperatures:[12.75 17.49929 27.46]}
```

### Environment variables

```bash
$ export DATABASE_HOSTNAME=localhost-env
$ export DATABASE_PORT=1234
$ export CAN_RESTART=true
$ export TAGS=one,two
$ export TEMPERATURES=123.321,234.432,456.654
$ ./app

&{Database:{Hostname:localhost-env Port:1234} CanRestart:true Tags:[one two] Temperatures:[123.321 234.432 456.654]}
```

### Command-line flags

```bash
$ ./app -database-hostname localhost-flags -database-port 6543 -can-restart -tags aaa -tags bbb -tags ccc -temperatures 1.223 -temperatures 2.332

&{Database:{Hostname:localhost-flags Port:6543} CanRestart:true Tags:[aaa bbb ccc] Temperatures:[1.223 2.332]}
```

---

## Documentation

### Global params

---

#### SetEnvPrefix

Use this to set a global prefix for environment variables:

```go
yacl.SetEnvPrefix("APP")
```

For example, struct field `HTTPPort uint` will become `APP_HTTP_PORT`.

---

#### SetEnvDelimiter

```go
yacl.SetEnvDelimiter("__")
```

For example, struct field `HTTPPort uint` will become `HTTP__PORT`.

---

#### SetFilePath

Specifies the path where to look for config files for all the formats: yaml, json, and bin. Also, YACL always looks for config files in the current working directory.

---

#### SetFilename

Sets the default filename for the config filename. For example:

```go
yacl.SetFilename("my-config")
```

So YACL will look for those files: my-config.yaml, my-config.json, and my-config.bin.

---

#### SetFlagDelimiter

```go
yacl.SetFlagDelimiter("_")
```

So, struct field `DatabasePort uint` will become command-line flag `-database_port`.

---

#### Parse

Parses all the config source hierarchically. See the usage section for more details and examples.

---

### 🌟 Instance parameters

In order to create a separate instance of YACL with its own params you can:

```go
y := yacl.New()
```

---

#### SetEnvPrefix

Same as the global version.

---

#### SetEnvDelimiter

Same as the global version.

---

#### AddFilePath

Adds a path to look for a config.

---

#### SetFilename

Same as the global version.

---

#### SetFlagDelimiter

Same as the global version.

---

#### SetIgnoreFlags

In some cases, you may need not to touch flag params. For such a case, you can `yacl.SetIgnoreFlags(true)` and YACL will not clear default flags.

---

#### Parse

Same as the global version.
