package algorithm

import (
	"math/rand"
)

// 子集划分算法: 参考：https://xie.infoq.cn/article/2075b4d8473854bd606ab49e0
// backends: 表示当前所有的后端
// clientId:  客户端唯一ID,某一台客户端
// subsetSize: 子集大小,也就是要连接多少个后端
func SubSet(backends []string, clientId, subsetSize int) []string {
	subsetCount := len(backends) / subsetSize
	round := clientId / subsetCount

	r := rand.New(rand.NewSource(int64(round)))
	r.Shuffle(len(backends), func(i, j int) {
		backends[i], backends[j] = backends[j], backends[i]
	})

	subsetId := clientId % subsetCount

	start := subsetId * subsetSize
	return backends[start : start+subsetSize]
}
