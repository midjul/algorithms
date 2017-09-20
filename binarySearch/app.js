function bSearch(n, hs) {
    const l = hs.length
    let low = 0
    let hight = l - 1
    while (low <= hight) {
        let median = Math.floor((low + hight) / 2)
        if (hs[median] < n) {
            low = median + 1
        }
        else {
            hight = median - 1
        }
    }
    if (low === l || hs[low] !== n) {
        return false
    }
    return true
}

const hs = [1, 4, 6, 7, 9, 12, 24, 67, 74]
const n = 90
console.log(bSearch(n, hs))