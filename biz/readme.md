📊 业务场景常用数据结构
1. 缓存类
数据结构	实际应用	典型场景
LRU Cache	Redis LRU、浏览器缓存	最近最少使用淘汰
LFU Cache	Redis LFU	最不经常使用淘汰
哈希表 + 链表	LinkedHashMap	缓存实现
布隆过滤器	缓存穿透防护、黑名单过滤	快速判断"不存在"
2. 存储索引
数据结构	实际应用	典型场景
B+树	MySQL InnoDB索引	范围查询、排序
LSM树	LevelDB、RocksDB	写多读少场景
跳表	Redis ZSet	有序集合、排行榜
前缀树(Trie)	自动补全、敏感词过滤	前缀匹配
3. 限流与流量控制
算法	实际应用	典型场景
令牌桶	网关限流、API限流	突发流量控制
漏桶	网络流量整形	均匀流量输出
滑动窗口	滑动时间窗口统计	实时统计、限流
漏斗算法	登录防刷、下单限流	流量过滤
⏰ 时间轮 (Timing Wheel)
核心应用场景
1. 定时任务调度
// Kafka定时任务实现
type TimingWheel struct {
    tickMs      int64        // 毫秒间隔
    wheelSize   int          // 时间轮大小
    startTime   int64        // 起始时间
    buckets     []*Bucket    // 时间槽
    currentTime int64        // 当前时间
}

type Job struct {
    delay    int64  // 延迟时间
    key      string // 业务key
    callback func() // 回调函数
}

// 应用场景：
// - Kafka消息延迟投递
// - Redis过期键删除
// - 网关超时处理
// - 订单超时取消
// - 心跳检测
2. 实际工程案例
Kafka：延迟消息、事务超时
Netty：连接超时管理
xxl-job：分布式定时任务
Elasticsearch：定时刷新索引
🤖 RAG (检索增强生成) 编排
核心算法流程
// 简化RAG流程
type RAGPipeline struct {
    documents   []Document     // 文档库
    vectorDB    VectorDB       // 向量数据库
    llm         LLM            // 大语言模型
    promptMgr   PromptManager  // 提示词管理
}

func (r *RAGPipeline) Query(question string) string {
    // 1. 向量化查询
    queryVector := r.vectorDB.Embed(question)
    
    // 2. 相似度检索 (常用算法：余弦相似度)
    topDocs := r.vectorDB.Search(queryVector, topK=5)
    
    // 3. 上下文构建
    context := r.buildContext(topDocs)
    
    // 4. 提示词编排
    prompt := r.promptMgr.BuildPrompt(question, context)
    
    // 5. LLM生成
    return r.llm.Generate(prompt)
}
核心算法与数据结构
算法	作用	实际应用
向量嵌入	文本→向量	Word2Vec, BERT, OpenAI Embeddings
相似度计算	余弦相似度、欧氏距离	FAISS, Milvus
图遍历	RAG链路编排	DAG调度、A*搜索
Top-K检索	K近邻搜索	HNSW, IVF
滑动窗口	上下文窗口管理	Token窗口控制
🎯 推荐系统算法
1. 协同过滤 (Collaborative Filtering)
// 用户相似度计算
func UserSimilarity(userItems map[int][]int) map[int]map[int]float64 {
    // 算法：余弦相似度
    // sim(u,v) = |N(u) ∩ N(v)| / sqrt(|N(u)| * |N(v)|)
    
    // 应用场景：
    // - 电商推荐
    // - 视频推荐
    // - 音乐推荐
}
2. 矩阵分解 (Matrix Factorization)
// 经典ALS算法
func ALS(userItemMatrix [][]float64) (U, V [][]float64) {
    // R ≈ U * V^T
    // U: 用户隐语义矩阵
    // V: 物品隐语义矩阵
    
    // 应用场景：
    // - Netflix推荐
    // - 广告CTR预估
}
3. 深度学习推荐
模型	特点	应用
Wide&Deep	记忆+泛化	Google Play推荐
DeepFM	自动特征交叉	CTR预估
DIN	注意力机制	阿里巴巴广告推荐
Graph Neural Network	图神经网络	社交推荐
🔗 分布式算法
1. 一致性哈希 (Consistent Hashing)
type ConsistentHash struct {
    nodes     []string
    hashRing  map[uint32]string  // 哈希环
    virtualNodes int            // 虚拟节点数
}

func (c *ConsistentHash) Get(key string) string {
    hash := c.hash(key)
    node := c.getNode(hash)
    return node
}

// 应用场景：
// - Redis Cluster分片
// - Cassandra数据分布
// - 负载均衡
// - CDN内容分发
2. Raft/Paxos (分布式一致性)
type RaftNode struct {
    state     NodeState  // Follower, Candidate, Leader
    log       []LogEntry // 日志
    commitIdx int        // 已提交位置
}

// 应用场景：
// - etcd一致性存储
// - Consul服务发现
// - Kafka Controller选举
// - TiKV分布式事务
📈 实时计算算法
1. 滑动窗口统计
type SlidingWindowCounter struct {
    windowSize  int           // 窗口大小
    data        []int64       // 环形缓冲区
    currentIdx  int           // 当前索引
    currentSum  int64         // 当前窗口和
}

// 应用场景：
// - 实时统计API调用量
// - 用户行为分析
// - 监控指标计算
2. HyperLogLog (基数估算)
type HyperLogLog struct {
    registers []uint8  // 寄存器数组
    m         int      // 桶数
}

// 应用场景：
// - UV统计 (DAU/MAU)
// - Redis PFADD
// - 数据库去重基数估算
3. Count-Min Sketch (频次统计)
type CountMinSketch struct {
    depth  int         // 哈希函数数量
    width  int         // 数组宽度
    cms    [][]int64   // 计数数组
}

// 应用场景：
// - Top-K热词统计
// - 流量去重
// - 频率估算
🎨 其他高频应用
1. 搜索排序
算法	应用场景
BM25	文本相关性排序
TF-IDF	关键词权重计算
PageRank	Google网页排序
HNSW	向量搜索
2. 图算法
算法	应用场景
Dijkstra	路径规划、导航
A*	游戏寻路、AI导航
最小生成树	网络拓扑设计
拓扑排序	依赖管理、任务调度
3. 机器学习
算法	应用场景
LR (逻辑回归)	CTR预估
GBDT	推荐、排序
K-Means	用户聚类
决策树	规则引擎、风控
💡 总结：业务优先级
最常用TOP 10：
哈希表 - 缓存、唯一性检查
B+树 - 数据库索引
LRU/LFU - 缓存淘汰
布隆过滤器 - 去重、穿透防护
一致性哈希 - 分布式存储
时间轮 - 定时任务
滑动窗口 - 实时统计
令牌桶 - 限流
LR/GBDT - 推荐、CTR
向量检索 - 语义搜索、RAG
实际工作中：
80%问题用基础数据结构解决（数组、链表、哈希表、堆）
15%需要高级算法（DP、图算法、搜索）
5%需要机器学习/深度学习

安照顺序写一遍，理解一遍