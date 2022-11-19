package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"projectname/internal/project/domain/configuration"
	"strings"
)

const ServiceName = `DefaultAppConfiguration`

func New() *viper.Viper {
	return viper.NewWithOptions(viper.KeyDelimiter(configuration.LayerSeparator))
}

// SetDefaults устанавливает дефолтные значения конфигурации приложения
func SetDefaults(cfg *viper.Viper) error {
	cfg.SetDefault(configuration.LogLevel, configuration.DefaultLogLevel)
	cfg.SetDefault(configuration.HttpHost, configuration.DefaultHttpHost)
	cfg.SetDefault(configuration.HttpPort, configuration.DefaultHttpPort)
	cfg.SetDefault(configuration.FileLogName, configuration.DefaultFileLogName)
	cfg.SetDefault(configuration.LogMaxAge, configuration.DefaultLogMaxAge)
	cfg.SetDefault(configuration.LogMaxSize, configuration.DefaultLogMaxSize)
	cfg.SetDefault(configuration.LogMaxBackups, configuration.DefaultLogMaxBackups)
	cfg.SetDefault(configuration.DevelopmentMode, configuration.DefaultDevelopmentMode)
	cfg.SetDefault(configuration.DatabaseConn, configuration.DefaultDatabaseConn)
	cfg.SetDefault(configuration.OpenWeatherKey, configuration.DefaultOpenWeatherKey)
	cfg.SetDefault(configuration.OpenWeatherUrl, configuration.DefaultOpenWeatherUrl)
	cfg.SetDefault(configuration.SyncGetCityWeatherTime, configuration.DefaultGetCityWeatherTime)
	cfg.SetDefault(configuration.SyncСntDayArchiveTime, configuration.DefaultСntDayArchiveTime)
	cfg.SetDefault(configuration.СntDayArchive, configuration.DefaultCntDayArchive)

	dir, err := determineBaseDir()

	if err != nil {
		return err
	}

	dirBin := filepath.Join(dir, configuration.DefaultDirBin)
	dirVar := filepath.Join(dir, configuration.DefaultDirVar)
	dirLog := filepath.Join(dirVar, configuration.DefaultDirLog)

	cfg.SetDefault(configuration.DirApp, dir)
	cfg.SetDefault(configuration.DirBin, dirBin)
	cfg.SetDefault(configuration.DirVar, dirVar)
	cfg.SetDefault(configuration.DirLog, dirLog)

	return nil
}

// ReadEnv читает конфигурации из ENV
func ReadEnv(cfg *viper.Viper) error {
	if db := os.Getenv(configuration.FullEnvPrefix + configuration.DatabaseConn); db != "" {
		_ = os.Setenv(strings.ToUpper(configuration.FullEnvPrefix+configuration.DatabaseConn), db)
	}

	cfg.SetEnvPrefix(configuration.EnvPrefix)
	cfg.AutomaticEnv()

	return nil
}
