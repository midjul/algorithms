function lSearch(arr, key) {
    const l = arr.length

    for (let i = 0; i < l; i++) {
        if (arr[i] === key) {
            return true
        }
    }
    return false
}

const arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
const key =11 

console.log(lSearch(arr, key))