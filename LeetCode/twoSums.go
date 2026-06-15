package LeetCode

import(

)
// Complexidade: O(n²)

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}


// Melhor solução
// Utilizando um HashMap
// Complexidade: O(n)

func twoSumHashMap(nums []int, target int) []int {
    seen := make(map[int]int) // valor -> índice

    for i, num := range nums {
        complement := target - num // Calculando o que é necessário

        if j, ok := seen[complement]; ok { // Se ja foi visto o cmplemento
            return []int{j, i} // então os indices j e i
        }

        seen[num] = i // se não achou ainda, registra a atual
    }

    return nil
}