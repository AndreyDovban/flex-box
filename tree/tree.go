package tree

import "slices"

func BuidHashList(objects [][]string) map[string][]string {
	hash := map[string][]string{}

	for _, arr := range objects {
		for i := 0; i < len(arr)-1; i++ {
			_, ok := hash[arr[i]]
			if ok {
				children := hash[arr[i]]
				if !slices.Contains(children, arr[i+1]) {
					children = append(children, arr[i+1])
					hash[arr[i]] = children
				}
			} else {
				if i != len(arr)-1 {
					hash[arr[i]] = []string{}
				}
			}
		}
	}

	hash[""] = []string{objects[0][0]}

	return hash
}
