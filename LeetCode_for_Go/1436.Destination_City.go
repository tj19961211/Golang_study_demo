package main

```
You are given the array paths, where paths[i] = [cityAi, cityBi] means there exists a direct path going from cityAi to cityBi. 
Return the destination city, that is, the city without any path outgoing to another city.

It is guaranteed that the graph of paths forms a line without any loop, 
therefore, there will be exactly one destination city.
```

// 大概思路，因为 [cityA, cityB]两个城市间有路能到达目的地(目的地唯一)，所以在 val[1] 即目的地的坐标中，
// 存在Val[0] 不存在的值即是最终目的地(即 return 的值)
func destCity(paths [][]string) string {
	startCities := make(map[string]struct{}, len(paths))
	for _, val := range paths {
		startCities[val[0]] = struct{}{}
	}

	for _, val := range paths {
		_, ok := startCities[val[1]]
		if !ok {
			return val[1]
		}
	}
	return ""
}

