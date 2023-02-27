package leetcode

// toGoatLatin  山羊拉丁文
// 给你一个由若干单词组成的句子 sentence ，单词间由空格分隔。每个单词仅由大写和小写英文字母组成。
//
// 请你将句子转换为 “山羊拉丁文（Goat Latin）”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。山羊拉丁文的规则如下：
//
// 如果单词以元音开头（'a', 'e', 'i', 'o', 'u'），在单词后添加"ma"。
// 例如，单词 "apple" 变为 "applema" 。
// 如果单词以辅音字母开头（即，非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
// 例如，单词 "goat" 变为 "oatgma" 。
// 根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从 1 开始。
// 例如，在第一个单词后添加 "a" ，在第二个单词后添加 "aa" ，以此类推。
// 返回将 sentence 转换为山羊拉丁文后的句子。
//
// 提示：
//
// 1 <= sentence.length <= 150
// sentence 由英文字母和空格组成
// sentence 不含前导或尾随空格
// sentence 中的所有单词由单个空格分隔
//
// 示例 1：
//
// 输入：sentence = "I speak Goat Latin"
// 输出："Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
//
// 示例 2：
//
// 输入：sentence = "The quick brown fox jumped over the lazy dog"
// 输出："heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
//
//
func toGoatLatin(s string) string {
	const fillLetter = 'a'
	getSuffix := func(s byte, count int) []byte {
		if !isVowel(s) {
			return append([]byte{s, 'm', 'a'}, repeat(fillLetter, count)...)
		}

		return append([]byte{'m', 'a'}, repeat(fillLetter, count)...)
	}

	ans := make([]byte, 0)
	index := 1
	var suffix []byte

	suffix = getSuffix(s[0], index)
	if isVowel(s[0]) {
		ans = append(ans, s[0])
	}

	for i := 1; i < len(s); i++ {
		letter := s[i]
		switch letter {
		case ' ':
			ans = append(ans, suffix...)
			index++
			suffix = getSuffix(s[i+1], index)
			if !isVowel(s[i+1]) {
				i++ // 跳过下一个单词首字符
			}

		default:
		}
		ans = append(ans, letter)

	}
	ans = append(ans, suffix...)
	return string(ans)
}

func isVowel(s byte) bool {
	switch s {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true

	}
	return false
}

func repeat(s byte, count int) []byte {
	ans := make([]byte, 0, count)
	for i := 0; i < count; i++ {
		ans = append(ans, s)
	}
	return ans
}
