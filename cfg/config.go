package cfg

import (
	"fmt"
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	flag "github.com/spf13/pflag"
)

type Config struct {
	*koanf.Koanf
}

func NewConfig() *Config {
	k := koanf.New(".")
	// Use the POSIX compliant pflag lib instead of Go's flag lib.
	f := flag.NewFlagSet("config", flag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	// Path to one or more config files to load into koanf along with some config params.
	f.StringSlice("conf", []string{"config.yaml"}, "path to one or more .yaml config files")
	f.String("dbdriver", "postgres", "database driver")
	f.String("dbhost", "localhost:5432", "database host name")
	f.String("dbuser", "postgres", "database user name")
	f.String("dbpasswd", "postgres", "database password")
	f.String("dbname", "limits", "database name")

	f.String("listen", ":10002", "service listen port")

	f.Parse(os.Args[1:])

	// Load the config files provided in the commandline.
	cFiles, _ := f.GetStringSlice("conf")
	for _, c := range cFiles {
		if err := k.Load(file.Provider(c), yaml.Parser()); err != nil {
			log.Fatalf("error loading file: %v", err)
		}
	}

	// The values may be set in the config. But we can still override at runtime
	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	return &Config{k}
}

func (c *Config) DatabaseConnectionString() string {
	cs := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable", c.String("dbdriver"),
		c.String("dbuser"),
		c.String("dbpasswd"),
		c.String("dbhost"),
		c.String("dbname"))
	return cs
}
