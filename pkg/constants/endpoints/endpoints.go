package endpoints

const (
	KafkaEndpoint           = "localhost:9092"
	RabbitMQEndpoint        = "localhost:5672"
	PulsarEndpoint          = "localhost:6650"
	ActiveMQEndpoint        = "localhost:61616"
	NATSEndpoint            = "localhost:4222"
	RocketMQEndpoint        = "localhost:9876"
	VerneMQEndpoint         = "localhost:1883"
	MosquittoEndpoint       = "localhost:1883"
	ElasticsearchEndpoint   = "localhost:9200"
	SolrEndpoint            = "localhost:8983"
	LuceneEndpoint          = "localhost:7474"
	MySQLEndpoint           = "localhost:3306"
	PostgreSQLEndpoint      = "localhost:5432"
	MongoDBEndpoint         = "localhost:27017"
	CassandraEndpoint       = "localhost:9042"
	RedisEndpoint           = "localhost:6379"
	OracleEndpoint          = "localhost:1521"
	SQLServerEndpoint       = "localhost:1433"
	DB2Endpoint             = "localhost:50000"
	InformixEndpoint        = "localhost:9088"
	SQLiteEndpoint          = "localhost:/path/to/sqlite.db"
	IBMWebSphereEndpoint    = "localhost:1414"
	IBMMessageSightEndpoint = "localhost:1883"
	TIBCOEMSQueueEndpoint   = "localhost:7222"
	ApacheMQEndpoint        = "localhost:61616"
	ApacheKafkaEndpoint     = "localhost:9092"
	ApachePulsarEndpoint    = "localhost:6650"
	NATSStreamingEndpoint   = "localhost:4222"
	PravegaEndpoint         = "localhost:9090"
	ApacheFlinkEndpoint     = "localhost:8081"
	EventStoreEndpoint      = "localhost:2113"
	AxonServerEndpoint      = "localhost:8124"
	HadoopEndpoint          = "localhost:9000"
	SparkEndpoint           = "localhost:7077"
	FlinkEndpoint           = "localhost:8081"
	HiveEndpoint            = "localhost:10000"
	PrestoEndpoint          = "localhost:8080"
	TensorFlowEndpoint      = "localhost:8501"
	PyTorchEndpoint         = "localhost:8000"
	ScikitLearnEndpoint     = "localhost:5000"
	AzureCosmosDBEndpoint   = "https://your-cosmosdb-endpoint.documents.azure.com:443/"
	GCPFirestoreEndpoint    = "https://firestore.googleapis.com/v1/projects/your-project-id/databases/(default)/documents"
	AWSDynamoDBEndpoint     = "https://dynamodb.your-region.amazonaws.com"
)

var (
	KnownBrokers          = []string{KafkaEndpoint, RabbitMQEndpoint, PulsarEndpoint, ActiveMQEndpoint, NATSEndpoint, RocketMQEndpoint, VerneMQEndpoint, MosquittoEndpoint, IBMWebSphereEndpoint, IBMMessageSightEndpoint, TIBCOEMSQueueEndpoint, ApacheMQEndpoint, ApacheKafkaEndpoint, ApachePulsarEndpoint, NATSStreamingEndpoint}
	KnownSearchEngines    = []string{ElasticsearchEndpoint, SolrEndpoint, LuceneEndpoint}
	KnownDatabases        = []string{MySQLEndpoint, PostgreSQLEndpoint, MongoDBEndpoint, CassandraEndpoint, RedisEndpoint, OracleEndpoint, SQLServerEndpoint, DB2Endpoint, InformixEndpoint, SQLiteEndpoint}
	KnownStreamingEngines = []string{ApacheKafkaEndpoint, ApachePulsarEndpoint, NATSStreamingEndpoint, PravegaEndpoint, ApacheFlinkEndpoint}
	KnownEventStores      = []string{EventStoreEndpoint, AxonServerEndpoint}
	KnownBigDataEngines   = []string{HadoopEndpoint, SparkEndpoint, FlinkEndpoint, HiveEndpoint, PrestoEndpoint}
	KnownMLEngines        = []string{TensorFlowEndpoint, PyTorchEndpoint, ScikitLearnEndpoint}
	KnownCloudDatastores  = []string{AzureCosmosDBEndpoint, GCPFirestoreEndpoint, AWSDynamoDBEndpoint}

	BrokerMap          = map[string]string{}
	SearchEngineMap    = map[string]string{}
	DatabaseMap        = map[string]string{}
	StreamingEngineMap = map[string]string{}
	EventStoreMap      = map[string]string{}
	BigDataEngineMap   = map[string]string{}
	MLEngineMap        = map[string]string{}
	CloudDatastoreMap  = map[string]string{}
)

func init() {
	for _, broker := range KnownBrokers {
		BrokerMap[broker] = broker
	}

	for _, searchEngine := range KnownSearchEngines {
		SearchEngineMap[searchEngine] = searchEngine
	}

	for _, database := range KnownDatabases {
		DatabaseMap[database] = database
	}

	for _, streamingEngine := range KnownStreamingEngines {
		StreamingEngineMap[streamingEngine] = streamingEngine
	}

	for _, eventStore := range KnownEventStores {
		EventStoreMap[eventStore] = eventStore
	}

	for _, bigDataEngine := range KnownBigDataEngines {
		BigDataEngineMap[bigDataEngine] = bigDataEngine
	}

	for _, mlEngine := range KnownMLEngines {
		MLEngineMap[mlEngine] = mlEngine
	}

	for _, cloudDatastore := range KnownCloudDatastores {
		CloudDatastoreMap[cloudDatastore] = cloudDatastore
	}
}
