package arrays

func RotLeft(a []int32, d int32) []int32 {
    arrLength := int32(len(a))
    
    result := make([]int32, arrLength)
    rotateArray := d
    i := 0
    
    for rotateArray < arrLength {
        result[i] = a[rotateArray]
        i++
        rotateArray++
    }
    
    rotateArray = 0
    
    for rotateArray < d {
        result[i] = a[rotateArray]
		i++
        rotateArray++
    }
    
    return result
}