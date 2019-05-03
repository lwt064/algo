package stringmatch

var primes = []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,
71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,
167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,
271,277,281,283,293,307,311,313,317,331,337,347,349,353,359,367,373,379,383,
389,397,401,409,419,421,431,433,439,443,449,457,461,463,467,479,487,491,499,
503,509,521,523,541,547,557,563,569,571,577,587,593,599,601,607,613,617,619,
631,641,643,647,653,659,661,673,677,683,691,701,709,719,727,733,739,743,751,
757,761,769,773,787,797,809,811,821,823,827,829,839,853,857,859,863,877,881,
883,887,907,911,919,929,937,941,947,953,967,971,977,983,991,997}

func isMatch(s string, froms int, p string) bool {
	return s[froms:froms+len(p)] == p
}

func getHash(s string, froms int, length int) uint64 {
	hash := uint64(0)
	tmpS := s[froms:froms+length]
	for _, c := range tmpS {
		hash += uint64(primes[int(c)])
	}
	return hash
}

func RKMatch(s string, p string) int {
	if len(s) < len(p) {
		return -1
	}
	sHash := getHash(s, 0, len(p))
	pHash := getHash(p, 0, len(p))
	if sHash == pHash {
		return 0
	}

	end := len(s) - len(p)
	for i := 1; i <= end; i++ {
		lastStartChar := int(s[i-1])
		curEndChar := int(s[i+len(p)-1])
		sHash = sHash - uint64(primes[lastStartChar]) + uint64(primes[curEndChar])
		if sHash == pHash && isMatch(s, i, p) {
			return i
		}
	}
	return -1
}

// 生成坏字符在模式串中的位置表，-1表示不存在，如果出现多次，则记录最后一次的位置
func GenerateBC(p string) []int {
	const CHAC_SIZE = 256
	bc := make([]int, CHAC_SIZE)
	for i := 0; i < CHAC_SIZE; i++ {
		bc[i] = -1
	}
	for i := 0; i < len(p); i++ {
		c := int(p[i])
		bc[c] = i
	}
	return bc
}

// 生成模式串的后缀子串u在模式串中重复出现的位置表suffix, 由于最后一个字符固定，suffix下标为u的长度，可以唯一确定一个u
// 生成模式串的后缀子串u与模式串的前缀子串匹配的位置表prefix, 由于最后一个字符固定，prefix下标为u的长度，可以唯一确定一个u
func GenerateGS(p string) ([]int, []bool) {
	m := len(p)
	suffix := make([]int, m)
	prefix := make([]bool, m)
	for i, _ := range(p) {
		suffix[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ {  // 求p[0:i+1) 与 p[0:m)的公共后缀
		j := i
		for j >= 0 && p[j] == p[m-1-i+j] {
			suffix[i-j+1] = j   // i-j+1 为公共子串长度
			j--
		}
		if j < 0 {
			prefix[i+1] = true 
		}
	}
	return suffix, prefix
}

func moveByGS(j int, m int, suffix []int, prefix []bool) int {
	k := m - 1 -j
	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}
	// j是坏字符位置，j+1为已匹配后缀串开始位置，j+2为后缀串的子串开始位置
	for r := j+2; r <= m-1; r++ {
		if prefix[m-r] {  // m-r为模式串后缀串子串的长度，prefix[m-r]=true 表示这段子串与前缀串匹配，可以将模式串头部移动到r的位置
			return r
		}
	}
	return m
}

func BMMatch(s string, p string) int {
	n, m := len(s), len(p)
	if m < 1 {
		return -1
	}
	bc := GenerateBC(p)
	suffix, prefix := GenerateGS(p)
	for i := 0; i <= n - m; {
		j := m-1
		for j >= 0 {
			if s[i+j] != p[j] {
				break
			}
			j--
		}
		if j < 0 {
			return i
		}
		x := j - bc[int(s[i+j])] // 根据坏字符规则，后移的距离
		y := 0
		if j < m-1 {
			y = moveByGS(j, m, suffix, prefix)
		}
		if x > y {
			i = i + x
		} else {
			i = i + y
		}
	}
	return -1
}