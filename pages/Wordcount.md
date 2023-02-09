- ```
  func WordCount(s string) (m map[string]int) {
  	m = make(map[string]int)
  	for _, w:= range(strings.Fields(s)) {
  		m[w]++
  	}
  	return
  }
  ```
-