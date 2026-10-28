func carFleet(target int, position []int, speed []int) int {
	table := make([]float32, target)

	for i := range position {
		table[position[i]] = float32(speed[i])
	}

	result := 0
	var previous float32 = 0

	for i := target - 1; i >= 0; i-- {
		if table[i] == 0 {
			continue
		}

		t := float32(target - i) / table[i]

		if t > previous || previous == 0 {
			result++
			previous = t
		}
	}

	return result

	// type Car struct {
	// 	position int
	// 	speed int
	// }

	// cars := make([]Car, len(position))

	// for i := 0; i < len(position); i ++ {
	// 	cars[i] = Car{
	// 		position: position[i],
	// 		speed: speed[i],
	// 	}
	// }

	// // sort by position
	// sort.Slice(cars, func(i, j int) bool {
	// 	return cars[i].position > cars[j].position
	// })

	// fleets := make([]float64, 0, len(position))

	// for _, car := range cars {
	// 	// target=10
	// 	// position=[4,1,0,7]
	// 	// speed=[2,2,1,1]
	// 	// sort = [{7,1}, {4,2}, {1,2}, {0,1}]
	// 	// i = 0 -> d = 3, t = 3
	// 	// i = 1 -> d = 6, t = 3
	// 	// i = 2 -> d = 9, t = 4.5
	// 	// i = 3 -> d = 10, t = 9
	// 	distance := target - car.position

	// 	time := float64(distance) / float64(car.speed)

	// 	if len(fleets) == 0 {
	// 		fleets = append(fleets, time)
	// 		continue
	// 	}

	// 	// // if time more than next fleet  
	// 	if time > fleets[len(fleets)-1] {
	// 		fleets = append(fleets, time)
	// 	}
	// }

	// return len(fleets)
}
