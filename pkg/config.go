package config

import (
	"github.com/spf13/viper"
)

// Config adalah struct yang akan digunakan untuk menyimpan konfigurasi aplikasi.
type Config struct {
	// Tambahkan field konfigurasi sesuai dengan kebutuhan Anda.
	// Misalnya:
	AppHost      string
	AppPort      string
	DatabaseURL  string
	DatabasePort string
}

// InitConfig inisialisasi konfigurasi menggunakan Viper.
func InitConfig() (*Config, error) {
	// Inisialisasi Viper
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// Baca konfigurasi dari file (opsional)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Buat instance Config dan isi sesuai dengan konfigurasi yang dibaca
	config := &Config{
		AppHost:      viper.GetString("AppHost"),
		AppPort:      viper.GetString("AppPort"),
		DatabasePort: viper.GetString("DatabasePort"),
		DatabaseURL:  viper.GetString("DatabaseURL"),
	}

	return config, nil
}
