package leetcode

// RecentCounter 最近的请求次数
// 写一个 RecentCounter 类来计算特定时间范围内最近的请求。
//
// 请你实现 RecentCounter 类：
//
// RecentCounter() 初始化计数器，请求数为 0 。
// int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。确切地说，返回在 [t-3000, t] 内发生的请求数。
// 保证 每次对 ping 的调用都使用比之前更大的 t 值。
//
// 提示：
//
// 1 <= t <= 109
// 保证每次对 ping 调用所使用的 t 值都 严格递增
// 至多调用 ping 方法 104 次
//
//
//
// 示例 1：
//
// 输入：
// ["RecentCounter", "ping", "ping", "ping", "ping"]
// [[], [1], [100], [3001], [3002]]
// 输出：
// [null, 1, 2, 3, 3]
//
// 解释：
// RecentCounter recentCounter = new RecentCounter();
// recentCounter.ping(1);     // requests = [1]，范围是 [-2999,1]，返回 1
// recentCounter.ping(100);   // requests = [1, 100]，范围是 [-2900,100]，返回 2
// recentCounter.ping(3001);  // requests = [1, 100, 3001]，范围是 [1,3001]，返回 3
// recentCounter.ping(3002);  // requests = [1, 100, 3001, 3002]，范围是 [2,3002]，返回 3
//
//
type RecentCounter struct {
	pings []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (r *RecentCounter) Ping(t int) int {
	r.add(t)

	return r.count(t) - r.count(t-3001)
}

func (r *RecentCounter) add(t int) {
	r.pings = append(r.pings, t)
}

// count 返回 0 - t 时间范围内的请求次数
func (r *RecentCounter) count(t int) int {
	nums := 0
	pings := r.pings
	if pings[0] > t {
		return 0
	}

	if pings[len(pings)-1] < t {
		return 0
	}

	left, right := 0, len(pings)-1
	for left <= right {
		k := (right + left) / 2

		if pings[k] == t {
			nums = k + 1
			break
		}

		if k > 0 && pings[k-1] < t && pings[k] > t {
			nums = k
			break
		}

		if pings[k] > t {
			right = k - 1
		}

		if pings[k] < t {
			left = k + 1
		}

	}
	return nums
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
