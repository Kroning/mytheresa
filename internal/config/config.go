package config

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	configDir         = "./"
	configFilePathEnv = "CONFIG_FILE_PATH"
	appNameKey        = "APP_NAME"
	envNameKey        = "ENVIRONMENT_NAME"
)

var (
	logger *zap.Logger
)

type Options struct {
	Dir                string // dir where configs located, set to "./config/" if empty
	Type               string // yaml or json. Set to yaml if empty
	DevFile            string // dev config file name, set to dev.(yaml|json) if empty
	ProdFile           string // prod config file name, set to main.(yaml|json) if empty
	ReplaceFromEnvVars bool   // replace config values from ENV VARS, false by default
	EnvVarsPrefix      string // prefix for ENV VARS, empty by default
	Validator          *validator.Validate
}

func NewConfig(log *zap.Logger) (*Config, error) {
	cfg := &Config{}
	err := Parse(cfg, Options{Dir: "config", Type: "yaml"}, log)
	if err != nil {
		return nil, err
	}
	enrichAppEnvName(cfg)

	return cfg, nil
}

func Parse(configStruct interface{}, opts Options, log *zap.Logger) error {
	if log == nil {
		logger = zap.NewNop()
	} else {
		logger = log
	}

	t := reflect.TypeOf(configStruct)
	if t.Kind() != reflect.Ptr {
		return errors.New("configStruct arg must be pointer")
	}
	if t.Elem().Kind() != reflect.Struct {
		return errors.New("configStruct arg must be pointer to struct")
	}

	err := opts.fill()
	if err != nil {
		return err
	}
	loadedFromFile, err := loadFromFile(opts)
	if err != nil {
		return err
	}

	if !loadedFromFile {
		return errors.New("cannot load from config")
	}

	if opts.ReplaceFromEnvVars {
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AllowEmptyEnv(true)
		viper.SetEnvPrefix(opts.EnvVarsPrefix)
		viper.AutomaticEnv()
	}

	err = viper.Unmarshal(configStruct)
	if err != nil {
		return err
	}

	// some checks
	err = opts.Validator.Struct(configStruct)
	if err != nil {
		return err
	}
	return nil

}

func (o *Options) fill() error {
	if o.Type == "" {
		o.Type = "yaml"
	}
	if o.DevFile == "" {
		o.DevFile = "dev." + o.Type
	}
	if o.ProdFile == "" {
		o.ProdFile = "main." + o.Type
	}
	if o.Dir == "" {
		o.Dir = configDir
	}
	if os.Getenv(configFilePathEnv) != "" {
		o.Dir = path.Dir(os.Getenv(configFilePathEnv))
		o.ProdFile = path.Base(os.Getenv(configFilePathEnv))
		o.Type = strings.ToLower(path.Ext(os.Getenv(configFilePathEnv)))
		o.Type = strings.ReplaceAll(o.Type, ".", "")
	}

	if o.Type != "json" && o.Type != "yaml" {
		return errors.Errorf("bad format %s, must be json or yaml", o.Type)
	}
	if !strings.HasSuffix(o.Dir, string(os.PathSeparator)) {
		o.Dir = o.Dir + string(os.PathSeparator)
	}
	if o.Validator == nil {
		o.Validator = validator.New()
	}
	return nil
}

func loadFromFile(opts Options) (loaded bool, err error) {
	configPath := opts.Dir + opts.DevFile
	if err := fileExists(configPath); err != nil {
		logger.Debug("dev file not exists or empty", zap.String("path", configPath), zap.Error(err))
		configPath = opts.Dir + opts.ProdFile
	}

	if err := fileExists(configPath); err != nil {
		logger.Debug("prod file not exists or empty", zap.String("path", configPath), zap.Error(err))
		return false, nil
	}

	logger.Info("load config from file", zap.String("path", configPath))
	configPath, err = filepath.Abs(configPath)
	if err != nil {
		return false, err
	}
	viper.SetConfigFile(configPath)
	viper.SetConfigType(opts.Type)
	//fmt.Printf("start parsing config file %s, env prefix %s\n", configPath, prefix)
	err = viper.ReadInConfig()
	if err != nil {
		return false, err
	}
	logger.Info("done loading config from file")
	return true, nil
}

func fileExists(path string) (err error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return
	}
	info, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("file not exists: " + absPath)
		}
		return
	}
	if info.IsDir() {
		return errors.New("must be file: " + absPath)
	}
	if info.Size() == 0 {
		return errors.New("file is empty: " + absPath)
	}
	return
}

func enrichAppEnvName(cfg *Config) {
	if appName, ok := os.LookupEnv(appNameKey); ok && appName != "" {
		cfg.App.Name = appName
	}

	if envName, ok := os.LookupEnv(envNameKey); ok && envName != "" {
		cfg.App.Env = envName
	}
}
