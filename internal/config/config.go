package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pedrocmart/maze-go/consts"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Host      string
	Port      int
	DBURL     string
	DBTimeout int
	DBRefresh int
	LogOutput string
	LogFormat string
	Run       string
	Migrate   consts.Migration
}

func NewConfig(args ...string) (cfg *Config, err error) {
	migration := setup()

	initFlags()
	if len(args) == 0 {
		err = loadFlags(os.Args[1:])
	} else {
		err = loadFlags(args)
	}
	if err != nil {
		return
	}

	cfg = &Config{Migrate: migration}
	if err = loadConfig(cfg, migration); err != nil {
		return nil, err
	}

	logrus.Infof("Configuration loaded with params - %+v", cfg)
	return
}

func setup() (migration consts.Migration) {
	viper.SetEnvPrefix(consts.AppName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	var cmdMigrateUp = &cobra.Command{
		Use:   "migrate up/down",
		Short: "Run migrations",
		Long:  "Run migrations",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			migration = consts.Migration(args[0])
		},
	}

	var rootCmd = &cobra.Command{Use: consts.AppName}
	rootCmd.AddCommand(cmdMigrateUp)
	_ = rootCmd.Execute()

	return
}

func initFlags() {
	pflag.String(consts.DatabaseURL, "", "database url")

	pflag.Int(consts.DatabaseRefresh, 5, "database refresh interval in seconds")
	pflag.Int(consts.DatabaseTimeout, 120, "database timeout in seconds")

	pflag.String(consts.HttpHost, "0.0.0.0", "HTTP host bind interface")
	pflag.Int(consts.HttpPort, 5000, "HTTP port to listen on")
	pflag.String(consts.LogOutput, "STDOUT", "log output")
	pflag.String(consts.LogFormat, "json", "log format")
}

func loadFlags(args []string) error {
	if err := pflag.CommandLine.Parse(args); err != nil {
		return err
	}
	return viper.BindPFlags(pflag.CommandLine)
}

func loadConfig(cfg *Config, migration consts.Migration) error {
	err := addConfigParam(&cfg.LogOutput, consts.LogOutput, false, "STDOUT")
	if err != nil {
		return err
	}

	err = addConfigParam(&cfg.LogFormat, consts.LogFormat, false, "json")
	if err != nil {
		return err
	}

	err = addConfigParam(&cfg.DBURL, consts.DatabaseURL, true, "")
	if err != nil {
		return err
	}

	err = addConfigParam(&cfg.DBTimeout, consts.DatabaseTimeout, false, 120)
	if err != nil {
		return err
	}

	err = addConfigParam(&cfg.DBRefresh, consts.DatabaseRefresh, false, 5)
	if err != nil {
		return err
	}

	if migration != consts.MigrateUp && migration != consts.MigrateDown && migration != consts.MigratePrint {
		err = addConfigParam(&cfg.Host, consts.HttpHost, true, "")
		if err != nil {
			return err
		}

		err = addConfigParam(&cfg.Port, consts.HttpPort, true, 0)
		if err != nil {
			return err
		}
	}

	return nil
}

func addConfigParam(dest interface{}, variable string, required bool, defaultValue interface{}) error {
	if dest == nil {
		return fmt.Errorf("nil destination provided for variable %s", variable)
	}

	switch d := dest.(type) {
	case *string:
		_, isStr := defaultValue.(string)
		if !isStr {
			return fmt.Errorf("default value %v for string config field is not a string", defaultValue)
		}

		if !viper.IsSet(variable) {
			if required {
				return fmt.Errorf("%v %s", variable, consts.RequiredParam)
			}
			*d = defaultValue.(string)
		} else {
			*d = viper.GetString(variable)
		}

	case *int:
		_, isInt := defaultValue.(int)
		if !isInt {
			return fmt.Errorf("default value %v for int config field is not an int", defaultValue)
		}

		if !viper.IsSet(variable) {
			if required {
				return fmt.Errorf("%v %s", variable, consts.RequiredParam)
			}
			*d = defaultValue.(int)
		} else {
			*d = viper.GetInt(variable)
		}

	case *bool:
		_, isBool := defaultValue.(bool)
		if !isBool {
			return fmt.Errorf("default value %v for bool config field is not a bool", defaultValue)
		}

		if !viper.IsSet(variable) {
			if required {
				return fmt.Errorf("%v %s", variable, consts.RequiredParam)
			}
			*d = defaultValue.(bool)
		} else {
			*d = viper.GetBool(variable)
		}

	default:
		return errors.New("invalid destination type provided")
	}

	return nil
}
