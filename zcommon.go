package gimg

import (
	_ "encoding/json"
	_ "fmt"
)

const (
	PROJECT_VERSION = "1.0"
	MAX_LINE        = 1024
	CACHE_KEY_SIZE  = 128
	RETRY_TIME_WAIT = 1000
	CACHE_MAX_SIZE  = 1048576 //1024*1024
	PATH_MAX_SIZE   = 512
)

type ZRequest struct {
	Md5        string
	ImageType  string
	Width      int
	Height     int
	Proportion int
	Gary       int
	X          int
	Y          int
	Rotate     int
	Format     string
	Save       int
	Quality    int
}

type ZImageInfo struct {
	Size    int    `json:"size"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Quality int    `json:"quality"`
	Format  string `json:"format"`
}

type ZContext struct {
	Config AppConfig
	Logger *ZLogger
	Cache  *ZCache
	Image  *ZImage
	Redis  *ZRedisDB
}

func NewContext(cfgFile string) (*ZContext, error) {
	cfg, err := LoadConfig(cfgFile)
	if err != nil {
		return nil, err
	}

	//log, err := logger.NewFileLogger("zimg", 0, cfg.System.LogName)
	log, err := NewLogger("gimg", 0)
	if err != nil {
		return nil, err
	}

	var cache *ZCache
	if cfg.Cache.Cache == 1 {
		// var client *memcache.Client
		// cacheAddr := fmt.Sprintf("%s:%d", cfg.Cache.MemcacheHost, cfg.Cache.MemcachePort)
		// client = memcache.New(cacheAddr)
		cache = NewCache(cfg.Cache.MemcacheHost, cfg.Cache.MemcachePort)
	} else {
		cache = nil
	}

	img := NewImage()

	redisDB, err := NewRedisDB(cfg.Storage.SsdbHost, cfg.Storage.SsdbPort)
	if err != nil {
		return nil, err
	}

	return &ZContext{Config: cfg,
		Logger: log,
		Cache:  cache,
		Image:  img,
		Redis:  redisDB,
	}, nil
}
