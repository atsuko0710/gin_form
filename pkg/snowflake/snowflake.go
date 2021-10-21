package snowflake

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// Init 用来初始化 snowflake
// startTime 开始时间
// machineID 分布式架构下机器标识
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return 
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}

// GetID 返回生成的id
func GetID() int64 {
	return node.Generate().Int64()
}