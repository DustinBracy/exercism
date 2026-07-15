fun reverse(input: String): String {
    var reversed = ""
    for (char in input) {
        reversed = "$char$reversed"
    }
    return reversed
}
