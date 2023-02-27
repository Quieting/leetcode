package leetcode

import (
	"strings"
)

// lengthLongestPath. 文件的最长绝对路径
// 假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：
//
//
//
// 这里将 dir 作为根目录中的唯一目录。dir 包含两个子目录 subdir1 和 subdir2 。subdir1 包含文件 file1.ext 和子目录 subsubdir1；subdir2 包含子目录 subsubdir2，该子目录下包含文件 file2.ext 。
//
// 在文本格式中，如下所示(⟶表示制表符)：
//
// dir
// ⟶ subdir1
// ⟶ ⟶ file1.ext
// ⟶ ⟶ subsubdir1
// ⟶ subdir2
// ⟶ ⟶ subsubdir2
// ⟶ ⟶ ⟶ file2.ext
// 如果是代码表示，上面的文件系统可以写为 "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" 。'\n' 和 '\t' 分别是换行符和制表符。
//
// 文件系统中的每个文件和文件夹都有一个唯一的 绝对路径 ，即必须打开才能到达文件/目录所在位置的目录顺序，所有路径用 '/' 连接。上面例子中，指向 file2.ext 的 绝对路径 是 "dir/subdir2/subsubdir2/file2.ext" 。每个目录名由字母、数字和/或空格组成，每个文件名遵循 name.extension 的格式，其中 name 和 extension由字母、数字和/或空格组成。
//
// 给定一个以上述格式表示文件系统的字符串 input ，返回文件系统中 指向 文件 的 最长绝对路径 的长度 。 如果系统中没有文件，返回 0。
//
// 提示：
//
// 1 <= input.length <= 104
// input 可能包含小写或大写的英文字母，一个换行符 '\n'，一个制表符 '\t'，一个点 '.'，一个空格 ' '，和数字。
//
//
//
// 示例 1：
//
//
// 输入：input = "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext"
// 输出：20
// 解释：只有一个文件，绝对路径为 "dir/subdir2/file.ext" ，路径长度 20
//
// 示例 2：
//
// 输入：input = "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"
// 输出：32
// 解释：存在两个文件：
// "dir/subdir1/file1.ext" ，路径长度 21
// "dir/subdir2/subsubdir2/file2.ext" ，路径长度 32
// 返回 32 ，因为这是最长的路径
//
// 示例 3：
//
// 输入：input = "a"
// 输出：0
// 解释：不存在任何文件
// 示例 4：
//
// 输入：input = "file1.txt\nfile2.txt\nlongfile.txt"
// 输出：12
// 解释：根目录下有 3 个文件。
// 因为根目录中任何东西的绝对路径只是名称本身，所以答案是 "longfile.txt" ，路径长度为 12
//
func lengthLongestPath(s string) int {
	const ext = "."   // 扩展符，判定是否是文件
	const split = "/" // 路径分隔符
	const entry = '\n'
	const tab = '\t'

	longest := 0 // 存储最长的路径长度

	var prefix []string // 解析当前文件或者目录的前缀路径
	var tmp []byte      // 正在解析的文件或者目录
	var level int       // 即将解析的文件或者目录所属层级

	for i := 0; i < len(s); i++ {
		letter := int32(s[i])
		switch letter {
		case entry:
			// 验证当前路径是否为文件
			path := string(tmp)
			if strings.Contains(path, ext) { // 文件
				longest = max(longest, len(strings.Join(append(prefix[:level], path), split)))
			} else { // 目录
				prefix = append(prefix[:level], path)
			}

			level = 0
			tmp = tmp[:0]
		case tab:
			level++
		default:
			if level+1 > len(prefix) {
				for l := 0; l < level-len(prefix); l++ {
					prefix = append(prefix, "")

				}
			} else if level > 0 {
				prefix = prefix[:level]
			}
			tmp = append(tmp, s[i])
		}
	}

	path := string(tmp)
	if strings.Contains(path, ext) { // 文件
		longest = max(longest, len(strings.Join(append(prefix[:level], path), split)))
	}

	return longest
}

// lengthLongestPath1 官方解法
func lengthLongestPath1(input string) (ans int) {
	n := len(input)
	level := make([]int, n+1)
	for i := 0; i < n; {
		// 检测当前文件的深度
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		// 统计当前文件名的长度
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++ // 跳过换行符

		if depth > 1 {
			length += level[depth-1] + 1
		}
		if isFile {
			ans = max(ans, length)
		} else {
			level[depth] = length
		}
	}
	return
}
