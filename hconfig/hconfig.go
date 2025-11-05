package hconfig

var Conf = new(Config)

type Config struct {
	Hin           HinConfig           `mapstructure:"hin"`
	Server        ServerConfig        `mapstructure:"server"`
	Client        ClientConfig        `mapstructure:"client"`
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
type HinConfig struct {
	Version          string `mapstructure:"version"`             // The version of the Hin framework.(当前Hin版本号)
	MaxPacketSize    uint32 `mapstructure:"max_pack_size"`       // The maximum size of the packets that can be sent or received.(读写数据包的最大值)
	MaxConn          int    `mapstructure:"max_connection"`      // The maximum number of connections that the server can handle.(当前服务器主机允许的最大连接个数)
	WorkerPoolSize   uint32 `mapstructure:"worker_pool_siz"`     // The number of worker pools in the business logic.(业务工作Worker池的数量)
	MaxWorkerTaskLen uint32 `mapstructure:"max_worker_task_len"` // The maximum number of tasks that a worker pool can handle.(业务工作Worker对应负责的任务队列最大任务存储数量)
	WorkerMode       string `mapstructure:"worker_mode"`         // The way to assign workers to connections.(为连接分配worker的方式)
	MaxMsgChanLen    uint32 `mapstructure:"max_msg_chan_len"`    // The maximum length of the send buffer message queue.(SendBuffMsg发送消息的缓冲最大长度)
	IOReadBuffSize   uint32 `mapstructure:"io_read_buff_size"`   // The maximum size of the read buffer for each IO operation.(每次IO最大的读取长度)
}
type ServerConfig struct {
	Host    string `mapstructure:"host"`     // The IP address of the current server. (当前服务器主机IP)
	TCPPort int    `mapstructure:"tcp_port"` // The port number on which the server listens for TCP connections.(当前服务器主机监听端口号)
	Name    string `mapstructure:"name"`     // The name of the current server.(当前服务器名称)
}
type ClientConfig struct {
	Host    string `mapstructure:"host"`     // The IP address of the current server. (当前客户端连接主机IP)
	TCPPort int    `mapstructure:"tcp_port"` // The port number on which the server listens for TCP connections.(当前客户端主机连接端口号)
	Name    string `mapstructure:"name"`     // The name of the current server.(当前客户端名称)
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
