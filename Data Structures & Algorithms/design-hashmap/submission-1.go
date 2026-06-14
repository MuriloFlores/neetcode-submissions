type MyHashMap struct {
	buckets [][]Pair
}

type Pair struct {
	Key 	int
	Value 	int
}

func Constructor() MyHashMap {
	return MyHashMap {
		buckets: make([][]Pair, 1000),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % 1000
}


func (this *MyHashMap) Put(key int, value int) {
    index := this.hash(key)
	
	for i, pair := range this.buckets[index] {
		if pair.Key == key {
			this.buckets[index][i].Value = value
			return
		}
	}

	newPair := Pair{Key: key, Value: value}
	this.buckets[index] = append(this.buckets[index], newPair)
}

func (this *MyHashMap) Get(key int) int {
    index := this.hash(key)

	for i, pair := range this.buckets[index] {
		if pair.Key == key {
			return this.buckets[index][i].Value
		}
	}

	return -1
}

func (this *MyHashMap) Remove(key int) {
    index := this.hash(key)

	for i, pair := range this.buckets[index] {
		if pair.Key == key {
			lastItem := this.buckets[index][len(this.buckets[index]) - 1]
			this.buckets[index][i] = lastItem

			this.buckets[index] = this.buckets[index][:len(this.buckets[index]) - 1]
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */