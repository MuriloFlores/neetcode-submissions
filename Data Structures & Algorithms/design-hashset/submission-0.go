type MyHashSet struct {
	buckets [][]int
}

func Constructor() MyHashSet {
	return MyHashSet{
		buckets: make([][]int, 1000),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % 1000
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		index := this.hash(key)

		this.buckets[index] = append(this.buckets[index], key)
	}
}

func (this *MyHashSet) Remove(key int) {
	index := this.hash(key)

	for i, value := range this.buckets[index] {
		if value == key {
			lastItem := this.buckets[index][len(this.buckets[index])-1]
			this.buckets[index][i] = lastItem

			this.buckets[index] = this.buckets[index][:len(this.buckets[index])-1]
			return
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	index := this.hash(key)

	for _, value := range this.buckets[index] {
		if value == key {
			return true
		}
	}

	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
