package configs

var Conf = new(Config)

type Config struct {
	App           ApplicationConfig   `mapstructure:"app"`
	Log           LoggerConfig        `mapstructure:"log"`
	DB            DBConfig            `mapstructure:"database"`
	Redis         RedisConfig         `mapstructure:"redis"`
	Email         EmailConfig         `mapstructure:"email"`
	JWT           JWTConfig           `mapstructure:"jwt"`
	Oss           OssConfig           `mapstructure:"oss"`
	Elasticsearch ElasticsearchConfig `mapstructure:"elasticsearch"`
	MQ            MQConfig            `mapstructure:"mq"`
	Coze          CozeConfig          `mapstructure:"coze"`
}

type ApplicationConfig struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Env         string `mapstructure:"env"`
	LogfilePath string `mapstructure:"logfilePath"`
}
type LoggerConfig struct {
	Level    int8   `mapstructure:"level"`
	Format   string `mapstructure:"format"`
	Director string `mapstructure:"director"`
	ShowLine bool   `mapstructure:"show-line"`
}

type DBConfig struct {
	Driver []string `mapstructure:"driver"`
	MySQL  struct {
		AutoMigrate bool   `mapstructure:"migrate"`
		Dsn         string `mapstructure:"dsn"`
	} `mapstructure:"mysql"`
}

type MQConfig struct {
	Enabled  []string `mapstructure:"enabled"`
	RabbitMQ struct {
		Dsn             string `mapstructure:"dsn"`
		ChannelPoolSize string `mapstructure:"channelPoolSize"`
	} `mapstructure:"rabbitmq"`
	Kafka struct {
		Brokers []string `mapstructure:"brokers"`
		host    string   `mapstructure:"host"`
		port    int      `mapstructure:"port"`
	} `mapstructure:"kafka"`
}

type RedisConfig struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type EmailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type OssConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"accessKeyID"`
	AccessKeySecret string `mapstructure:"accessKeySecret"`
	BucketName      string `mapstructure:"bucketName"`
}

type ElasticsearchConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	UserName string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type CozeConfig struct {
	Token string `mapstructure:"token"`
	//BotID string `mapstructure:"botID"`
}
