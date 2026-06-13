func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, word := range strs {
        var count [26]int

        for _, letter := range word {
            count[letter - 'a']++
        }

        groups[count] = append(groups[count], word)
    }

    var result [][]string
    for _, group := range groups {
        result = append(result, group)
    }

    return result
}
