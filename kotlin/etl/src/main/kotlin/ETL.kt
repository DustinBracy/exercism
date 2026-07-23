object ETL {
    fun transform(source: Map<Int, Collection<Char>>): Map<Char, Int> {
        val output = mutableMapOf<Char, Int>()
        for ((point, letters) in source) {
            for (letter in letters) {
                output.putIfAbsent(letter.lowercaseChar(), point)
            }
        }
        return output
    }
}
