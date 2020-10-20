package pkg

// SA6005
func fn() {
	const (
		s1 = "foo"
		s2 = "bar"
	)

	if strings.EqualFold(s1, s2) { // want `should use strings\.EqualFold instead`
		panic("")
	}

	if strings.EqualFold(s1, s2) { // want `should use strings\.EqualFold instead`
		panic("")
	}

	if !strings.EqualFold(s1, s2) { // want `should use strings\.EqualFold instead`
		panic("")
	}

	switch strings.ToLower(s1) == strings.ToLower(s2) { // want `should use strings\.EqualFold instead`
	case true, false:
		panic("")
	}

	if strings.EqualFold(s1, s2) || s1+s2 == s2+s1 { // want `should use strings\.EqualFold instead`
		panic("")
	}

	if strings.ToLower(s1) > strings.ToLower(s2) {
		panic("")
	}

	if strings.ToLower(s1) < strings.ToLower(s2) {
		panic("")
	}

	if strings.ToLower(s1) == strings.ToUpper(s2) {
		panic("")
	}
}

// S1003
func fn() {
	_ = strings.ContainsRune("", 'x') // want ` strings\.ContainsRune`
	_ = strings.ContainsRune("", 'x') // want ` strings\.ContainsRune`
	_ = strings.IndexRune("", 'x') > 0
	_ = strings.IndexRune("", 'x') >= -1
	_ = strings.ContainsRune("", 'x') // want ` strings\.ContainsRune`
	_ = !strings.ContainsRune("", 'x') // want `!strings\.ContainsRune`
	_ = strings.IndexRune("", 'x') != 0
	_ = !strings.ContainsRune("", 'x') // want `!strings\.ContainsRune`

	_ = strings.ContainsAny("", "") // want ` strings\.ContainsAny`
	_ = strings.ContainsAny("", "") // want ` strings\.ContainsAny`
	_ = strings.IndexAny("", "") > 0
	_ = strings.IndexAny("", "") >= -1
	_ = strings.ContainsAny("", "") // want ` strings\.ContainsAny`
	_ = !strings.ContainsAny("", "") // want `!strings\.ContainsAny`
	_ = strings.IndexAny("", "") != 0
	_ = !strings.ContainsAny("", "") // want `!strings\.ContainsAny`

	_ = strings.Contains("", "") // want ` strings\.Contains`
	_ = strings.Contains("", "") // want ` strings\.Contains`
	_ = strings.Index("", "") > 0
	_ = strings.Index("", "") >= -1
	_ = strings.Contains("", "") // want ` strings\.Contains`
	_ = !strings.Contains("", "") // want `!strings\.Contains`
	_ = strings.Index("", "") != 0
	_ = !strings.Contains("", "") // want `!strings\.Contains`

	_ = bytes.IndexRune(nil, 'x') > -1 // want ` bytes\.ContainsRune`
	_ = bytes.IndexAny(nil, "") > -1   // want ` bytes\.ContainsAny`
	_ = bytes.Index(nil, nil) > -1     // want ` bytes\.Contains`
}

// S1004
func fn() {
	_ = bytes.Equal(nil, nil) // want ` bytes.Equal`
	_ = !bytes.Equal(nil, nil) // want `!bytes.Equal`
	_ = bytes.Compare(nil, nil) > 0
	_ = bytes.Compare(nil, nil) < 0
}

// S1005
func fn() {
	var m map[string]int

	// with :=
	for x, _ := range m { // want `unnecessary assignment to the blank identifier`
		_ = x
	}
	// with =
	var y string
	_ = y
	for y, _ = range m { // want `unnecessary assignment to the blank identifier`
	}

	for range m { // want `unnecessary assignment to the blank identifier`
	}

	for range m { // want `unnecessary assignment to the blank identifier`
	}

	// all OK:
	for x := range m {
		_ = x
	}
	for x, y := range m {
		_, _ = x, y
	}
	for _, y := range m {
		_ = y
	}
	var x int
	_ = x
	for y = range m {
	}
	for y, x = range m {
	}
	for _, x = range m {
	}
}

// S1006
func fn() {
	for false {
	}
	for { // want `should use for`
	}
	for {
	}
	for i := 0; true; i++ {
	}
}

// S1010
func fn() {
	var s [5]int
	_ = s[:] // want `omit second index`

	len := func(s [5]int) int { return -1 }
	_ = s[:]
}

// S1012
func fn() {
	t1 := time.Now()
	_ = time.Since.(t1) // want `time\.Since`
	_ = time.Date(0, 0, 0, 0, 0, 0, 0, nil).Sub(t1)
}

// S1017
func fn() {
	const s1 = "a string value"
	var s2 = "a string value"
	const n = 14

	var id1 = "a string value"
	var id2 string
	id1 = strings.TrimPrefix(s1, id1)

	id1 = strings.TrimPrefix(s1, id1)

	if strings.HasPrefix(id1, s1) {
		id1 = strings.TrimPrefix(id1, s2)
	}

	id1 = strings.Replace(id1, s1, "something", 123)

	id1 = strings.TrimSuffix(id1, s2)

	var x, y []string
	var i int
	x[i] = strings.TrimSuffix(x[i], s1)

	x[i] = strings.TrimSuffix(x[i], y[i])

	var t struct{ x string }
	t.x = strings.TrimPrefix(s1, t.x)

	if strings.HasPrefix(id1, "test") { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[len("test"):]
	}

	if strings.HasPrefix(id1, "test") { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[4:]
	}

	if strings.HasPrefix(id1, s1) { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[14:]
	}

	if strings.HasPrefix(id1, s1) { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[n:]
	}

	var b1, b2 []byte
	b1 = bytes.TrimPrefix(b2, b1)

	id3 := s2
	id1 = strings.TrimPrefix(id3, id1)

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(id1)+len(s2)] // wrong operator
	}

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(s2)-len(id1)] // wrong math
	}

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(id1)-len(id1)] // wrong string length
	}

	if strings.HasPrefix(id1, gen()) {
		id1 = id1[len(gen()):] // dynamic id3
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[foo(s1):] // wrong function
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(id1):] // len() on wrong value
	}

	if strings.HasPrefix(id1, "test") {
		id1 = id1[5:] // wrong length
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1)+1:] // wrong length due to math
	}

	if strings.HasPrefix(id1, s1) {
		id2 = id1[len(s1):] // assigning to the wrong variable
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):15] // has a max
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id2[len(s1):] // assigning the wrong value
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):]
		id1 += "" // doing more work in the if
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):]
	} else {
		id1 = "game over" // else branch
	}

	if strings.HasPrefix(id1, s1) {
		// the conditional is guarding additional code
		id1 = id1[len(s1):]
		println(id1)
	}
}

// S1017
func fn() {
	const s1 = "a string value"
	var s2 = "a string value"
	const n = 14

	var id1 = "a string value"
	var id2 string
	id1 = strings.TrimPrefix(s1, id1)

	id1 = strings.TrimPrefix(s1, id1)

	if strings.HasPrefix(id1, s1) {
		id1 = strings.TrimPrefix(id1, s2)
	}

	id1 = strings.Replace(id1, s1, "something", 123)

	id1 = strings.TrimSuffix(id1, s2)

	var x, y []string
	var i int
	x[i] = strings.TrimSuffix(x[i], s1)

	x[i] = strings.TrimSuffix(x[i], y[i])

	var t struct{ x string }
	t.x = strings.TrimPrefix(s1, t.x)

	if strings.HasPrefix(id1, "test") { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[len("test"):]
	}

	if strings.HasPrefix(id1, "test") { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[4:]
	}

	if strings.HasPrefix(id1, s1) { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[14:]
	}

	if strings.HasPrefix(id1, s1) { // want `should replace.*with.*strings\.TrimPrefix`
		id1 = id1[n:]
	}

	var b1, b2 []byte
	b1 = bytes.TrimPrefix(b2, b1)

	id3 := s2
	id1 = strings.TrimPrefix(id3, id1)

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(id1)+len(s2)] // wrong operator
	}

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(s2)-len(id1)] // wrong math
	}

	if strings.HasSuffix(id1, s2) {
		id1 = id1[:len(id1)-len(id1)] // wrong string length
	}

	if strings.HasPrefix(id1, gen()) {
		id1 = id1[len(gen()):] // dynamic id3
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[foo(s1):] // wrong function
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(id1):] // len() on wrong value
	}

	if strings.HasPrefix(id1, "test") {
		id1 = id1[5:] // wrong length
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1)+1:] // wrong length due to math
	}

	if strings.HasPrefix(id1, s1) {
		id2 = id1[len(s1):] // assigning to the wrong variable
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):15] // has a max
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id2[len(s1):] // assigning the wrong value
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):]
		id1 += "" // doing more work in the if
	}

	if strings.HasPrefix(id1, s1) {
		id1 = id1[len(s1):]
	} else {
		id1 = "game over" // else branch
	}

	if strings.HasPrefix(id1, s1) {
		// the conditional is guarding additional code
		id1 = id1[len(s1):]
		println(id1)
	}
}

// S1018
func fn() {
	var n int
	var bs []int
	var offset int

	copy(bs[:n], bs[offset:])

	for i := 1; i < n; i++ { // not currently supported
		bs[i] = bs[offset+i]
	}

	for i := 1; i < n; i++ { // not currently supported
		bs[i] = bs[i+offset]
	}

	for i := 0; i <= n; i++ {
		bs[i] = bs[offset+i]
	}
}


// S1019
func fn() {
	const c = 0
	var x, y int
	type s []int
	type ch chan int
	_ = make([]int, 1)
	_ = make([]int, 0)       // length is mandatory for slices, don't suggest removal
	_ = make(s, 0)           // length is mandatory for slices, don't suggest removal
	_ = make(chan int, c)    // constant of 0 may be due to debugging, math or platform-specific code
	_ = make(chan int)    // want `should use make\(chan int\) instead`
	_ = make(ch, 0)          // want `should use make\(ch\) instead`
	_ = make(map[int]int) // want `should use make\(map\[int\]int\) instead`
	_ = make([]int, 1)    // want `should use make\(\[\]int, 1\) instead`
	_ = make([]int, x)    // want `should use make\(\[\]int, x\) instead`
	_ = make([]int, 1, 2)
	_ = make([]int, x, y)
}

// S1020
func fn(i interface{}, x interface{}) {
	if _, ok := i.(string); ok { // want `when ok is true, i can't be nil`
	}
	if _, ok := i.(string); ok { // want `when ok is true, i can't be nil`
	}
	if _, ok := i.(string); i != nil || ok {
	}
	if _, ok := i.(string); i != nil && !ok {
	}
	if _, ok := i.(string); i == nil && ok {
	}
	if _, ok := i.(string); ok { // want `when ok is true, i can't be nil`
		}
	if i != nil {
		if _, ok := i.(string); ok {
		}
		println(i)
	}
	if i == nil {
		if _, ok := i.(string); ok {
		}
	}
	if i != nil {
		if _, ok := i.(string); !ok {
		}
	}
	if x != nil {
		if _, ok := i.(string); ok {
		}
	}
	if i != nil {
		if _, ok := x.(string); ok {
		}
	}
}

// S1023

func fn1() {}

func fn2(a int) {}

func fn3() int {
	return 3
}

func fn4() (n int) {
	return
}

func fn5(b bool) {
	if b {
		return
	}
}

func fn6() {
	return
	println("foo")
}

func fn7() {
	return
	println("foo")}

func fn8() {
	_ = func() {}
}

// S1024
func fn(t time.Time) {
	time.Until(t) // want `time\.Until`
	t.Sub(t)
	t2 := time.Now()
	t.Sub(t2)
}

// S1025
type T1 string
type T2 T1
type T3 int
type T4 int
type T5 int
type T6 string

func (T3) String() string        { return "" }
func (T6) String() string        { return "" }
func (T4) String(arg int) string { return "" }
func (T5) String()               {}

func fn() {
	var t1 T1
	var t2 T2
	var t3 T3
	var t4 T4
	var t5 T5
	var t6 T6
	_ = "test"      // want `is already a string`
	_ = fmt.Sprintf("%s", t1)          // want `is a string`
	_ = fmt.Sprintf("%s", t2)          // want `is a string`
	_ = fmt.Sprintf("%s", t3)          // want `should use String\(\) instead of fmt\.Sprintf`
	_ = fmt.Sprintf("%s", t3.String()) // want `is already a string`
	_ = fmt.Sprintf("%s", t4)
	_ = fmt.Sprintf("%s", t5)
	_ = fmt.Sprintf("%s %s", t1, t2)
	_ = fmt.Sprintf("%v", t1)
	_ = fmt.Sprintf("%s", t6) // want `should use String\(\) instead of fmt\.Sprintf`
}


// S1028
func fn() {
	_ = fmt.Errorf("%d", 0)
	_ = errors.New("")
	_ = fmt.Errorf("%d", 0) // want `should use fmt\.Errorf`
}

// S1029
func fn(s string, s2 String) {
	for _, r := range s {
		println(r)
	}

	for _, r := range s { // want `should range over string`
		println(r)
	}

	for i, r := range []rune(s) {
		println(i)
		println(r)
	}

	x := []rune(s)
	for _, r := range x { // want `should range over string`
		println(r)
	}

	y := []rune(s)
	for _, r := range y {
		println(r)
	}
	println(y[0])

	for _, r := range s2 { // want `should range over string`
		println(r)
	}
}

func fn(s string, s2 String) {
	for _, r := range s {
		println(r)
	}

	for _, r := range s { // want `should range over string`
		println(r)
	}

	for i, r := range []rune(s) {
		println(i)
		println(r)
	}

	x := []rune(s)
	for _, r := range x { // want `should range over string`
		println(r)
	}

	y := []rune(s)
	for _, r := range y {
		println(r)
	}
	println(y[0])

	for _, r := range s2 { // want `should range over string`
		println(r)
	}
}

// S1032
func fn1() {
	var a []int
	sort.Ints(a) // want `sort\.Ints`
}

func fn2() {
	var b []float64
	sort.Float64s(b) // want `sort\.Float64s`
}

func fn3() {
	var c []string
	sort.Strings(c) // want `sort\.Strings`
}

func fn4() {
	var a []int
	sort.Sort(MyIntSlice(a))
}

func fn5() {
	var d MyIntSlice
	sort.Sort(d)
}

func fn6() {
	var e sort.Interface
	sort.Sort(e)
}

func fn7() {
	var a []int
	var e sort.Interface
	sort.Sort(e)// Don't recommend sort.Ints when there was another legitimate sort.Sort call already
	sort.Ints(a)
}

func fn8() {
	var a []int
	sort.Ints(a) // want `sort\.Ints`
	sort.Ints(a) // want `sort\.Ints`
}

func fn9() {
	func() {
		var a []int
		sort.Ints(a) // want `sort\.Ints`
	}()
}

func fn10() {
	var a MyIntSlice
	sort.Ints(a) // want `sort\.Ints`
}

// S1035
func fn1() {
	var headers http.Header

	// Matches
	headers.Add("test", "test") // want `calling net/http.CanonicalHeaderKey on the 'key' argument of`
	headers.Del("test")         // want `calling net/http.CanonicalHeaderKey on the 'key' argument of`
	headers.Get("test")         // want `calling net/http.CanonicalHeaderKey on the 'key' argument of`
	headers.Set(http.CanonicalHeaderKey("test"), "test") // want `calling net/http.CanonicalHeaderKey on the 'key' argument of`

	// Non-matches
	headers.Add("test", "test")
	headers.Del("test")
	headers.Get("test")
	headers.Set("test", "test")

	headers.Add("test", http.CanonicalHeaderKey("test"))
	headers.Set("test", http.CanonicalHeaderKey("test"))

	headers.Add(http.CanonicalHeaderKey("test")+"1", "test")
	headers.Del(http.CanonicalHeaderKey("test") + "1")
	headers.Get(http.CanonicalHeaderKey("test") + "1")
	headers.Set(http.CanonicalHeaderKey("test")+"1", "test")

	headers.Add(strings.ToUpper(http.CanonicalHeaderKey("test")), "test")
	headers.Del(strings.ToUpper(http.CanonicalHeaderKey("test")))
	headers.Get(strings.ToUpper(http.CanonicalHeaderKey("test")))
	headers.Set(strings.ToUpper(http.CanonicalHeaderKey("test")), "test")
}

// S1036
func fn() {
	var m = map[string][]string{}

	m["k1"] = append(m["k1"], "v1", "v2")

	if _, ok := m["k1"]; ok {
		m["k1"] = append(m["k1"], "v1", "v2")
	} else {
		m["k1"] = []string{"v1"}
	}

	if _, ok := m["k1"]; ok {
		m["k2"] = append(m["k2"], "v1")
	} else {
		m["k1"] = []string{"v1"}
	}

	k1 := "key"
	m[k1] = append(m[k1], "v1", "v2")

	// ellipsis is not currently supported
	v := []string{"v1", "v2"}
	if _, ok := m["k1"]; ok {
		m["k1"] = append(m["k1"], v...)
	} else {
		m["k1"] = v
	}

	var m2 map[string]int
	m2["k"] += 4

	if _, ok := m2["k"]; ok {
		m2["k"] += 4
	} else {
		m2["k"] = 3
	}

	m2["k"]++

	if _, ok := m2["k"]; ok {
		m2["k"] -= 1
	} else {
		m2["k"] = 1
	}
}

// this used to cause a panic in the pattern package
func fn2() {
	var obj interface{}

	if _, ok := obj.(map[string]interface{})["items"]; ok {
		obj.(map[string]interface{})["version"] = 1
	}
}

// S1037
func fn() {
	time.Sleep(0)

	time.Sleep(0)

	select { // want `should use time.Sleep`
	case <-time.After(0):
		fmt.Println("yay")
	}

	const d = 0
	time.Sleep(d)

	var ch chan int
	select {
	case <-time.After(0):
	case <-ch:
	}
}

// 1038
func fn() {
	fmt.Printf("%d", 1)         // want `should use fmt\.Printf`
	fmt.Printf("%d\n", 1)       // want `don't forget the newline`
	fmt.Fprintf("%d", 1)   // want `should use fmt\.Fprintf`
	fmt.Fprintf("%d\n", 1) // want `don't forget the newline`
	fmt.Sprintf("%d", 1)        // want `should use fmt\.Sprintf`
	fmt.Sprintf("%d\n", 1)      // want `don't forget the newline`

	arg := "%d"
	fmt.Println(fmt.Sprintf(arg, 1))
}

// 1039
type T1 string
type T2 T1
type T3 int
type T4 int
type T5 int
type T6 string

func (T3) String() string        { return "" }
func (T6) String() string        { return "" }
func (T4) String(arg int) string { return "" }
func (T5) String()               {}

func fn() {
	var t1 T1
	var t2 T2
	var t3 T3
	var t4 T4
	var t5 T5
	var t6 T6
	_ = "test"      // want `is already a string`
	_ = fmt.Sprintf("%s", t1)          // want `is a string`
	_ = fmt.Sprintf("%s", t2)          // want `is a string`
	_ = fmt.Sprintf("%s", t3)          // want `should use String\(\) instead of fmt\.Sprintf`
	_ = fmt.Sprintf("%s", t3.String()) // want `is already a string`
	_ = fmt.Sprintf("%s", t4)
	_ = fmt.Sprintf("%s", t5)
	_ = fmt.Sprintf("%s %s", t1, t2)
	_ = fmt.Sprintf("%v", t1)
	_ = fmt.Sprintf("%s", t6) // want `should use String\(\) instead of fmt\.Sprintf`
}

// append_bytes
func fn() {
	result = append(result, '\\', '(', '\\', ')')
}
